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
	oserver "gopkg.in/oauth2.v3/server"

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
	Store    *Store
	Config   *config.API
	Oauth    *oserver.Server
	cronTask *cron.Cron
}

func NewServer() *Server {
	cnf := config.New()
	source := fmt.Sprintf("user=%s password=%s host=%s port=%v dbname=%s sslmode=disable", cnf.Database.User, cnf.Database.Password, cnf.Database.Host, cnf.Database.Port, cnf.Database.Database)
	pg, err := sql.Open("postgres", source)
	if err != nil {
		log.Fatal(err)
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
