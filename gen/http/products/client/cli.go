// Code generated by goa v3.5.2, DO NOT EDIT.
//
// products HTTP client CLI support package
//
// Command:
// $ goa gen api_crud/design

package client

import (
	products "api_crud/gen/products"

	goa "goa.design/goa/v3/pkg"
)

// BuildGetAllProductsByCategoryPayload builds the payload for the products
// getAllProductsByCategory endpoint from CLI flags.
func BuildGetAllProductsByCategoryPayload(productsGetAllProductsByCategoryCategory string, productsGetAllProductsByCategoryOauth string) (*products.GetAllProductsByCategoryPayload, error) {
	var err error
	var category string
	{
		category = productsGetAllProductsByCategoryCategory
		if !(category == "men" || category == "women" || category == "hat" || category == "jacket" || category == "sneaker" || category == "nothing") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("category", category, []interface{}{"men", "women", "hat", "jacket", "sneaker", "nothing"}))
		}
		if err != nil {
			return nil, err
		}
	}
	var oauth *string
	{
		if productsGetAllProductsByCategoryOauth != "" {
			oauth = &productsGetAllProductsByCategoryOauth
		}
	}
	v := &products.GetAllProductsByCategoryPayload{}
	v.Category = category
	v.Oauth = oauth

	return v, nil
}

// BuildGetAllProductsPayload builds the payload for the products
// getAllProducts endpoint from CLI flags.
func BuildGetAllProductsPayload(productsGetAllProductsOauth string) (*products.GetAllProductsPayload, error) {
	var oauth *string
	{
		if productsGetAllProductsOauth != "" {
			oauth = &productsGetAllProductsOauth
		}
	}
	v := &products.GetAllProductsPayload{}
	v.Oauth = oauth

	return v, nil
}

// BuildGetProductPayload builds the payload for the products getProduct
// endpoint from CLI flags.
func BuildGetProductPayload(productsGetProductID string, productsGetProductOauth string) (*products.GetProductPayload, error) {
	var err error
	var id string
	{
		id = productsGetProductID
		err = goa.MergeErrors(err, goa.ValidateFormat("id", id, goa.FormatUUID))

		if err != nil {
			return nil, err
		}
	}
	var oauth *string
	{
		if productsGetProductOauth != "" {
			oauth = &productsGetProductOauth
		}
	}
	v := &products.GetProductPayload{}
	v.ID = id
	v.Oauth = oauth

	return v, nil
}
