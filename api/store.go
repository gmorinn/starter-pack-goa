package api

import (
	"api_crud/config"
	sqlc "api_crud/internal/db"
	"context"
	"database/sql"
	"fmt"

	"log"
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

// execTx executes a function within a database transaction
func (store *Store) ExecTx(ctx context.Context, fn func(*sqlc.Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := sqlc.New(tx)
	if err := fn(q); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx: err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

type Server struct {
	Store  *Store
	Config *config.API
}

func NewServer() *Server {
	cnf := config.New()
	source := fmt.Sprintf("user=%s password=%s host=%s port=%v dbname=%s sslmode=disable TimeZone=%s", cnf.Database.User, cnf.Database.Password, cnf.Database.Host, cnf.Database.Port, cnf.Database.Database, cnf.TZ)
	pg, err := sql.Open("postgres", source)
	if err != nil {
		log.Fatal(err)
	}
	store := NewStore(pg)

	server := Server{
		Store:  store,
		Config: cnf,
	}

	return &server
}
