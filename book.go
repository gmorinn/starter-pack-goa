package basic

import (
	"api_crud/api"
	book "api_crud/gen/book"
	"api_crud/internal/db"
	"context"
	"log"

	"github.com/google/uuid"
)

// book service example implementation.
// The example methods log the requests and return zero values.
type booksrvc struct {
	logger *log.Logger
	server *api.Server
}

func Error_ID(msg, id string, err error) *book.IDDoesntExist {
	return &book.IDDoesntExist{
		Err: err.Error(),
		ID:  id,
	}
}

func ErrorResponse(msg string, err error) *book.UnknownError {
	return &book.UnknownError{
		Err:       err.Error(),
		ErrorCode: msg,
	}
}

// NewBook returns the book service implementation.
func NewBook(logger *log.Logger, server *api.Server) book.Service {
	return &booksrvc{logger, server}
}

// Get one item
func (s *booksrvc) GetBook(ctx context.Context, p *book.GetBookPayload) (res *book.GetBookResult, err error) {
	b, err := s.server.Store.GetBook(ctx, uuid.MustParse(p.ID))
	if err != nil {
		return nil, Error_ID("ERROR_GET_BOOK", p.ID, err)
	}

	response := book.GetBookResult{
		ID:      b.ID.String(),
		Name:    b.Name,
		Price:   b.Price,
		Success: true,
	}
	return &response, nil
}

// Update one item
func (s *booksrvc) UpdateBook(ctx context.Context, p *book.UpdateBookPayload) (res *book.UpdateBookResult, err error) {
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
	response := book.UpdateBookResult{
		ID:      newBook.ID.String(),
		Name:    newBook.Name,
		Price:   newBook.Price,
		Success: true,
	}
	return &response, nil
}

// Read All items
func (s *booksrvc) GetAllBooks(ctx context.Context) (res *book.GetAllBooksResult, err error) {
	books, err := s.server.Store.GetBooks(ctx)
	if err != nil {
		return nil, ErrorResponse("ERROR_GET_ALL_BOOKS", err)

	}

	var BookResponse []*book.BookResponse
	for _, v := range books {
		id := v.ID.String()
		BookResponse = append(BookResponse, &book.BookResponse{
			ID:    id,
			Name:  v.Name,
			Price: v.Price,
		})
	}

	response := book.GetAllBooksResult{
		Books:   BookResponse,
		Success: true,
	}
	return &response, nil
}

// Delete one item by ID
func (s *booksrvc) DeleteBook(ctx context.Context, p *book.DeleteBookPayload) (res *book.DeleteBookResult, err error) {
	if err := s.server.Store.DeleteBook(ctx, uuid.MustParse(p.ID)); err != nil {
		return nil, ErrorResponse("ERROR_DELETE_BOOK", err)
	}
	return &book.DeleteBookResult{Success: true}, nil
}

// Create one item
func (s *booksrvc) CreateBook(ctx context.Context, p *book.CreateBookPayload) (res *book.CreateBookResult, err error) {
	arg := db.CreateBookParams{
		Price: p.Price,
		Name:  p.Name,
	}
	b, err := s.server.Store.CreateBook(ctx, arg)
	if err != nil {
		return nil, ErrorResponse("ERROR_CREATE_BOOK", err)
	}

	newBook, err := s.GetBook(ctx, &book.GetBookPayload{ID: b.ID.String()})
	if err != nil {
		return nil, Error_ID("ERROR_GET_BOOK", b.ID.String(), err)
	}
	response := book.CreateBookResult{
		Book: &book.BookResponse{
			ID:    newBook.ID,
			Name:  newBook.Name,
			Price: newBook.Price,
		},
		Success: true,
	}

	return &response, nil
}
