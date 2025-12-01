package db

import "errors"

var (
	// ErrMissingDSN is returned when DATABASE_URL is not set.
	ErrMissingDSN = errors.New("DATABASE_URL is not set")
)


