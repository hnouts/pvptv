package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"arenatv/internal/db"
	httpserver "arenatv/internal/http"
)

func main() {
	// Best-effort: load .env for local development if present.
	loadDotEnv(".env")

	// Open DB early so readiness can depend on it later.
	database, err := db.Open()
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer database.Close()

	srv := httpserver.NewServer(database)

	addr := ":" + httpserverEnvOrDefault("PORT", "8080")
	log.Printf("starting pvptv v2 server on %s", addr)

	if err := srv.Run(addr); err != nil {
		log.Fatalf("server exited: %v", err)
	}
}

// httpserverEnvOrDefault is a tiny adapter to avoid re-exporting config helpers from internal/http.
func httpserverEnvOrDefault(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

// loadDotEnv loads KEY=VALUE pairs from a simple .env file into the process
// environment. It is intended for local development only.
func loadDotEnv(path string) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])
		// Don't override already-set env vars (e.g., from Docker).
		if _, exists := os.LookupEnv(key); !exists && key != "" {
			_ = os.Setenv(key, val)
		}
	}
}



