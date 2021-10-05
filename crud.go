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

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

// crud service example implementation.
// The example methods log the requests and return zero values.
type crudsrvc struct {
	logger *log.Logger
	store  *api.Store
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
	return &crudsrvc{logger, store}
}

// Read Book
func (s *crudsrvc) GetBook(ctx context.Context, p *crud.GetBookPayload) (res *crud.GetBookResult, err error) {
	book, err := s.store.GetBook(ctx, uuid.MustParse(p.ID))
	if err != nil {
		return nil, &crud.IDDoesntExist{
			Message: "ERROR_GET_BOOK: " + err.Error(),
			ID:      p.ID,
			Success: false,
		}
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
	if _, err := s.store.GetBook(ctx, uuid.MustParse(p.ID)); err != nil {
		return nil, &crud.IDDoesntExist{
			Message: "ERROR_GET_BOOK: " + err.Error(),
			ID:      p.ID,
			Success: false,
		}
	}

	if err := s.store.DeleteBook(ctx, uuid.MustParse(p.ID)); err != nil {
		return nil, &crud.UnknownError{
			Message: "ERROR_DELETE_BOOK: " + err.Error(),
			Success: false,
		}
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
		return nil, &crud.UnknownError{
			Message: "ERROR_CREATE_BOOK: " + err.Error(),
			Success: false,
		}
	}

	newBook, err := s.GetBook(ctx, &crud.GetBookPayload{ID: book.ID.String()})
	if err != nil {
		return nil, &crud.UnknownError{
			Message: "ERROR_GET_BOOK: " + err.Error(),
			Success: false,
		}
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
		return nil, &crud.UnknownError{
			Success: false,
			Message: "ERROR_GET_ALL_BOOK: " + err.Error(),
		}
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
		return nil, &crud.UnknownError{
			Success: false,
			Message: "ERROR_UPDATE_BOOK: " + err.Error(),
		}
	}

	newBook, err := s.store.GetBook(ctx, uuid.MustParse(p.ID))
	if err != nil {
		return nil, &crud.IDDoesntExist{
			Message: "ERROR_GET_BOOK: " + err.Error(),
			ID:      p.ID,
			Success: false,
		}
	}

	response := crud.UpdateBookResult{
		ID:      newBook.ID.String(),
		Name:    newBook.Name,
		Price:   newBook.Price,
		Success: true,
	}
	return &response, nil
}
