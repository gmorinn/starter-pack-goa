// Code generated by goa v3.5.2, DO NOT EDIT.
//
// products service
//
// Command:
// $ goa gen api_crud/design

package products

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Products of the E-Commerce
type Service interface {
	// Get All products by category
	GetAllProductsByCategory(context.Context, *GetAllProductsByCategoryPayload) (res *GetAllProductsByCategoryResult, err error)
	// Get All products
	GetAllProducts(context.Context, *GetAllProductsPayload) (res *GetAllProductsResult, err error)
	// Get one product
	GetProduct(context.Context, *GetProductPayload) (res *GetProductResult, err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// OAuth2Auth implements the authorization logic for the OAuth2 security scheme.
	OAuth2Auth(ctx context.Context, token string, schema *security.OAuth2Scheme) (context.Context, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "products"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [3]string{"getAllProductsByCategory", "getAllProducts", "getProduct"}

// GetAllProductsByCategoryPayload is the payload type of the products service
// getAllProductsByCategory method.
type GetAllProductsByCategoryPayload struct {
	Category string
	// Use to generate Oauth with /authorization
	Oauth *string
}

// GetAllProductsByCategoryResult is the result type of the products service
// getAllProductsByCategory method.
type GetAllProductsByCategoryResult struct {
	// Result is an array of object
	Products []*ResProduct
	Success  bool
}

// GetAllProductsPayload is the payload type of the products service
// getAllProducts method.
type GetAllProductsPayload struct {
	// Use to generate Oauth with /authorization
	Oauth *string
}

// GetAllProductsResult is the result type of the products service
// getAllProducts method.
type GetAllProductsResult struct {
	// Result is an array of object
	Products []*ResProduct
	Success  bool
}

// GetProductPayload is the payload type of the products service getProduct
// method.
type GetProductPayload struct {
	// Unique ID of the product
	ID string
	// Use to generate Oauth with /authorization
	Oauth *string
}

// GetProductResult is the result type of the products service getProduct
// method.
type GetProductResult struct {
	// Result is an object
	Product *ResProduct
	Success bool
}

type ResProduct struct {
	ID       string
	Name     string
	Price    float64
	Cover    string
	Category string
}

type UnknownError struct {
	Err       string
	ErrorCode string
	Success   bool
}

// Error returns an error description.
func (e *UnknownError) Error() string {
	return ""
}

// ErrorName returns "unknownError".
func (e *UnknownError) ErrorName() string {
	return "unknown_error"
}

// MakeTimeout builds a goa.ServiceError from an error.
func MakeTimeout(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "timeout",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
		Timeout: true,
	}
}
