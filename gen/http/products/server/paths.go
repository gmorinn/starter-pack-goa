// Code generated by goa v3.5.2, DO NOT EDIT.
//
// HTTP request path constructors for the products service.
//
// Command:
// $ goa gen api_crud/design

package server

import (
	"fmt"
)

// GetAllProductsProductsPath returns the URL path to the products service getAllProducts HTTP endpoint.
func GetAllProductsProductsPath() string {
	return "/web/products"
}

// GetAllProductsByCategoryProductsPath returns the URL path to the products service getAllProductsByCategory HTTP endpoint.
func GetAllProductsByCategoryProductsPath(category string) string {
	return fmt.Sprintf("/web/products/category/%v", category)
}

// DeleteProductProductsPath returns the URL path to the products service deleteProduct HTTP endpoint.
func DeleteProductProductsPath(id string) string {
	return fmt.Sprintf("/web/product/remove/%v", id)
}

// CreateProductProductsPath returns the URL path to the products service createProduct HTTP endpoint.
func CreateProductProductsPath() string {
	return "/web/product/add"
}

// UpdateProductProductsPath returns the URL path to the products service updateProduct HTTP endpoint.
func UpdateProductProductsPath(id string) string {
	return fmt.Sprintf("/web/product/%v", id)
}

// GetProductProductsPath returns the URL path to the products service getProduct HTTP endpoint.
func GetProductProductsPath(id string) string {
	return fmt.Sprintf("/web/product/%v", id)
}
