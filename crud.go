package basic

import (
	"api_crud/api"
	"api_crud/config"
	crud "api_crud/gen/crud"
	"api_crud/internal/db"
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"goa.design/goa/v3/security"
)

// crud service example implementation.
// The example methods log the requests and return zero values.
type crudsrvc struct {
	logger *log.Logger
	store  *api.Store
	config *config.API
}

// NewCrud returns the crud service implementation.
func NewCrud(logger *log.Logger) crud.Service {
	cnf := config.Get()
	source := fmt.Sprintf("user=%s password=%s host=%s port=%v dbname=%s sslmode=disable TimeZone=%s", cnf.Database.User, cnf.Database.Password, cnf.Database.Host, cnf.Database.Port, cnf.Database.Database, cnf.TZ)
	pg, err := sql.Open("postgres", source)
	if err != nil {
		log.Fatal(err)
	}
	store := api.NewStore(pg)
	return &crudsrvc{logger, store, cnf}
}

var (
	// ErrUnauthorized is the error returned by Login when the request credentials
	// are invalid.
	ErrUnauthorized error = crud.Unauthorized("invalid username and password combination")

	// ErrInvalidToken is the error returned when the JWT token is invalid.
	ErrInvalidToken error = crud.Unauthorized("invalid token")

	// ErrInvalidTokenScopes is the error returned when the scopes provided in
	// the JWT token claims are invalid.
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

func ErrorEmail() *crud.EmailAlreadyExist {
	return &crud.EmailAlreadyExist{
		Message: "EMAIL_ALREADY_EXIST",
	}
}

func (s *crudsrvc) JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error) {

	claims := make(jwt.MapClaims)

	// authorize request
	// 1. parse JWT token, token key is hardcoded to "secret" in this example
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		b := ([]byte(s.config.Security.Secret))
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
	book, err := s.store.GetBook(ctx, uuid.MustParse(p.ID))
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
	if err := s.store.DeleteBook(ctx, uuid.MustParse(p.ID)); err != nil {
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
	book, err := s.store.CreateBook(ctx, arg)
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
	books, err := s.store.GetBooks(ctx)
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
	if err := s.store.UpdateBook(ctx, arg); err != nil {
		return nil, ErrorResponse("ERROR_UPDATE_BOOK", err)

	}
	newBook, err := s.store.GetBook(ctx, uuid.MustParse(p.ID))
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

func (s *crudsrvc) Signup(ctx context.Context, p *crud.SignupPayload) (res *crud.Sign, err error) {

	isExist, err := s.store.ExistUserByEmail(ctx, p.Email)
	if err != nil {
		return nil, ErrorResponse("ERROR_GET_MAIL", err)
	}
	if isExist {
		return nil, ErrorEmail()
	}

	arg := db.SignupParams{
		Firstname: p.Firstname,
		Lastname:  p.Lastname,
		Email:     p.Email,
		Crypt:     p.Password,
	}
	user, err := s.store.Signup(ctx, arg)
	if err != nil {
		return nil, ErrorResponse("ERROR_CREATE_USER", err)
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":     user.ID.String(),
		"exp":    time.Now().Add(time.Duration((time.Hour * 24) * time.Duration(s.config.Security.AccessTokenDuration))).Unix(),
		"scopes": []string{"api:read", "api:write"},
	})
	t, err := accessToken.SignedString(s.config.Security.Secret)
	if err != nil {
		return nil, ErrorResponse("ERROR_GENERATE_ACCESS_JWT", err)
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":     user.ID.String(),
		"exp":    time.Now().Add(time.Duration((time.Hour * 24) * time.Duration(s.config.Security.RefreshTokenDuration))).Unix(),
		"scopes": []string{"api:read", "api:write"},
	})
	r, err := refreshToken.SignedString(s.config.Security.Secret)
	if err != nil {
		return nil, ErrorResponse("ERROR_GENERATE_REFRESH_JWT", err)
	}

	response := crud.Sign{
		AccessToken:  t,
		RefreshToken: r,
		Success:      true,
	}
	return &response, nil
}

func (s *crudsrvc) Signin(ctx context.Context, p *crud.SigninPayload) (res *crud.Sign, err error) {
	// Request Login
	arg := db.LoginUserParams{
		Email: p.Email,
		Crypt: p.Password,
	}
	user, err := s.store.LoginUser(ctx, arg)
	if err != nil {
		return nil, ErrorResponse("ERROR_LOGIN_USER", err)
	}

	// Generate access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":     user.ID.String(),
		"exp":    time.Now().Add(time.Duration((time.Hour * 24) * time.Duration(s.config.Security.AccessTokenDuration))).Unix(),
		"scopes": []string{"api:read", "api:write"},
	})
	t, err := accessToken.SignedString([]byte(s.config.Security.Secret))
	if err != nil {
		return nil, ErrorResponse("ERROR_GENERATE_ACCESS_JWT", err)
	}

	// Generate refresh token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":     user.ID.String(),
		"exp":    time.Now().Add(time.Duration((time.Hour * 24) * time.Duration(s.config.Security.RefreshTokenDuration))).Unix(),
		"scopes": []string{"api:read", "api:write"},
	})
	r, err := refreshToken.SignedString([]byte(s.config.Security.Secret))
	if err != nil {
		return nil, ErrorResponse("ERROR_GENERATE_REFRESH_JWT", err)
	}

	response := crud.Sign{
		AccessToken:  t,
		RefreshToken: r,
		Success:      true,
	}
	return &response, nil
}
