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

func Error_ID(msg, id string, err error) *crud.IDDoesntExist {
	return &crud.IDDoesntExist{
		Message: msg + " " + err.Error(),
		ID:      id,
	}
}

func ErrorResponse(msg string, err error) *crud.UnknownError {
	return &crud.UnknownError{
		Message: msg + " " + err.Error(),
	}
}

func ErrorEmail() *crud.UnknownError {
	return &crud.UnknownError{
		Message: "EMAIL_ALREADY_EXIST",
	}
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

func (s *crudsrvc) Signup(ctx context.Context, p *crud.SignupPayload) (res *crud.Register, err error) {

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

	claims := make(jwt.MapClaims)
	claims["id"] = user.ID.String()
	claims["exp"] = time.Now().Add(time.Duration((time.Hour * 24) * time.Duration(s.config.Security.AccessTokenDuration))).Unix()
	t, err := generateJWT(claims, s.config.Security.Secret)
	if err != nil {
		return nil, ErrorResponse("ERROR_GENERATE_ACCESS_JWT", err)
	}

	expt := time.Now().Add(time.Duration((time.Hour * 24) * time.Duration(s.config.Security.RefreshTokenDuration)))
	exp := expt.Unix()
	claims["exp"] = exp
	r, err := generateJWT(claims, s.config.Security.Secret)
	if err != nil {
		return nil, ErrorResponse("ERROR_GENERATE_REFRESH_JWT", err)
	}

	response := crud.Register{
		AccessToken:  t,
		RefreshToken: r,
		Success:      true,
	}
	return &response, nil
}
