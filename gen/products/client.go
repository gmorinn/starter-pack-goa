// Code generated by goa v3.5.2, DO NOT EDIT.
//
// products client
//
// Command:
// $ goa gen api_crud/design

package products

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "products" service client.
type Client struct {
	GetAllProductsByCategoryEndpoint goa.Endpoint
	GetAllProductsEndpoint           goa.Endpoint
	GetProductEndpoint               goa.Endpoint
}

// NewClient initializes a "products" service client given the endpoints.
func NewClient(getAllProductsByCategory, getAllProducts, getProduct goa.Endpoint) *Client {
	return &Client{
		GetAllProductsByCategoryEndpoint: getAllProductsByCategory,
		GetAllProductsEndpoint:           getAllProducts,
		GetProductEndpoint:               getProduct,
	}
}

// GetAllProductsByCategory calls the "getAllProductsByCategory" endpoint of
// the "products" service.
func (c *Client) GetAllProductsByCategory(ctx context.Context, p *GetAllProductsByCategoryPayload) (res *GetAllProductsByCategoryResult, err error) {
	var ires interface{}
	ires, err = c.GetAllProductsByCategoryEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*GetAllProductsByCategoryResult), nil
}

// GetAllProducts calls the "getAllProducts" endpoint of the "products" service.
func (c *Client) GetAllProducts(ctx context.Context, p *GetAllProductsPayload) (res *GetAllProductsResult, err error) {
	var ires interface{}
	ires, err = c.GetAllProductsEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*GetAllProductsResult), nil
}

// GetProduct calls the "getProduct" endpoint of the "products" service.
func (c *Client) GetProduct(ctx context.Context, p *GetProductPayload) (res *GetProductResult, err error) {
	var ires interface{}
	ires, err = c.GetProductEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*GetProductResult), nil
}
