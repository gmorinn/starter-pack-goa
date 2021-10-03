package api

import (
	sqlc "api_crud/internal/db"
	"database/sql"
)

type Store struct {
	*sqlc.Queries
	db *sql.DB
}

// NewStore create new Store
func NewStore(db *sql.DB) *Store {
	// db.SetMaxOpenConns(140)
	// db.SetMaxIdleConns(140)
	return &Store{
		db:      db,
		Queries: sqlc.New(db),
	}
}
