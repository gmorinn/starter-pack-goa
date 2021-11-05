package api

import (
	"api_crud/config"
	jwttoken "api_crud/gen/jwt_token"
	db "api_crud/internal"
	sqlc "api_crud/internal"

	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/robfig/cron/v3"
	"goa.design/goa/v3/security"
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

func (server *Server) CheckAuth(ctx context.Context, token string, scheme *security.OAuth2Scheme) (context.Context, error) {

	claims := make(jwt.MapClaims)

	// authorize request
	// 1. parse JWT token, token key is hardcoded to "secret" in this example
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		b := ([]byte(server.Config.Security.Secret))
		return b, nil
	})
	if err != nil {
		return ctx, ErrInvalidToken
	}

	// 2. validate provided "scopes" claim
	if claims["scopes"] == nil {
		return ctx, ErrInvalidTokenScopes
	}
	if claims["expires_in"] == nil {
		return ctx, ErrInvalidTokenScopes
	}
	if claims["token_type"] != "Bearer" {
		return ctx, ErrInvalidToken
	}
	scopes, ok := claims["scopes"].([]interface{})
	if !ok {
		return ctx, ErrInvalidTokenScopes
	}
	scopesInToken := make([]string, len(scopes))
	for _, scp := range scopes {
		scopesInToken = append(scopesInToken, scp.(string))
	}
	if err := scheme.Validate(scopesInToken); err != nil {
		return ctx, jwttoken.InvalidScopes(err.Error())
	}

	// 3. add authInfo to context
	ctx = contextWithAuthInfo(ctx, authInfo{
		oAuth: claims,
	})
	return ctx, nil
}

func (server *Server) CheckJWT(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error) {

	claims := make(jwt.MapClaims)

	// authorize request
	// 1. parse JWT token, token key is hardcoded to "secret" in this example
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		b := ([]byte(server.Config.Security.Secret))
		return b, nil
	})
	if err != nil {
		return ctx, ErrInvalidToken
	}

	// 2. validate provided "scopes" claim
	if claims["scopes"] == nil {
		return ctx, ErrInvalidTokenScopes
	}
	if claims["id"] == nil {
		return ctx, ErrInvalidTokenScopes
	}
	if claims["exp"] == nil {
		return ctx, ErrInvalidTokenScopes
	}
	scopes, ok := claims["scopes"].([]interface{})
	if !ok {
		return ctx, ErrInvalidTokenScopes
	}
	scopesInToken := make([]string, len(scopes))
	for _, scp := range scopes {
		scopesInToken = append(scopesInToken, scp.(string))
	}
	if err := schema.Validate(scopesInToken); err != nil {
		return ctx, jwttoken.InvalidScopes(err.Error())
	}

	// 3. add authInfo to context
	ctx = contextWithAuthInfo(ctx, authInfo{
		jwtToken: claims,
	})
	return ctx, nil
}
