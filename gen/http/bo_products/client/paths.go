// Code generated by goa v3.5.2, DO NOT EDIT.
//
// HTTP request path constructors for the boProducts service.
//
// Command:
// $ goa gen api_crud/design

package client

import (
	"fmt"
)

// GetAllProductsBoProductsPath returns the URL path to the boProducts service getAllProducts HTTP endpoint.
func GetAllProductsBoProductsPath(offset int32, limit int32) string {
	return fmt.Sprintf("/v1/bo/products/%v/%v", offset, limit)
}

// GetAllProductsByCategoryBoProductsPath returns the URL path to the boProducts service getAllProductsByCategory HTTP endpoint.
func GetAllProductsByCategoryBoProductsPath(category string) string {
	return fmt.Sprintf("/v1/bo/products/category/%v", category)
}

// DeleteProductBoProductsPath returns the URL path to the boProducts service deleteProduct HTTP endpoint.
func DeleteProductBoProductsPath(id string) string {
	return fmt.Sprintf("/v1/bo/product/remove/%v", id)
}

// CreateProductBoProductsPath returns the URL path to the boProducts service createProduct HTTP endpoint.
func CreateProductBoProductsPath() string {
	return "/v1/bo/product/add"
}

// UpdateProductBoProductsPath returns the URL path to the boProducts service updateProduct HTTP endpoint.
func UpdateProductBoProductsPath(id string) string {
	return fmt.Sprintf("/v1/bo/product/%v", id)
}

// DeleteManyProductsBoProductsPath returns the URL path to the boProducts service deleteManyProducts HTTP endpoint.
func DeleteManyProductsBoProductsPath() string {
	return "/v1/bo/products/remove"
}

// GetProductBoProductsPath returns the URL path to the boProducts service getProduct HTTP endpoint.
func GetProductBoProductsPath(id string) string {
	return fmt.Sprintf("/v1/bo/product/%v", id)
}
