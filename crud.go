package basic

import (
	"api_crud/api"
	crud "api_crud/gen/crud"
	"api_crud/internal/db"
	"context"
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"goa.design/goa/v3/security"
)

type crudsrvc struct {
	logger *log.Logger
	server *api.Server
}

func NewCrud(logger *log.Logger, server *api.Server) crud.Service {
	return &crudsrvc{logger, server}
}

var (
	ErrInvalidToken error = crud.Unauthorized("invalid token")

	ErrInvalidTokenScopes error = crud.InvalidScopes("invalid scopes in token")
)

func Error_ID(msg, id string, err error) *crud.IDDoesntExist {
	return &crud.IDDoesntExist{
		Err: err.Error(),
		ID:  id,
	}
}

func ErrorResponse(msg string, err error) *crud.UnknownError {
	return &crud.UnknownError{
		Err:       err.Error(),
		ErrorCode: msg,
	}
}

func (s *crudsrvc) JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error) {

	claims := make(jwt.MapClaims)

	// authorize request
	// 1. parse JWT token, token key is hardcoded to "secret" in this example
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		b := ([]byte(s.server.Config.Security.Secret))
		return b, nil
	})
	if err != nil {
		return ctx, ErrInvalidToken
	}

	// 2. validate provided "scopes" claim
	if claims["scopes"] == nil {
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
		return ctx, crud.InvalidScopes(err.Error())
	}

	// 3. add authInfo to context
	ctx = contextWithAuthInfo(ctx, authInfo{
		claims: claims,
	})
	return ctx, nil
}

// Read Book
func (s *crudsrvc) GetBook(ctx context.Context, p *crud.GetBookPayload) (res *crud.GetBookResult, err error) {
	book, err := s.server.Store.GetBook(ctx, uuid.MustParse(p.ID))
	if err != nil {
		return nil, Error_ID("ERROR_GET_BOOK", p.ID, err)
	}
	response := crud.GetBookResult{
		ID:      book.ID.String(),
		Name:    book.Name,
		Price:   book.Price,
		Success: true,
	}
	return &response, nil
}

// Delete Book
func (s *crudsrvc) DeleteBook(ctx context.Context, p *crud.DeleteBookPayload) (res *crud.DeleteBookResult, err error) {
	if err := s.server.Store.DeleteBook(ctx, uuid.MustParse(p.ID)); err != nil {
		return nil, ErrorResponse("ERROR_DELETE_BOOK", err)
	}
	return &crud.DeleteBookResult{Success: true}, nil
}

// Create Book
func (s *crudsrvc) CreateBook(ctx context.Context, p *crud.CreateBookPayload) (res *crud.CreateBookResult, err error) {
	arg := db.CreateBookParams{
		Price: p.Price,
		Name:  p.Name,
	}
	book, err := s.server.Store.CreateBook(ctx, arg)
	if err != nil {
		return nil, ErrorResponse("ERROR_CREATE_BOOK", err)
	}

	newBook, err := s.GetBook(ctx, &crud.GetBookPayload{ID: book.ID.String()})
	if err != nil {
		return nil, Error_ID("ERROR_GET_BOOK", book.ID.String(), err)
	}
	response := crud.CreateBookResult{
		Book: &crud.BookResponse{
			ID:    newBook.ID,
			Name:  newBook.Name,
			Price: newBook.Price,
		},
		Success: true,
	}

	return &response, nil
}

// Read Books
func (s *crudsrvc) GetAllBooks(ctx context.Context) (res *crud.GetAllBooksResult, err error) {
	books, err := s.server.Store.GetBooks(ctx)
	if err != nil {
		return nil, ErrorResponse("ERROR_GET_ALL_BOOKS", err)

	}

	var BookResponse []*crud.BookResponse
	for _, v := range books {
		id := v.ID.String()
		BookResponse = append(BookResponse, &crud.BookResponse{
			ID:    id,
			Name:  v.Name,
			Price: v.Price,
		})
	}

	response := crud.GetAllBooksResult{
		Books:   BookResponse,
		Success: true,
	}
	return &response, nil
}

func (s *crudsrvc) UpdateBook(ctx context.Context, p *crud.UpdateBookPayload) (res *crud.UpdateBookResult, err error) {
	arg := db.UpdateBookParams{
		ID:    uuid.MustParse(p.ID),
		Name:  p.Name,
		Price: p.Price,
	}
	if err := s.server.Store.UpdateBook(ctx, arg); err != nil {
		return nil, ErrorResponse("ERROR_UPDATE_BOOK", err)

	}
	newBook, err := s.server.Store.GetBook(ctx, uuid.MustParse(p.ID))
	if err != nil {
		return nil, Error_ID("ERROR_GET_BOOK", p.ID, err)

	}
	response := crud.UpdateBookResult{
		ID:      newBook.ID.String(),
		Name:    newBook.Name,
		Price:   newBook.Price,
		Success: true,
	}
	return &response, nil
}

func (s *crudsrvc) OAuth(ctx context.Context, p *crud.OAuthPayload) (res *crud.OAuthResponse, err error) {
	if p.GrantType != "grant_type" {
		return nil, ErrorResponse("OAUTH", nil)
	}
	if p.ClientID != s.server.Config.Security.OAuthID {

	}
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"access_token": time.Now().Add(time.Duration((time.Hour * 24) * time.Duration(s.server.Config.Security.AccessTokenDuration))).Unix(),
	// 	"expires_in":   time.Now().Add(time.Duration(time.Hour * 2)).Unix(),
	// 	"token_type":   "Bearer",
	// })
	return nil, nil
}
