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
func (s *crudsrvc) GetBook(ctx context.Context, p *crud.GetBookPayload) (res *crud.BookResponse, err error) {
	if _, err := uuid.Parse(p.ID); err != nil {
		return nil, &crud.CannotConvertStringToUUID{
			Message: err.Error(),
			ID:      p.ID,
		}
	}
	book, err := s.store.GetBook(ctx, uuid.MustParse(p.ID))
	if err != nil {
		return nil, &crud.IDDoesntExist{
			Message: err.Error(),
			ID:      p.ID,
		}
	}
	id := book.ID.String()
	response := crud.BookResponse{
		Name:  &book.Name,
		Price: &book.Price,
		ID:    &id,
	}
	return &response, nil
}

// Delete Book
func (s *crudsrvc) DeleteBook(ctx context.Context, p string) (err error) {
	if err := s.store.DeleteBook(ctx, uuid.MustParse(p)); err != nil {
		if err != nil {
			s.logger.Print("ERROR_DELETE_ONE_BOOK")
			return err
		}
	}
	s.logger.Print("Success")
	return nil
}

// Create Book
func (s *crudsrvc) CreateBook(ctx context.Context, p *crud.CreateBookPayload) (res *crud.BookResponse, err error) {
	arg := db.CreateBookParams{
		Price: *p.Price,
		Name:  *p.Name,
	}
	book, err := s.store.CreateBook(ctx, arg)
	if err != nil {
		s.logger.Print("ERROR_CREATE_BOOK")
		return nil, err
	}

	newBook, err := s.GetBook(ctx, &crud.GetBookPayload{ID: book.ID.String()})
	if err != nil {
		s.logger.Print("ERROR_GET_ONE_BOOK")
		return nil, err
	}

	return newBook, nil
}

// Read Books
func (s *crudsrvc) GetAllBooks(ctx context.Context) (res []*crud.BookResponse, err error) {
	books, err := s.store.GetBooks(ctx)
	if err != nil {
		s.logger.Print("ERROR_GET_ALL_BOOK")
		return nil, err
	}
	var response []*crud.BookResponse
	for _, v := range books {
		id := v.ID.String()
		response = append(response, &crud.BookResponse{
			ID:    &id,
			Name:  &v.Name,
			Price: &v.Price,
		})
	}
	return response, nil
}
