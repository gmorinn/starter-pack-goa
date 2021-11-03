package api

import (
	book "api_crud/gen/book"
	db "api_crud/internal"
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
)

// book service example implementation.
// The example methods log the requests and return zero values.
type booksrvc struct {
	logger *log.Logger
	server *Server
}

func (s *booksrvc) ErrorResponse(msg string, err error) *book.UnknownError {
	return &book.UnknownError{
		Err:       err.Error(),
		ErrorCode: msg,
	}
}

// NewBook returns the book service implementation.
func NewBook(logger *log.Logger, server *Server) book.Service {
	return &booksrvc{logger, server}
}

// Get one item
func (s *booksrvc) GetBook(ctx context.Context, p *book.GetBookPayload) (res *book.GetBookResult, err error) {
	err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		b, err := q.GetBook(ctx, uuid.MustParse(p.ID))
		if err != nil {
			return fmt.Errorf("ERROR_GET_BOOK_BY_ID %v", err)
		}
		res = &book.GetBookResult{
			ID:      b.ID.String(),
			Name:    b.Name,
			Price:   b.Price,
			Success: true,
		}
		return nil
	})

	if err != nil {
		return nil, s.ErrorResponse("TX_GET_BOOK", err)
	}

	return res, nil
}

// Update one item
func (s *booksrvc) UpdateBook(ctx context.Context, p *book.UpdateBookPayload) (res *book.UpdateBookResult, err error) {

	err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		arg := db.UpdateBookParams{
			ID:    uuid.MustParse(p.ID),
			Name:  p.Name,
			Price: p.Price,
		}
		if err := q.UpdateBook(ctx, arg); err != nil {
			return fmt.Errorf("ERROR_UPDATE_BOOK %v", err)
		}
		newBook, err := q.GetBook(ctx, uuid.MustParse(p.ID))
		if err != nil {
			return fmt.Errorf("ERROR_GET_BOOK_BY_ID %v", err)
		}
		res = &book.UpdateBookResult{
			Book: &book.BookResponse{
				ID:    newBook.ID.String(),
				Name:  newBook.Name,
				Price: newBook.Price,
			},
			Success: true,
		}
		return nil
	})

	if err != nil {
		return nil, s.ErrorResponse("TX_UPDATE_BOOK", err)
	}

	return res, nil
}

// Read All items
func (s *booksrvc) GetAllBooks(ctx context.Context) (res *book.GetAllBooksResult, err error) {
	err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		books, err := s.server.Store.GetBooks(ctx)
		if err != nil {
			return fmt.Errorf("ERROR_GET_ALL_BOOKS %v", err)
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
		res = &book.GetAllBooksResult{
			Books:   BookResponse,
			Success: true,
		}
		return nil
	})

	if err != nil {
		return nil, s.ErrorResponse("TX_GET_ALL_BOOKS", err)
	}

	return res, nil
}

// Delete one item by ID
func (s *booksrvc) DeleteBook(ctx context.Context, p *book.DeleteBookPayload) (res *book.DeleteBookResult, err error) {
	err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		if err := q.DeleteBook(ctx, uuid.MustParse(p.ID)); err != nil {
			return fmt.Errorf("ERROR_DELETE_BOOK %v", err)
		}
		return nil
	})
	if err != nil {
		return nil, s.ErrorResponse("TX_DELETE_BOOK", err)
	}
	return &book.DeleteBookResult{Success: true}, nil
}

// Create one item
func (s *booksrvc) CreateBook(ctx context.Context, p *book.CreateBookPayload) (res *book.CreateBookResult, err error) {
	err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		arg := db.CreateBookParams{
			Price: p.Price,
			Name:  p.Name,
		}
		b, err := q.CreateBook(ctx, arg)
		if err != nil {
			return fmt.Errorf("ERROR_CREATE_BOOK %v", err)
		}

		newBook, err := q.GetBook(ctx, b.ID)
		if err != nil {
			return fmt.Errorf("ERROR_GET_BOOK %v", err)
		}
		res = &book.CreateBookResult{
			Book: &book.BookResponse{
				ID:    newBook.ID.String(),
				Name:  newBook.Name,
				Price: newBook.Price,
			},
			Success: true,
		}
		return nil
	})

	if err != nil {
		return nil, s.ErrorResponse("TX_CREATE_BOOK", err)
	}

	return res, nil
}
