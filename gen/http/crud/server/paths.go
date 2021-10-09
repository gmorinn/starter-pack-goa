// Code generated by goa v3.5.2, DO NOT EDIT.
//
// HTTP request path constructors for the crud service.
//
// Command:
// $ goa gen api_crud/design

package server

import (
	"fmt"
)

// GetBookCrudPath returns the URL path to the crud service getBook HTTP endpoint.
func GetBookCrudPath(id string) string {
	return fmt.Sprintf("/book/%v", id)
}

// UpdateBookCrudPath returns the URL path to the crud service updateBook HTTP endpoint.
func UpdateBookCrudPath(id string) string {
	return fmt.Sprintf("/book/%v", id)
}

// GetAllBooksCrudPath returns the URL path to the crud service getAllBooks HTTP endpoint.
func GetAllBooksCrudPath() string {
	return "/books"
}

// DeleteBookCrudPath returns the URL path to the crud service deleteBook HTTP endpoint.
func DeleteBookCrudPath(id string) string {
	return fmt.Sprintf("/book/remove/%v", id)
}

// CreateBookCrudPath returns the URL path to the crud service createBook HTTP endpoint.
func CreateBookCrudPath() string {
	return "/book/add"
}

// SignupCrudPath returns the URL path to the crud service signup HTTP endpoint.
func SignupCrudPath() string {
	return "/signup"
}

// SigninCrudPath returns the URL path to the crud service signin HTTP endpoint.
func SigninCrudPath() string {
	return "/signin"
}
