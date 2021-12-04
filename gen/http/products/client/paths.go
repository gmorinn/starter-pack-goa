// Code generated by goa v3.5.2, DO NOT EDIT.
//
// HTTP request path constructors for the products service.
//
// Command:
// $ goa gen api_crud/design

package client

import (
	"fmt"
)

// GetAllProductsByCategoryProductsPath returns the URL path to the products service getAllProductsByCategory HTTP endpoint.
func GetAllProductsByCategoryProductsPath(category string) string {
	return fmt.Sprintf("/v1/web/products/category/%v", category)
}

// GetAllProductsProductsPath returns the URL path to the products service getAllProducts HTTP endpoint.
func GetAllProductsProductsPath() string {
	return "/v1/web/products"
}

// GetProductProductsPath returns the URL path to the products service getProduct HTTP endpoint.
func GetProductProductsPath(id string) string {
	return fmt.Sprintf("/v1/web/product/%v", id)
}
