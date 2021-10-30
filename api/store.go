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
	oerrors "gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	oserver "gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"

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
	source := fmt.Sprintf("user=%s password=%s host=%s port=%v dbname=%s sslmode=disable TimeZone=%s", cnf.Database.User, cnf.Database.Password, cnf.Database.Host, cnf.Database.Port, cnf.Database.Database, cnf.TZ)
	pg, err := sql.Open("postgres", source)
	if err != nil {
		log.Fatal(err)
	}
	store := NewStore(pg)

	server := &Server{Store: store}
	server.Config = cnf
	server.Oauth = server.oAuth()
	server.runCron(&server.cronTask, server.Config)
	initCron(server)
	return server
}

func initCron(server *Server) {
	c := cron.New()
	c.AddFunc("@hourly", func() { server.Store.DeleteOldRefreshToken(context.Background()) })
	c.Start()
}

// oAuth manage and store oAuth token
func (server *Server) oAuth() *oserver.Server {
	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	// token store
	manager.MustTokenStorage(store.NewFileTokenStore("data.db"))

	// client store
	clientStore := store.NewClientStore()
	manager.MapClientStorage(clientStore)

	srv := oserver.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(false)
	srv.SetClientInfoHandler(oserver.ClientFormHandler)
	manager.SetRefreshTokenCfg(manage.DefaultRefreshTokenCfg)

	srv.SetInternalErrorHandler(func(err error) (re *oerrors.Response) {
		log.Println("Internal Error:", err.Error())
		return re
	})

	srv.SetResponseErrorHandler(func(re *oerrors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	// client domain
	var domain string
	if server.Config.SSL {
		domain = fmt.Sprintf("https://%s", server.Config.Host)
	} else {
		domain = fmt.Sprintf("http://%s", server.Config.Host)
	}

	// client store
	err := clientStore.Set(server.Config.Security.OAuthID, &models.Client{
		ID:     server.Config.Security.OAuthID,
		Secret: server.Config.Security.OAuthSecret,
		Domain: domain,
	})

	if err != nil {
		log.Println("clientStore", err)
	}

	return srv
}

// storeRefresh store refres_token into database
func (server *Server) StoreRefresh(ctx context.Context, token string, exp time.Time, userID uuid.UUID) error {
	return server.Store.CreateRefreshToken(ctx, db.CreateRefreshTokenParams{
		Ip:        "À COMPLÉTER",
		UserAgent: "À COMPLÉTER",
		Token:     token,
		ExpirOn:   exp,
		UserID:    userID,
	})
}
