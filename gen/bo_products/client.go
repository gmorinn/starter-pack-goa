// Code generated by goa v3.5.2, DO NOT EDIT.
//
// boProducts client
//
// Command:
// $ goa gen api_crud/design

package boproducts

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "boProducts" service client.
type Client struct {
	GetAllProductsEndpoint           goa.Endpoint
	GetAllProductsByCategoryEndpoint goa.Endpoint
	DeleteProductEndpoint            goa.Endpoint
	CreateProductEndpoint            goa.Endpoint
	UpdateProductEndpoint            goa.Endpoint
	DeleteManyProductsEndpoint       goa.Endpoint
	GetProductEndpoint               goa.Endpoint
}

// NewClient initializes a "boProducts" service client given the endpoints.
func NewClient(getAllProducts, getAllProductsByCategory, deleteProduct, createProduct, updateProduct, deleteManyProducts, getProduct goa.Endpoint) *Client {
	return &Client{
		GetAllProductsEndpoint:           getAllProducts,
		GetAllProductsByCategoryEndpoint: getAllProductsByCategory,
		DeleteProductEndpoint:            deleteProduct,
		CreateProductEndpoint:            createProduct,
		UpdateProductEndpoint:            updateProduct,
		DeleteManyProductsEndpoint:       deleteManyProducts,
		GetProductEndpoint:               getProduct,
	}
}

// GetAllProducts calls the "getAllProducts" endpoint of the "boProducts"
// service.
func (c *Client) GetAllProducts(ctx context.Context, p *GetAllProductsPayload) (res *GetAllProductsResult, err error) {
	var ires interface{}
	ires, err = c.GetAllProductsEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*GetAllProductsResult), nil
}

// GetAllProductsByCategory calls the "getAllProductsByCategory" endpoint of
// the "boProducts" service.
func (c *Client) GetAllProductsByCategory(ctx context.Context, p *GetAllProductsByCategoryPayload) (res *GetAllProductsByCategoryResult, err error) {
	var ires interface{}
	ires, err = c.GetAllProductsByCategoryEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*GetAllProductsByCategoryResult), nil
}

// DeleteProduct calls the "deleteProduct" endpoint of the "boProducts" service.
func (c *Client) DeleteProduct(ctx context.Context, p *DeleteProductPayload) (res *DeleteProductResult, err error) {
	var ires interface{}
	ires, err = c.DeleteProductEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*DeleteProductResult), nil
}

// CreateProduct calls the "createProduct" endpoint of the "boProducts" service.
func (c *Client) CreateProduct(ctx context.Context, p *CreateProductPayload) (res *CreateProductResult, err error) {
	var ires interface{}
	ires, err = c.CreateProductEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*CreateProductResult), nil
}

// UpdateProduct calls the "updateProduct" endpoint of the "boProducts" service.
func (c *Client) UpdateProduct(ctx context.Context, p *UpdateProductPayload) (res *UpdateProductResult, err error) {
	var ires interface{}
	ires, err = c.UpdateProductEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*UpdateProductResult), nil
}

// DeleteManyProducts calls the "deleteManyProducts" endpoint of the
// "boProducts" service.
func (c *Client) DeleteManyProducts(ctx context.Context, p *DeleteManyProductsPayload) (res *DeleteManyProductsResult, err error) {
	var ires interface{}
	ires, err = c.DeleteManyProductsEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*DeleteManyProductsResult), nil
}

// GetProduct calls the "getProduct" endpoint of the "boProducts" service.
func (c *Client) GetProduct(ctx context.Context, p *GetProductPayload) (res *GetProductResult, err error) {
	var ires interface{}
	ires, err = c.GetProductEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*GetProductResult), nil
}
