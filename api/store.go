package api

import (
	"api_crud/config"
	sqlc "api_crud/internal/db"
	"context"
	"database/sql"
	"fmt"

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
	Store  *Store
	Config *config.API
	Oauth  *oserver.Server
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

	server.Oauth = server.oAuth()

	return server
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
