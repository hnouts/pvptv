package http

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"sort"
	"strings"
	"time"

	"arenatv/internal/twitch"

	"github.com/gin-gonic/gin"
)

// getAllClasses fetches all classes for the sidebar navigation.
func getAllClasses(db *sql.DB) ([]map[string]string, error) {
	rows, err := db.Query("SELECT class_slug, class_name FROM classes WHERE game_slug = 'wow' ORDER BY class_name ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var classes []map[string]string
	for rows.Next() {
		var slug, name string
		if err := rows.Scan(&slug, &name); err != nil {
			continue
		}
		// Use static icon file path instead of DB SVG
		// Assumes icons are at /assets/icons/{slug}.jpg
		classes = append(classes, map[string]string{
			"Slug":     slug,
			"Name":     name,
			"IconPath": fmt.Sprintf("/assets/icons/%s.jpg", slug),
		})
	}
	return classes, nil
}

// homeHandler renders the public home page.
func homeHandler(db *sql.DB, twitchClient *twitch.HelixClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		classes, err := getAllClasses(db)
		if err != nil {
			c.String(http.StatusInternalServerError, "DB Error: %v", err)
			return
		}

		// Featured Stream Logic
		var featuredStream map[string]interface{}
		var onlineCount int
		
		if twitchClient != nil {
			// 1. Get candidate channels (high priority first)
			rows, err := db.Query(`
				SELECT twitch_login, display_name 
				FROM channels 
				WHERE is_published = true 
				ORDER BY sort_weight DESC 
				LIMIT 100
			`)
			if err == nil {
				defer rows.Close()
				var logins []string
				var channelMap = make(map[string]string) // login -> display_name
				
				for rows.Next() {
					var login, name string
					if err := rows.Scan(&login, &name); err == nil {
						logins = append(logins, login)
						channelMap[strings.ToLower(login)] = name
					}
				}

				// 2. Check who is online
				if len(logins) > 0 {
					ctx := c.Request.Context()
					// Twitch API limit is usually 100, which matches our limit
					resp, err := twitchClient.GetStreamsByLogin(ctx, logins)
					if err == nil && len(resp.Data) > 0 {
						onlineCount = len(resp.Data)
						
						// 3. Pick a random one (or just the first one since we sorted by weight)
						rand.Seed(time.Now().UnixNano())
						rand.Shuffle(len(resp.Data), func(i, j int) {
							resp.Data[i], resp.Data[j] = resp.Data[j], resp.Data[i]
						})
						stream := resp.Data[0]
						
						featuredStream = map[string]interface{}{
							"login":       stream.UserLogin,
							"displayName": channelMap[strings.ToLower(stream.UserLogin)],
							"title":       stream.Title,
							"viewers":     stream.ViewerCount,
							"game":        stream.GameName,
							"online":      true,
						}
					}
				}
			}
		}

		canonical := canonicalURL(c)
		jsonLD := map[string]interface{}{
			"@context": "https://schema.org",
			"@type":    "WebSite",
			"name":     "PvPtv",
			"url":      "https://pvptv.hnts.dev/",
		}

		c.HTML(http.StatusOK, "home", gin.H{
			"Canonical":       canonical,
			"OGImage":         "/assets/logo.png",
			"Classes":         classes,
			"FeaturedStream":  featuredStream,
			"OnlineCount":     onlineCount,
			"Title":           "PvPtv – Best WoW PvP Streamers & Arena Gameplay",
			"MetaDescription": "The ultimate directory for World of Warcraft PvP streamers. Find top Arena and RBG players by class and spec. Watch live WoW PvP now.",
			"JSONLD":          jsonLD,
		})
	}
}

// classHandler renders a class listing page (e.g., /rogue).
func classHandler(db *sql.DB, twitchClient *twitch.HelixClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		classSlug := c.Param("class")
		
		// Redirect Plunderstorm to home
		if classSlug == "plunderstorm" {
			c.Redirect(http.StatusMovedPermanently, "/")
			return
		}

		classes, err := getAllClasses(db)
		if err != nil {
			c.String(http.StatusInternalServerError, "getAllClasses Error: %v", err)
			return
		}

		// Get class info
		var classID int64
		var className string
		err = db.QueryRow(
			"SELECT id, class_name FROM classes WHERE game_slug = 'wow' AND class_slug = $1",
			classSlug,
		).Scan(&classID, &className)
		if err == sql.ErrNoRows {
			c.HTML(http.StatusNotFound, "404", gin.H{})
			return
		}
		if err != nil {
			c.String(http.StatusInternalServerError, "GetClass Error: %v", err)
			return
		}

		// Get channels for this class (via specs)
		// Only show channels that have specs assigned to this class
		rows, err := db.Query(`
			SELECT DISTINCT
				c.id, c.twitch_login, c.display_name, c.socials, c.sort_weight
			FROM channels c
			INNER JOIN channel_specs cs ON c.id = cs.channel_id
			INNER JOIN specs s ON cs.spec_id = s.id
			WHERE s.class_id = $1
			  AND c.is_published = true
			ORDER BY c.sort_weight DESC, c.display_name ASC
			LIMIT 100
		`, classID)
		if err != nil {
			c.String(http.StatusInternalServerError, "QueryChannels Error: %v", err)
			return
		}
		defer rows.Close()

		var channels []map[string]interface{}
		var logins []string
		for rows.Next() {
			var id, login, name string
			var socialsJSON []byte
			var sortWeight int
			if err := rows.Scan(&id, &login, &name, &socialsJSON, &sortWeight); err != nil {
				continue
			}
			var socials map[string]string
			json.Unmarshal(socialsJSON, &socials)
			channels = append(channels, map[string]interface{}{
				"id":       id,
				"slug":     login,
				"name":     name,
				"socials":  socials,
			})
			logins = append(logins, login)
		}

		// Fetch Twitch stream data
		streamData := make(map[string]interface{})
		if len(logins) > 0 && twitchClient != nil {
			ctx := c.Request.Context()
			resp, err := twitchClient.GetStreamsByLogin(ctx, logins)
			if err == nil {
				for _, stream := range resp.Data {
					streamData[strings.ToLower(stream.UserLogin)] = map[string]interface{}{
						"online":      true,
						"title":       stream.Title,
						"viewers":     stream.ViewerCount,
						"game":        stream.GameName,
					}
				}
			}
		}

		// Mark channels as online/offline
		for i := range channels {
			slug := channels[i]["slug"].(string)
			if data, ok := streamData[strings.ToLower(slug)]; ok {
				channels[i]["stream"] = data
			} else {
				channels[i]["stream"] = map[string]interface{}{"online": false}
			}
		}

		// Sort channels: Online first -> Viewers (desc) -> Name (asc)
		sort.SliceStable(channels, func(i, j int) bool {
			streamI := channels[i]["stream"].(map[string]interface{})
			streamJ := channels[j]["stream"].(map[string]interface{})

			onlineI, _ := streamI["online"].(bool)
			onlineJ, _ := streamJ["online"].(bool)

			if onlineI != onlineJ {
				return onlineI // Online first
			}

			if onlineI {
				viewersI, _ := streamI["viewers"].(int)
				viewersJ, _ := streamJ["viewers"].(int)
				if viewersI != viewersJ {
					return viewersI > viewersJ // Higher viewers first
				}
			}

			// Fallback to SortWeight then Name (already sorted by SQL, but good to be explicit or stable)
			// SQL order was: sort_weight DESC, display_name ASC.
			// Since we used SliceStable, the SQL order is preserved for ties.
			// But since we are re-sorting the whole list based on online status, we might break SQL order if we don't include it?
			// Actually SliceStable preserves original order of equal elements.
			// The original order was sort_weight DESC, display_name ASC.
			// So if both are online (or both offline) and viewers are equal, they stay in SQL order.
			return false 
		})

		canonical := canonicalURL(c)
		title := fmt.Sprintf("%s WoW PvP Streamers - Watch Live | PvPtv", className)
		desc := fmt.Sprintf("Watch top %s WoW PvP streamers live. Best %s arena and RBG players online now.", className, className)

		jsonLD := map[string]interface{}{
			"@context":    "https://schema.org",
			"@type":       "CollectionPage",
			"name":        title,
			"description": desc,
			"url":         canonical,
		}

		c.HTML(http.StatusOK, "class", gin.H{
			"Canonical":       canonical,
			"OGImage":         "/assets/logo.png",
			"ClassSlug":       classSlug,
			"ClassName":       className,
			"Channels":        channels,
			"Classes":         classes,
			"Title":           title,
			"MetaDescription": desc,
			"JSONLD":          jsonLD,
		})
	}
}

// streamHandler renders an individual stream page (e.g., /rogue/pikabooirl).
func streamHandler(db *sql.DB, twitchClient *twitch.HelixClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		classSlug := c.Param("class")
		channelSlug := c.Param("channel")

		// Redirect Plunderstorm to home
		if classSlug == "plunderstorm" {
			c.Redirect(http.StatusMovedPermanently, "/")
			return
		}

		classes, err := getAllClasses(db)
		if err != nil {
			c.String(http.StatusInternalServerError, "getAllClasses Error: %v", err)
			return
		}

		// Get channel
		var channelID, channelName string
		var socialsJSON []byte
		err = db.QueryRow(
			"SELECT id, display_name, socials FROM channels WHERE twitch_login = $1 AND is_published = true",
			channelSlug,
		).Scan(&channelID, &channelName, &socialsJSON)
		if err == sql.ErrNoRows {
			c.HTML(http.StatusNotFound, "404", gin.H{})
			return
		}
		if err != nil {
			c.String(http.StatusInternalServerError, "GetChannel Error: %v", err)
			return
		}

		var socials map[string]string
		json.Unmarshal(socialsJSON, &socials)

		// Get class info
		var className string
		db.QueryRow(
			"SELECT class_name FROM classes WHERE game_slug = 'wow' AND class_slug = $1",
			classSlug,
		).Scan(&className)

		// Get all channels for this class (for navigation)
		var classID int64
		db.QueryRow(
			"SELECT id FROM classes WHERE game_slug = 'wow' AND class_slug = $1",
			classSlug,
		).Scan(&classID)

		rows, err := db.Query(`
			SELECT DISTINCT
				c.id, c.twitch_login, c.display_name, c.socials, c.sort_weight
			FROM channels c
			INNER JOIN channel_specs cs ON c.id = cs.channel_id
			INNER JOIN specs s ON cs.spec_id = s.id
			WHERE s.class_id = $1
			  AND c.is_published = true
			ORDER BY c.sort_weight DESC, c.display_name ASC
			LIMIT 100
		`, classID)
		if err != nil {
			c.String(http.StatusInternalServerError, "QueryChannels Error: %v", err)
			return
		}
		defer rows.Close()

		var channels []map[string]interface{}
		var logins []string
		for rows.Next() {
			var id, login, name string
			var socialsJSON []byte
			var sortWeight int
			if err := rows.Scan(&id, &login, &name, &socialsJSON, &sortWeight); err != nil {
				continue
			}
			var socials map[string]string
			json.Unmarshal(socialsJSON, &socials)
			channels = append(channels, map[string]interface{}{
				"id":       id,
				"slug":     login,
				"name":     name,
				"socials":  socials,
			})
			logins = append(logins, login)
		}

		// Fetch Twitch stream data for all channels
		streamDataMap := make(map[string]interface{})
		if len(logins) > 0 && twitchClient != nil {
			ctx := c.Request.Context()
			resp, err := twitchClient.GetStreamsByLogin(ctx, logins)
			if err == nil {
				for _, stream := range resp.Data {
					streamDataMap[strings.ToLower(stream.UserLogin)] = map[string]interface{}{
						"online":      true,
						"title":       stream.Title,
						"viewers":     stream.ViewerCount,
						"game":        stream.GameName,
					}
				}
			}
		}

		// Mark channels as online/offline
		for i := range channels {
			slug := channels[i]["slug"].(string)
			if data, ok := streamDataMap[strings.ToLower(slug)]; ok {
				channels[i]["stream"] = data
			} else {
				channels[i]["stream"] = map[string]interface{}{"online": false}
			}
		}

		// Sort channels: Online first -> Viewers (desc) -> Name (asc)
		sort.SliceStable(channels, func(i, j int) bool {
			streamI := channels[i]["stream"].(map[string]interface{})
			streamJ := channels[j]["stream"].(map[string]interface{})

			onlineI, _ := streamI["online"].(bool)
			onlineJ, _ := streamJ["online"].(bool)

			if onlineI != onlineJ {
				return onlineI
			}

			if onlineI {
				viewersI, _ := streamI["viewers"].(int)
				viewersJ, _ := streamJ["viewers"].(int)
				if viewersI != viewersJ {
					return viewersI > viewersJ
				}
			}

			return false
		})

		// Get current channel stream status
		var streamData map[string]interface{}
		if data, ok := streamDataMap[strings.ToLower(channelSlug)]; ok {
			streamData = data.(map[string]interface{})
		} else if twitchClient != nil {
			ctx := c.Request.Context()
			resp, err := twitchClient.GetStreamsByLogin(ctx, []string{channelSlug})
			if err == nil && len(resp.Data) > 0 {
				s := resp.Data[0]
				streamData = map[string]interface{}{
					"online":    true,
					"title":     s.Title,
					"viewers":   s.ViewerCount,
					"game":      s.GameName,
				}
			} else {
				streamData = map[string]interface{}{"online": false}
			}
		} else {
			streamData = map[string]interface{}{"online": false}
		}

		canonical := canonicalURL(c)
		title := fmt.Sprintf("%s – %s PvP – PvPtv", channelName, className)
		desc := fmt.Sprintf("Watch %s play %s PvP live on PvPtv", channelName, className)

		var jsonLD []interface{}
		person := map[string]interface{}{
			"@context":    "https://schema.org",
			"@type":       "Person",
			"name":        channelName,
			"url":         canonical,
			"image":       "/assets/logo.png",
			"description": desc,
		}
		jsonLD = append(jsonLD, person)

		if online, _ := streamData["online"].(bool); online {
			event := map[string]interface{}{
				"@context":  "https://schema.org",
				"@type":     "BroadcastEvent",
				"name":      streamData["title"],
				"startDate": time.Now().Format(time.RFC3339),
				"location": map[string]interface{}{
					"@type": "VirtualLocation",
					"url":   canonical,
				},
				"performer": person,
			}
			jsonLD = append(jsonLD, event)
		}

		c.HTML(http.StatusOK, "stream", gin.H{
			"Canonical":       canonical,
			"OGImage":         "/assets/logo.png",
			"ClassSlug":       classSlug,
			"ClassName":       className,
			"ChannelSlug":     channelSlug,
			"ChannelName":     channelName,
			"Channels":        channels,
			"Socials":         socials,
			"Stream":          streamData,
			"Classes":         classes,
			"Title":           title,
			"MetaDescription": desc,
			"JSONLD":          jsonLD,
		})
	}
}

// robotsHandler serves a basic robots.txt
func robotsHandler(c *gin.Context) {
	scheme := "https"
	if c.Request.TLS == nil && c.Request.Header.Get("X-Forwarded-Proto") != "https" {
		scheme = "http"
	}
	host := c.Request.Host
	if host == "" {
		host = "pvptv.hnts.dev"
	}

	sitemapURL := fmt.Sprintf("%s://%s/sitemap.xml", scheme, host)
	c.Header("Content-Type", "text/plain; charset=utf-8")
	c.String(http.StatusOK, fmt.Sprintf("User-agent: *\nDisallow: /admin\n\nSitemap: %s\n", sitemapURL))
}

// sitemapHandler returns a handler that will generate sitemap.xml from DB.
func sitemapHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/xml; charset=utf-8")

		scheme := "https"
		if c.Request.TLS == nil && c.Request.Header.Get("X-Forwarded-Proto") != "https" {
			scheme = "http"
		}
		host := c.Request.Host
		if host == "" {
			host = "pvptv.hnts.dev"
		}
		baseURL := fmt.Sprintf("%s://%s", scheme, host)

		type URL struct {
			Loc        string
			LastMod    string
			ChangeFreq string
			Priority   string
		}
		var urls []URL

		// Home
		urls = append(urls, URL{Loc: baseURL + "/", ChangeFreq: "daily", Priority: "1.0"})

		// Classes
		rows, err := db.Query("SELECT class_slug FROM classes WHERE game_slug = 'wow'")
		if err == nil {
			defer rows.Close()
			for rows.Next() {
				var slug string
				if err := rows.Scan(&slug); err == nil {
					urls = append(urls, URL{
						Loc:        fmt.Sprintf("%s/%s", baseURL, slug),
						ChangeFreq: "daily",
						Priority:   "0.8",
					})
				}
			}
		}

		// Channels
		// Only include channels that have a primary spec assigned
		rows, err = db.Query(`
			SELECT c.twitch_login, cl.class_slug, c.updated_at
			FROM channels c
			JOIN channel_specs cs ON c.id = cs.channel_id
			JOIN specs s ON cs.spec_id = s.id
			JOIN classes cl ON s.class_id = cl.id
			WHERE c.is_published = true AND cs.is_primary = true
		`)
		if err == nil {
			defer rows.Close()
			for rows.Next() {
				var login, classSlug string
				var updatedAt time.Time
				if err := rows.Scan(&login, &classSlug, &updatedAt); err == nil {
					urls = append(urls, URL{
						Loc:        fmt.Sprintf("%s/%s/%s", baseURL, classSlug, login),
						LastMod:    updatedAt.Format("2006-01-02"),
						ChangeFreq: "weekly",
						Priority:   "0.6",
					})
				}
			}
		}

		c.String(http.StatusOK, `<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
`)
		for _, u := range urls {
			c.String(http.StatusOK, fmt.Sprintf(`  <url>
    <loc>%s</loc>
    <changefreq>%s</changefreq>
    <priority>%s</priority>`, u.Loc, u.ChangeFreq, u.Priority))
			if u.LastMod != "" {
				c.String(http.StatusOK, fmt.Sprintf("\n    <lastmod>%s</lastmod>", u.LastMod))
			}
			c.String(http.StatusOK, "\n  </url>\n")
		}
		c.String(http.StatusOK, `</urlset>`)
	}
}

// canonicalURL builds an absolute canonical URL for the current request.
func canonicalURL(c *gin.Context) string {
	scheme := c.Request.Header.Get("X-Forwarded-Proto")
	if scheme == "" {
		if c.Request.TLS != nil {
			scheme = "https"
		} else {
			scheme = "http"
		}
	}
	host := c.Request.Host
	if host == "" {
		host = "pvptv.hnts.dev"
	}
	return scheme + "://" + host + c.Request.URL.Path
}

// renderFormWithError is a helper to render the suggestion form with an error message
func renderFormWithError(c *gin.Context, db *sql.DB, errorMsg string) {
	sidebarClasses, _ := getAllClasses(db)
	formClasses, _ := getFormClasses(db)
	c.HTML(http.StatusOK, "suggest_streamer", gin.H{
		"Canonical":   canonicalURL(c),
		"Title":       "Suggest a Streamer - PvPtv",
		"Classes":     sidebarClasses,
		"FormClasses": formClasses,
		"Error":       errorMsg,
	})
}

// getFormClasses fetches classes with their specs for the suggestion form
func getFormClasses(db *sql.DB) ([]interface{}, error) {
	rows, err := db.Query(`
		SELECT c.id, c.class_name, s.id, s.spec_name
		FROM classes c
		LEFT JOIN specs s ON c.id = s.class_id
		WHERE c.game_slug = 'wow'
		ORDER BY c.class_name, s.spec_name
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	type SpecView struct {
		ID   string
		Name string
	}
	type ClassView struct {
		ID    string
		Name  string
		Specs []SpecView
	}

	var classes []*ClassView
	classMap := make(map[string]*ClassView)

	for rows.Next() {
		var classID, className, specID, specName sql.NullString
		rows.Scan(&classID, &className, &specID, &specName)

		if !classID.Valid {
			continue
		}

		class, exists := classMap[classID.String]
		if !exists {
			class = &ClassView{
				ID:    classID.String,
				Name:  className.String,
				Specs: []SpecView{},
			}
			classMap[classID.String] = class
			classes = append(classes, class)
		}

		if specID.Valid && specName.Valid {
			class.Specs = append(class.Specs, SpecView{
				ID:   specID.String,
				Name: specName.String,
			})
		}
	}

	// Convert to interface{} slice for template
	result := make([]interface{}, len(classes))
	for i, c := range classes {
		result[i] = c
	}
	return result, nil
}

// suggestStreamerForm renders the public suggestion form
func suggestStreamerForm(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		formClasses, err := getFormClasses(db)
		if err != nil {
			c.String(http.StatusInternalServerError, "DB Error: %v", err)
			return
		}
		
		// Get Sidebar Classes (for the sidebar template)
		sidebarClasses, _ := getAllClasses(db)

		c.HTML(http.StatusOK, "suggest_streamer", gin.H{
			"Canonical":   canonicalURL(c),
			"Title":       "Suggest a Streamer - PvPtv",
			"Classes":     sidebarClasses, // For sidebar (standard structure)
			"FormClasses": formClasses,    // For form (complex structure with specs)
			"Success":     c.Query("success") == "true",
		})
	}
}

// suggestStreamerSubmit handles the public suggestion form submission
func suggestStreamerSubmit(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		displayName := c.PostForm("display_name")
		twitchLogin := c.PostForm("twitch_login")
		specIDs := c.PostFormArray("spec_ids")

		// Basic Validation
		if displayName == "" || twitchLogin == "" || len(specIDs) == 0 {
			renderFormWithError(c, db, "Missing required fields (Name, Twitch URL, Specs)")
			return
		}

		// Cleanup twitch login (remove url parts if pasted)
		twitchLogin = strings.TrimSuffix(twitchLogin, "/")
		parts := strings.Split(twitchLogin, "/")
		twitchLogin = parts[len(parts)-1]

		// Parse socials
		socials := make(map[string]string)
		socialKeys := []string{"twitter", "youtube", "discord", "instagram", "website"}
		for _, key := range socialKeys {
			if val := c.PostForm("social_" + key); val != "" {
				socials[key] = val
			}
		}
		socialsJSON, _ := json.Marshal(socials)

		// Insert Channel as unpublished (pending)
		var id string
		err := db.QueryRow(`
			INSERT INTO channels (display_name, twitch_login, is_published, socials)
			VALUES ($1, $2, false, $3)
			RETURNING id
		`, displayName, twitchLogin, socialsJSON).Scan(&id)
		
		if err != nil {
			// Check for duplicate key error
			if strings.Contains(err.Error(), "unique constraint") {
				renderFormWithError(c, db, "Streamer already referenced, or is in validation process")
				return
			}
			// Re-render form with generic error
			renderFormWithError(c, db, "An error occurred. Please try again later.")
			return
		}

		// Insert Specs
		tx, err := db.Begin()
		if err != nil {
			renderFormWithError(c, db, "An error occurred. Please try again later.")
			return
		}
		
		stmt, err := tx.Prepare("INSERT INTO channel_specs (channel_id, spec_id, is_primary) VALUES ($1, $2, false)")
		if err != nil {
			tx.Rollback()
			renderFormWithError(c, db, "An error occurred. Please try again later.")
			return
		}
		defer stmt.Close()

		for _, specID := range specIDs {
			// No primary spec for suggestions, admin decides.
			if _, err := stmt.Exec(id, specID); err != nil {
				tx.Rollback()
				renderFormWithError(c, db, "An error occurred. Please try again later.")
				return
			}
		}

		if err := tx.Commit(); err != nil {
			renderFormWithError(c, db, "An error occurred. Please try again later.")
			return
		}

		// Re-render form with success message
		// We need to fetch classes again for the form re-render
		// For simplicity, just redirect to same page with query param?
		// Or reuse the form handler logic.
		// Redirect is cleaner.
		c.Redirect(http.StatusFound, "/suggest-streamer?success=true")
	}
}
