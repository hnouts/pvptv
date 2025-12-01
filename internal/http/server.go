package http

import (
	"database/sql"
	"encoding/json"
	"html/template"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"arenatv/internal/twitch"
)

// NewServer constructs the Gin engine with all middleware and routes wired.
// It is used by cmd/server.
func NewServer(db *sql.DB) *gin.Engine {
	ginMode := envOrDefault("GIN_MODE", gin.ReleaseMode)
	gin.SetMode(ginMode)

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	// Static assets (reuse existing CSS/images from docs/web).
	r.Static("/assets", "./docs/web")

	// Register templates (public + admin will live here).
	r.SetFuncMap(template.FuncMap{
		"nowYear": func() int { return time.Now().Year() },
		"safeHTML": func(s string) template.HTML { return template.HTML(s) },
		"json": func(v interface{}) (template.JS, error) {
			a, err := json.Marshal(v)
			if err != nil {
				return "", err
			}
			return template.JS(a), nil
		},
	})
	r.LoadHTMLGlob("templates/*.html")

	// Health endpoints.
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	r.GET("/readyz", func(c *gin.Context) {
		// Later: check db.PingContext, Twitch client health, etc.
		if err := db.Ping(); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": "db-unavailable"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "ready"})
	})

	// SEO-related endpoints (will be filled out properly later).
	r.GET("/robots.txt", robotsHandler)
	r.GET("/sitemap.xml", sitemapHandler(db))

	// Initialize Twitch client (non-fatal if missing credentials)
	twitchClient, _ := twitch.NewHelixClient()

	// Public site routes.
	r.GET("/", homeHandler(db))
	r.GET("/:class", classHandler(db, twitchClient))
	r.GET("/:class/:channel", streamHandler(db, twitchClient))
	
	// Redirect Plunderstorm routes
	r.GET("/plunderstorm/:channel", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	// Admin routes.
	registerAdminRoutes(r, db)

	return r
}

func envOrDefault(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}


