package api

import (
	"api_crud/config"
	db "api_crud/internal"
	sqlc "api_crud/internal"

	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/robfig/cron/v3"

	"log"
)

type Store interface {
	db.Querier
	ExecTx(ctx context.Context, fn func(*sqlc.Queries) error) error
}

type SQLStore struct {
	*sqlc.Queries
	db *sql.DB
}

type Server struct {
	Store    Store
	Config   *config.API
	cronTask *cron.Cron
}

// NewStore create new Store
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: sqlc.New(db),
	}
}

// execTx executes a function within a database transaction
func (store *SQLStore) ExecTx(ctx context.Context, fn func(*sqlc.Queries) error) error {
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

func NewServer() *Server {
	cnf := config.New()
	source := fmt.Sprintf("postgresql://%s:%s@%s:%v/%s?sslmode=disable", cnf.Database.User, cnf.Database.Password, cnf.Database.Host, cnf.Database.Port, cnf.Database.Database)
	pg, err := sql.Open("postgres", source)
	if err != nil {
		log.Fatal("Err DB ==> ", err)
	}

	if err = pg.Ping(); err != nil {
		fmt.Printf("Postgres ping error : (%v)", err)
	}

	store := NewStore(pg)
	server := &Server{Store: store}
	server.Config = cnf
	server.runCron(&server.cronTask, server.Config)
	initCron(server)
	return server
}

func initCron(server *Server) {
	c := cron.New()
	c.AddFunc("@hourly", func() { server.Store.DeleteOldRefreshToken(context.Background()) })
	c.Start()
}

// storeRefresh store refres_token into database
func (server *Server) StoreRefresh(ctx context.Context, token string, exp time.Time, userID uuid.UUID) error {
	return server.Store.CreateRefreshToken(ctx, db.CreateRefreshTokenParams{
		Token:   token,
		ExpirOn: exp,
		UserID:  userID,
	})
}
