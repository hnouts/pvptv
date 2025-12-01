package http

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/argon2"
)

// registerAdminRoutes wires admin routes into the Gin engine.
func registerAdminRoutes(r *gin.Engine, db *sql.DB) {
	admin := r.Group("/admin")
	{
		admin.GET("/login", adminLoginForm)
		admin.POST("/login", adminLoginPost)
		admin.GET("/logout", adminLogout)

		// Protected routes
		admin.Use(requireAdmin)
		admin.GET("/channels", adminChannelsList(db))
		admin.GET("/channels/new", adminChannelNew(db))
		admin.POST("/channels/new", adminChannelCreate(db))
		admin.GET("/channels/:id/edit", adminChannelEdit(db))
		admin.POST("/channels/:id", adminChannelUpdate(db))
		admin.POST("/channels/:id/delete", adminChannelDelete(db))
	}
}

const adminSessionCookie = "pvptv_admin"

// Very small, cookie-based admin session.
func setAdminSession(c *gin.Context, username string) {
	// For v2 we use a simple opaque cookie; in a follow-up we can switch to signed/encrypted.
	// Secure=false for dev/localhost simplicity. In prod, use true if HTTPS.
	secure := os.Getenv("GIN_MODE") == "release"
	c.SetCookie(adminSessionCookie, username, int((7*24*time.Hour).Seconds()), "/", "", secure, true)
}

func clearAdminSession(c *gin.Context) {
	secure := os.Getenv("GIN_MODE") == "release"
	c.SetCookie(adminSessionCookie, "", -1, "/", "", secure, true)
}

func requireAdmin(c *gin.Context) {
	if v, err := c.Cookie(adminSessionCookie); err == nil && v != "" {
		c.Next()
		return
	}
	c.Redirect(http.StatusFound, "/admin/login")
	c.Abort()
}

func adminLoginForm(c *gin.Context) {
	c.HTML(http.StatusOK, "admin_login", gin.H{
		"Canonical": canonicalURL(c),
		"Title":     "Admin Login",
		"HideNav":   true,
	})
}

func adminLoginPost(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if validateAdminCredentials(username, password) {
		setAdminSession(c, username)
		c.Redirect(http.StatusFound, "/admin/channels")
		return
	}

	c.HTML(http.StatusUnauthorized, "admin_login", gin.H{
		"Canonical": canonicalURL(c),
		"Title":     "Admin Login",
		"Error":     "Invalid credentials",
		"HideNav":   true,
	})
}

func adminLogout(c *gin.Context) {
	clearAdminSession(c)
	c.Redirect(http.StatusFound, "/admin/login")
}

// validateAdminCredentials checks username + password against env-configured values.
// ADMIN_USERNAME, ADMIN_PASSWORD_HASH (Argon2id encoded as: salt:hexhash, where both parts are hex).
func validateAdminCredentials(username, password string) bool {
	wantUser := os.Getenv("ADMIN_USERNAME")
	hashEnv := os.Getenv("ADMIN_PASSWORD_HASH")
	if username == "" || password == "" || wantUser == "" || hashEnv == "" {
		return false
	}
	if username != wantUser {
		return false
	}

	// For simplicity, we assume hashEnv is a raw hex Argon2id hash with a fixed salt value baked in.
	// In a real deployment you would store full parameters (time, memory, threads, salt, hash).
	salt := []byte("pvptv-admin") // fixed salt for v2; replace with per-user salt if you add more users later.

	computed := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
	return hashEnv == fmt.Sprintf("%x", computed)
}

func adminChannelsList(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		search := c.Query("search")
		query := "SELECT id, display_name, twitch_login, is_published, sort_weight FROM channels"
		args := []interface{}{}
		
		if search != "" {
			query += " WHERE display_name ILIKE $1 OR twitch_login ILIKE $1"
			args = append(args, "%"+search+"%")
		}
		
		query += " ORDER BY sort_weight DESC, display_name ASC"
		
		rows, err := db.Query(query, args...)
		if err != nil {
			c.String(http.StatusInternalServerError, "DB Error: %v", err)
			return
		}
		defer rows.Close()

		var channels []map[string]interface{}
		for rows.Next() {
			var id, name, slug string
			var published bool
			var sortWeight int
			if err := rows.Scan(&id, &name, &slug, &published, &sortWeight); err != nil {
				continue
			}
			channels = append(channels, map[string]interface{}{
				"ID":         id,
				"Name":       name,
				"Slug":       slug,
				"Published":  published,
				"SortWeight": sortWeight,
			})
		}

		c.HTML(http.StatusOK, "admin_channels", gin.H{
			"Canonical": canonicalURL(c),
			"Channels":  channels,
			"Search":    search,
			"Title":     "Admin Channels",
		})
	}
}

func adminChannelEdit(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		
		// Fetch Channel
		var name, slug string
		var published bool
		var sortWeight int
		var socialsJSON []byte
		err := db.QueryRow("SELECT display_name, twitch_login, is_published, sort_weight, socials FROM channels WHERE id = $1", id).
			Scan(&name, &slug, &published, &sortWeight, &socialsJSON)
		if err == sql.ErrNoRows {
			c.String(http.StatusNotFound, "Channel not found")
			return
		}
		if err != nil {
			c.String(http.StatusInternalServerError, "DB Error: %v", err)
			return
		}

		// Parse socials JSON
		var socials map[string]string
		if len(socialsJSON) > 0 {
			if err := json.Unmarshal(socialsJSON, &socials); err != nil {
				socials = make(map[string]string)
			}
		} else {
			socials = make(map[string]string)
		}

		// Fetch Assigned Specs with primary flag and class info
		assigned := make(map[string]bool)
		assignedClasses := make(map[string]bool) // Track which classes have selected specs
		primarySpecID := ""
		rows, err := db.Query(`
			SELECT cs.spec_id, cs.is_primary, c.id as class_id
			FROM channel_specs cs
			JOIN specs s ON cs.spec_id = s.id
			JOIN classes c ON s.class_id = c.id
			WHERE cs.channel_id = $1
		`, id)
		if err == nil {
			defer rows.Close()
			for rows.Next() {
				var specID, classID string
				var isPrimary bool
				rows.Scan(&specID, &isPrimary, &classID)
				assigned[specID] = true
				assignedClasses[classID] = true
				if isPrimary {
					primarySpecID = specID
				}
			}
		}

		// Fetch All Classes with their Specs
		rows, err = db.Query(`
			SELECT c.id, c.class_name, s.id, s.spec_name
			FROM classes c
			LEFT JOIN specs s ON c.id = s.class_id
			WHERE c.game_slug = 'wow'
			ORDER BY c.class_name, s.spec_name
		`)
		if err != nil {
			c.String(http.StatusInternalServerError, "DB Error: %v", err)
			return
		}
		defer rows.Close()

		// Group specs by class
		type SpecView struct {
			ID         string
			Name       string
			Selected   bool
			IsPrimary  bool
		}
		type ClassView struct {
			ID       string
			Name     string
			Selected bool // True if any spec from this class is selected
			Specs    []SpecView
		}
		
		// Use a map to build classes, but we want stable order.
		var classes []*ClassView
		classMap := make(map[string]*ClassView)

		for rows.Next() {
			var classID, className, specID, specName sql.NullString
			rows.Scan(&classID, &className, &specID, &specName)
			
			if !classID.Valid {
				continue
			}
			
			// Get or create class
			class, exists := classMap[classID.String]
			if !exists {
				class = &ClassView{
					ID:       classID.String,
					Name:     className.String,
					Selected: assignedClasses[classID.String],
					Specs:    []SpecView{},
				}
				classMap[classID.String] = class
				classes = append(classes, class)
			}
			
			// Add spec if it exists
			if specID.Valid && specName.Valid {
				class.Specs = append(class.Specs, SpecView{
					ID:        specID.String,
					Name:      specName.String,
					Selected:  assigned[specID.String],
					IsPrimary: specID.String == primarySpecID,
				})
				// Update class selected status
				if assigned[specID.String] {
					class.Selected = true
				}
			}
		}

		c.HTML(http.StatusOK, "admin_channel_edit", gin.H{
			"Canonical": canonicalURL(c),
			"Channel": gin.H{
				"ID":         id,
				"Name":       name,
				"Slug":       slug,
				"Published":  published,
				"SortWeight": sortWeight,
				"Socials":    socials,
			},
			"Classes":      classes,
			"PrimarySpecID": primarySpecID,
			"Title":        "Edit Channel",
		})
	}
}

func adminChannelUpdate(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		displayName := c.PostForm("display_name")
		twitchLogin := c.PostForm("twitch_login")
		sortWeight := c.PostForm("sort_weight")
		isPublished := c.PostForm("is_published") == "true"
		specIDs := c.PostFormArray("spec_ids")
		primarySpecID := c.PostForm("primary_spec_id")

		// Parse socials from form
		socials := make(map[string]string)
		socialKeys := []string{"twitter", "youtube", "discord", "instagram", "website"}
		for _, key := range socialKeys {
			if val := c.PostForm("social_" + key); val != "" {
				socials[key] = val
			}
		}
		socialsJSON, _ := json.Marshal(socials)

		// Update Channel
		_, err := db.Exec(`
			UPDATE channels 
			SET display_name = $1, twitch_login = $2, is_published = $3, sort_weight = $4, socials = $5
			WHERE id = $6
		`, displayName, twitchLogin, isPublished, sortWeight, socialsJSON, id)
		if err != nil {
			c.String(http.StatusInternalServerError, "DB Error updating channel: %v", err)
			return
		}

		// Update Specs (Delete all, then Insert)
		tx, err := db.Begin()
		if err != nil {
			c.String(http.StatusInternalServerError, "DB Error: %v", err)
			return
		}
		
		_, err = tx.Exec("DELETE FROM channel_specs WHERE channel_id = $1", id)
		if err != nil {
			tx.Rollback()
			c.String(http.StatusInternalServerError, "DB Error deleting specs: %v", err)
			return
		}

		stmt, err := tx.Prepare("INSERT INTO channel_specs (channel_id, spec_id, is_primary) VALUES ($1, $2, $3)")
		if err != nil {
			tx.Rollback()
			c.String(http.StatusInternalServerError, "DB Error prepare: %v", err)
			return
		}
		defer stmt.Close()

		for _, specID := range specIDs {
			isPrimary := specID == primarySpecID
			if _, err := stmt.Exec(id, specID, isPrimary); err != nil {
				tx.Rollback()
				c.String(http.StatusInternalServerError, "DB Error inserting spec: %v", err)
				return
			}
		}

		if err := tx.Commit(); err != nil {
			c.String(http.StatusInternalServerError, "DB Error commit: %v", err)
			return
		}

		c.Redirect(http.StatusFound, "/admin/channels")
	}
}

func adminChannelNew(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Fetch All Classes with their Specs (same structure as edit)
		rows, err := db.Query(`
			SELECT c.id, c.class_name, s.id, s.spec_name
			FROM classes c
			LEFT JOIN specs s ON c.id = s.class_id
			WHERE c.game_slug = 'wow'
			ORDER BY c.class_name, s.spec_name
		`)
		if err != nil {
			c.String(http.StatusInternalServerError, "DB Error: %v", err)
			return
		}
		defer rows.Close()

		type SpecView struct {
			ID   string
			Name string
		}
		type ClassView struct {
			ID       string
			Name     string
			Selected bool
			Specs    []SpecView
		}

		var classes []*ClassView
		classMap := make(map[string]*ClassView)

		for rows.Next() {
			var classID, className, specID, specName sql.NullString
			rows.Scan(&classID, &className, &specID, &specName)

			if !classID.Valid {
				continue
			}

			// Get or create class
			class, exists := classMap[classID.String]
			if !exists {
				class = &ClassView{
					ID:       classID.String,
					Name:     className.String,
					Selected: false,
					Specs:    []SpecView{},
				}
				classMap[classID.String] = class
				classes = append(classes, class)
			}

			// Add spec if it exists
			if specID.Valid && specName.Valid {
				class.Specs = append(class.Specs, SpecView{
					ID:   specID.String,
					Name: specName.String,
				})
			}
		}

		c.HTML(http.StatusOK, "admin_channel_new", gin.H{
			"Canonical": canonicalURL(c),
			"Classes":   classes,
			"Title":     "New Channel",
		})
	}
}

func adminChannelCreate(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		displayName := c.PostForm("display_name")
		twitchLogin := c.PostForm("twitch_login")
		sortWeight := c.PostForm("sort_weight")
		isPublished := c.PostForm("is_published") == "true"
		specIDs := c.PostFormArray("spec_ids")
		primarySpecID := c.PostForm("primary_spec_id")

		// Parse socials from form
		socials := make(map[string]string)
		socialKeys := []string{"twitter", "youtube", "discord", "instagram", "website"}
		for _, key := range socialKeys {
			if val := c.PostForm("social_" + key); val != "" {
				socials[key] = val
			}
		}
		socialsJSON, _ := json.Marshal(socials)

		// Insert Channel
		var id string
		err := db.QueryRow(`
			INSERT INTO channels (display_name, twitch_login, is_published, sort_weight, socials)
			VALUES ($1, $2, $3, $4, $5)
			RETURNING id
		`, displayName, twitchLogin, isPublished, sortWeight, socialsJSON).Scan(&id)
		if err != nil {
			c.String(http.StatusInternalServerError, "DB Error creating channel: %v", err)
			return
		}

		// Insert Specs
		if len(specIDs) > 0 {
			tx, err := db.Begin()
			if err != nil {
				c.String(http.StatusInternalServerError, "DB Error: %v", err)
				return
			}

			stmt, err := tx.Prepare("INSERT INTO channel_specs (channel_id, spec_id, is_primary) VALUES ($1, $2, $3)")
			if err != nil {
				tx.Rollback()
				c.String(http.StatusInternalServerError, "DB Error prepare: %v", err)
				return
			}
			defer stmt.Close()

			for _, specID := range specIDs {
				isPrimary := specID == primarySpecID
				if _, err := stmt.Exec(id, specID, isPrimary); err != nil {
					tx.Rollback()
					c.String(http.StatusInternalServerError, "DB Error inserting spec: %v", err)
					return
				}
			}

			if err := tx.Commit(); err != nil {
				c.String(http.StatusInternalServerError, "DB Error commit: %v", err)
				return
			}
		}

		c.Redirect(http.StatusFound, "/admin/channels")
	}
}

func adminChannelDelete(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		_, err := db.Exec("DELETE FROM channels WHERE id = $1", id)
		if err != nil {
			c.String(http.StatusInternalServerError, "DB Error deleting channel: %v", err)
			return
		}
		c.Redirect(http.StatusFound, "/admin/channels")
	}
}
