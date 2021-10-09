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
)

// Endpoints wraps the "crud" service endpoints.
type Endpoints struct {
	GetBook     goa.Endpoint
	UpdateBook  goa.Endpoint
	GetAllBooks goa.Endpoint
	DeleteBook  goa.Endpoint
	CreateBook  goa.Endpoint
	Signup      goa.Endpoint
	Signin      goa.Endpoint
}

// NewEndpoints wraps the methods of the "crud" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		GetBook:     NewGetBookEndpoint(s),
		UpdateBook:  NewUpdateBookEndpoint(s),
		GetAllBooks: NewGetAllBooksEndpoint(s),
		DeleteBook:  NewDeleteBookEndpoint(s),
		CreateBook:  NewCreateBookEndpoint(s),
		Signup:      NewSignupEndpoint(s),
		Signin:      NewSigninEndpoint(s),
	}
}

// Use applies the given middleware to all the "crud" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.GetBook = m(e.GetBook)
	e.UpdateBook = m(e.UpdateBook)
	e.GetAllBooks = m(e.GetAllBooks)
	e.DeleteBook = m(e.DeleteBook)
	e.CreateBook = m(e.CreateBook)
	e.Signup = m(e.Signup)
	e.Signin = m(e.Signin)
}

// NewGetBookEndpoint returns an endpoint function that calls the method
// "getBook" of service "crud".
func NewGetBookEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*GetBookPayload)
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

// NewSignupEndpoint returns an endpoint function that calls the method
// "signup" of service "crud".
func NewSignupEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*SignupPayload)
		return s.Signup(ctx, p)
	}
}

// NewSigninEndpoint returns an endpoint function that calls the method
// "signin" of service "crud".
func NewSigninEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*SigninPayload)
		return s.Signin(ctx, p)
	}
}
