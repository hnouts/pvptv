package http

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"arenatv/internal/twitch"
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
func homeHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		classes, err := getAllClasses(db)
		if err != nil {
			c.String(http.StatusInternalServerError, "DB Error: %v", err)
			return
		}

		canonical := canonicalURL(c)
		jsonLD := map[string]interface{}{
			"@context": "https://schema.org",
			"@type":    "WebSite",
			"name":     "PvPtv.io",
			"url":      "https://pvptv.hnts.dev/",
		}

		c.HTML(http.StatusOK, "home", gin.H{
			"Canonical":       canonical,
			"OGImage":         "/assets/logo.png",
			"Classes":         classes,
			"Title":           "PvPtv.io – WoW PvP streams",
			"MetaDescription": "All of WoW PvP streaming in one place – discover top arena streamers by class and spec.",
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
		title := fmt.Sprintf("%s – PvPtv.io", className)
		desc := fmt.Sprintf("Watch top %s WoW PvP streamers live on PvPtv.io", className)

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
		title := fmt.Sprintf("%s – %s PvP – PvPtv.io", channelName, className)
		desc := fmt.Sprintf("Watch %s play %s PvP live on PvPtv.io", channelName, className)

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

// robotsHandler serves a basic robots.txt; will be extended as spec finalizes.
func robotsHandler(c *gin.Context) {
	c.Header("Content-Type", "text/plain; charset=utf-8")
	c.String(http.StatusOK, "User-agent: *\nDisallow: /admin\n\nSitemap: https://pvptv.hnts.dev/sitemap.xml\n")
}

// sitemapHandler returns a handler that will generate sitemap.xml from DB.
func sitemapHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/xml; charset=utf-8")

		type URL struct {
			Loc        string
			LastMod    string
			ChangeFreq string
			Priority   string
		}
		var urls []URL

		// Home
		urls = append(urls, URL{Loc: "https://pvptv.hnts.dev/", ChangeFreq: "daily", Priority: "1.0"})

		// Classes
		rows, err := db.Query("SELECT class_slug FROM classes WHERE game_slug = 'wow'")
		if err == nil {
			defer rows.Close()
			for rows.Next() {
				var slug string
				if err := rows.Scan(&slug); err == nil {
					urls = append(urls, URL{
						Loc:        fmt.Sprintf("https://pvptv.hnts.dev/%s", slug),
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
						Loc:        fmt.Sprintf("https://pvptv.hnts.dev/%s/%s", classSlug, login),
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
