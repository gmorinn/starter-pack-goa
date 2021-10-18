// Code generated by goa v3.5.2, DO NOT EDIT.
//
// crud endpoints
//
// Command:
// $ goa gen api_crud/design

package crud

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "crud" service endpoints.
type Endpoints struct {
	GetBook     goa.Endpoint
	UpdateBook  goa.Endpoint
	GetAllBooks goa.Endpoint
	DeleteBook  goa.Endpoint
	CreateBook  goa.Endpoint
}

// NewEndpoints wraps the methods of the "crud" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		GetBook:     NewGetBookEndpoint(s, a.JWTAuth),
		UpdateBook:  NewUpdateBookEndpoint(s),
		GetAllBooks: NewGetAllBooksEndpoint(s),
		DeleteBook:  NewDeleteBookEndpoint(s),
		CreateBook:  NewCreateBookEndpoint(s),
	}
}

// Use applies the given middleware to all the "crud" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.GetBook = m(e.GetBook)
	e.UpdateBook = m(e.UpdateBook)
	e.GetAllBooks = m(e.GetAllBooks)
	e.DeleteBook = m(e.DeleteBook)
	e.CreateBook = m(e.CreateBook)
}

// NewGetBookEndpoint returns an endpoint function that calls the method
// "getBook" of service "crud".
func NewGetBookEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*GetBookPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"api:read", "api:write"},
			RequiredScopes: []string{},
		}
		var token string
		if p.JWTToken != nil {
			token = *p.JWTToken
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return s.GetBook(ctx, p)
	}
}

// NewUpdateBookEndpoint returns an endpoint function that calls the method
// "updateBook" of service "crud".
func NewUpdateBookEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*UpdateBookPayload)
		return s.UpdateBook(ctx, p)
	}
}

// NewGetAllBooksEndpoint returns an endpoint function that calls the method
// "getAllBooks" of service "crud".
func NewGetAllBooksEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetAllBooks(ctx)
	}
}

// NewDeleteBookEndpoint returns an endpoint function that calls the method
// "deleteBook" of service "crud".
func NewDeleteBookEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*DeleteBookPayload)
		return s.DeleteBook(ctx, p)
	}
}

// NewCreateBookEndpoint returns an endpoint function that calls the method
// "createBook" of service "crud".
func NewCreateBookEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*CreateBookPayload)
		return s.CreateBook(ctx, p)
	}
}
