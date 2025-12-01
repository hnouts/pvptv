package db

import (
	"database/sql"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// Open returns a *sql.DB connected to the Postgres database specified by
// DATABASE_URL. It is intended to be called once at startup and shared.
func Open() (*sql.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return nil, ErrMissingDSN
	}
	return sql.Open("pgx", dsn)
}


