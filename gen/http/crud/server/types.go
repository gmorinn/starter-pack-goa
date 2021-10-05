// Code generated by goa v3.5.2, DO NOT EDIT.
//
// crud HTTP server types
//
// Command:
// $ goa gen api_crud/design

package server

import (
	crud "api_crud/gen/crud"
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// UpdateBookRequestBody is the type of the "crud" service "updateBook"
// endpoint HTTP request body.
type UpdateBookRequestBody struct {
	Name  *string  `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	Price *float64 `form:"price,omitempty" json:"price,omitempty" xml:"price,omitempty"`
}

// CreateBookRequestBody is the type of the "crud" service "createBook"
// endpoint HTTP request body.
type CreateBookRequestBody struct {
	Name  *string  `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	Price *float64 `form:"price,omitempty" json:"price,omitempty" xml:"price,omitempty"`
}

// GetBookResponseBody is the type of the "crud" service "getBook" endpoint
// HTTP response body.
type GetBookResponseBody struct {
	ID      string  `form:"id" json:"id" xml:"id"`
	Name    string  `form:"name" json:"name" xml:"name"`
	Price   float64 `form:"price" json:"price" xml:"price"`
	Success bool    `form:"success" json:"success" xml:"success"`
}

// UpdateBookResponseBody is the type of the "crud" service "updateBook"
// endpoint HTTP response body.
type UpdateBookResponseBody struct {
	ID      string  `form:"id" json:"id" xml:"id"`
	Name    string  `form:"name" json:"name" xml:"name"`
	Price   float64 `form:"price" json:"price" xml:"price"`
	Success bool    `form:"success" json:"success" xml:"success"`
}

// GetAllBooksResponseBody is the type of the "crud" service "getAllBooks"
// endpoint HTTP response body.
type GetAllBooksResponseBody struct {
	Books   []*BookResponseResponseBody `form:"books" json:"books" xml:"books"`
	Success bool                        `form:"success" json:"success" xml:"success"`
}

// DeleteBookResponseBody is the type of the "crud" service "deleteBook"
// endpoint HTTP response body.
type DeleteBookResponseBody struct {
	Success bool `form:"success" json:"success" xml:"success"`
}

// CreateBookResponseBody is the type of the "crud" service "createBook"
// endpoint HTTP response body.
type CreateBookResponseBody struct {
	Book    *BookResponseResponseBody `form:"book" json:"book" xml:"book"`
	Success bool                      `form:"success" json:"success" xml:"success"`
}

// GetBookIDDoesntExistResponseBody is the type of the "crud" service "getBook"
// endpoint HTTP response body for the "id_doesnt_exist" error.
type GetBookIDDoesntExistResponseBody struct {
	// Returning error
	Message string `form:"message" json:"message" xml:"message"`
	// Wrong Id
	ID      string `form:"id" json:"id" xml:"id"`
	Success bool   `form:"success" json:"success" xml:"success"`
}

// GetBookUnknownErrorResponseBody is the type of the "crud" service "getBook"
// endpoint HTTP response body for the "unknown_error" error.
type GetBookUnknownErrorResponseBody struct {
	// Returning error
	Message string `form:"message" json:"message" xml:"message"`
	Success bool   `form:"success" json:"success" xml:"success"`
}

// UpdateBookIDDoesntExistResponseBody is the type of the "crud" service
// "updateBook" endpoint HTTP response body for the "id_doesnt_exist" error.
type UpdateBookIDDoesntExistResponseBody struct {
	// Returning error
	Message string `form:"message" json:"message" xml:"message"`
	// Wrong Id
	ID      string `form:"id" json:"id" xml:"id"`
	Success bool   `form:"success" json:"success" xml:"success"`
}

// UpdateBookUnknownErrorResponseBody is the type of the "crud" service
// "updateBook" endpoint HTTP response body for the "unknown_error" error.
type UpdateBookUnknownErrorResponseBody struct {
	// Returning error
	Message string `form:"message" json:"message" xml:"message"`
	Success bool   `form:"success" json:"success" xml:"success"`
}

// GetAllBooksUnknownErrorResponseBody is the type of the "crud" service
// "getAllBooks" endpoint HTTP response body for the "unknown_error" error.
type GetAllBooksUnknownErrorResponseBody struct {
	// Returning error
	Message string `form:"message" json:"message" xml:"message"`
	Success bool   `form:"success" json:"success" xml:"success"`
}

// DeleteBookIDDoesntExistResponseBody is the type of the "crud" service
// "deleteBook" endpoint HTTP response body for the "id_doesnt_exist" error.
type DeleteBookIDDoesntExistResponseBody struct {
	// Returning error
	Message string `form:"message" json:"message" xml:"message"`
	// Wrong Id
	ID      string `form:"id" json:"id" xml:"id"`
	Success bool   `form:"success" json:"success" xml:"success"`
}

// DeleteBookUnknownErrorResponseBody is the type of the "crud" service
// "deleteBook" endpoint HTTP response body for the "unknown_error" error.
type DeleteBookUnknownErrorResponseBody struct {
	// Returning error
	Message string `form:"message" json:"message" xml:"message"`
	Success bool   `form:"success" json:"success" xml:"success"`
}

// CreateBookUnknownErrorResponseBody is the type of the "crud" service
// "createBook" endpoint HTTP response body for the "unknown_error" error.
type CreateBookUnknownErrorResponseBody struct {
	// Returning error
	Message string `form:"message" json:"message" xml:"message"`
	Success bool   `form:"success" json:"success" xml:"success"`
}

// BookResponseResponseBody is used to define fields on response body types.
type BookResponseResponseBody struct {
	ID    string  `form:"id" json:"id" xml:"id"`
	Name  string  `form:"name" json:"name" xml:"name"`
	Price float64 `form:"price" json:"price" xml:"price"`
}

// NewGetBookResponseBody builds the HTTP response body from the result of the
// "getBook" endpoint of the "crud" service.
func NewGetBookResponseBody(res *crud.GetBookResult) *GetBookResponseBody {
	body := &GetBookResponseBody{
		ID:      res.ID,
		Name:    res.Name,
		Price:   res.Price,
		Success: res.Success,
	}
	return body
}

// NewUpdateBookResponseBody builds the HTTP response body from the result of
// the "updateBook" endpoint of the "crud" service.
func NewUpdateBookResponseBody(res *crud.UpdateBookResult) *UpdateBookResponseBody {
	body := &UpdateBookResponseBody{
		ID:      res.ID,
		Name:    res.Name,
		Price:   res.Price,
		Success: res.Success,
	}
	return body
}

// NewGetAllBooksResponseBody builds the HTTP response body from the result of
// the "getAllBooks" endpoint of the "crud" service.
func NewGetAllBooksResponseBody(res *crud.GetAllBooksResult) *GetAllBooksResponseBody {
	body := &GetAllBooksResponseBody{
		Success: res.Success,
	}
	if res.Books != nil {
		body.Books = make([]*BookResponseResponseBody, len(res.Books))
		for i, val := range res.Books {
			body.Books[i] = marshalCrudBookResponseToBookResponseResponseBody(val)
		}
	}
	return body
}

// NewDeleteBookResponseBody builds the HTTP response body from the result of
// the "deleteBook" endpoint of the "crud" service.
func NewDeleteBookResponseBody(res *crud.DeleteBookResult) *DeleteBookResponseBody {
	body := &DeleteBookResponseBody{
		Success: res.Success,
	}
	return body
}

// NewCreateBookResponseBody builds the HTTP response body from the result of
// the "createBook" endpoint of the "crud" service.
func NewCreateBookResponseBody(res *crud.CreateBookResult) *CreateBookResponseBody {
	body := &CreateBookResponseBody{
		Success: res.Success,
	}
	if res.Book != nil {
		body.Book = marshalCrudBookResponseToBookResponseResponseBody(res.Book)
	}
	return body
}

// NewGetBookIDDoesntExistResponseBody builds the HTTP response body from the
// result of the "getBook" endpoint of the "crud" service.
func NewGetBookIDDoesntExistResponseBody(res *crud.IDDoesntExist) *GetBookIDDoesntExistResponseBody {
	body := &GetBookIDDoesntExistResponseBody{
		Message: res.Message,
		ID:      res.ID,
		Success: res.Success,
	}
	return body
}

// NewGetBookUnknownErrorResponseBody builds the HTTP response body from the
// result of the "getBook" endpoint of the "crud" service.
func NewGetBookUnknownErrorResponseBody(res *crud.UnknownError) *GetBookUnknownErrorResponseBody {
	body := &GetBookUnknownErrorResponseBody{
		Message: res.Message,
		Success: res.Success,
	}
	return body
}

// NewUpdateBookIDDoesntExistResponseBody builds the HTTP response body from
// the result of the "updateBook" endpoint of the "crud" service.
func NewUpdateBookIDDoesntExistResponseBody(res *crud.IDDoesntExist) *UpdateBookIDDoesntExistResponseBody {
	body := &UpdateBookIDDoesntExistResponseBody{
		Message: res.Message,
		ID:      res.ID,
		Success: res.Success,
	}
	return body
}

// NewUpdateBookUnknownErrorResponseBody builds the HTTP response body from the
// result of the "updateBook" endpoint of the "crud" service.
func NewUpdateBookUnknownErrorResponseBody(res *crud.UnknownError) *UpdateBookUnknownErrorResponseBody {
	body := &UpdateBookUnknownErrorResponseBody{
		Message: res.Message,
		Success: res.Success,
	}
	return body
}

// NewGetAllBooksUnknownErrorResponseBody builds the HTTP response body from
// the result of the "getAllBooks" endpoint of the "crud" service.
func NewGetAllBooksUnknownErrorResponseBody(res *crud.UnknownError) *GetAllBooksUnknownErrorResponseBody {
	body := &GetAllBooksUnknownErrorResponseBody{
		Message: res.Message,
		Success: res.Success,
	}
	return body
}

// NewDeleteBookIDDoesntExistResponseBody builds the HTTP response body from
// the result of the "deleteBook" endpoint of the "crud" service.
func NewDeleteBookIDDoesntExistResponseBody(res *crud.IDDoesntExist) *DeleteBookIDDoesntExistResponseBody {
	body := &DeleteBookIDDoesntExistResponseBody{
		Message: res.Message,
		ID:      res.ID,
		Success: res.Success,
	}
	return body
}

// NewDeleteBookUnknownErrorResponseBody builds the HTTP response body from the
// result of the "deleteBook" endpoint of the "crud" service.
func NewDeleteBookUnknownErrorResponseBody(res *crud.UnknownError) *DeleteBookUnknownErrorResponseBody {
	body := &DeleteBookUnknownErrorResponseBody{
		Message: res.Message,
		Success: res.Success,
	}
	return body
}

// NewCreateBookUnknownErrorResponseBody builds the HTTP response body from the
// result of the "createBook" endpoint of the "crud" service.
func NewCreateBookUnknownErrorResponseBody(res *crud.UnknownError) *CreateBookUnknownErrorResponseBody {
	body := &CreateBookUnknownErrorResponseBody{
		Message: res.Message,
		Success: res.Success,
	}
	return body
}

// NewGetBookPayload builds a crud service getBook endpoint payload.
func NewGetBookPayload(id string) *crud.GetBookPayload {
	v := &crud.GetBookPayload{}
	v.ID = id

	return v
}

// NewUpdateBookPayload builds a crud service updateBook endpoint payload.
func NewUpdateBookPayload(body *UpdateBookRequestBody, id string) *crud.UpdateBookPayload {
	v := &crud.UpdateBookPayload{
		Name:  *body.Name,
		Price: *body.Price,
	}
	v.ID = id

	return v
}

// NewDeleteBookPayload builds a crud service deleteBook endpoint payload.
func NewDeleteBookPayload(id string) *crud.DeleteBookPayload {
	v := &crud.DeleteBookPayload{}
	v.ID = id

	return v
}

// NewCreateBookPayload builds a crud service createBook endpoint payload.
func NewCreateBookPayload(body *CreateBookRequestBody) *crud.CreateBookPayload {
	v := &crud.CreateBookPayload{
		Name:  *body.Name,
		Price: *body.Price,
	}

	return v
}

// ValidateUpdateBookRequestBody runs the validations defined on
// UpdateBookRequestBody
func ValidateUpdateBookRequestBody(body *UpdateBookRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Price == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("price", "body"))
	}
	return
}

// ValidateCreateBookRequestBody runs the validations defined on
// CreateBookRequestBody
func ValidateCreateBookRequestBody(body *CreateBookRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Price == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("price", "body"))
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 3, true))
		}
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) > 10 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 10, false))
		}
	}
	return
}
