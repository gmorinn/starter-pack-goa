// Code generated by goa v3.5.2, DO NOT EDIT.
//
// boProducts HTTP client CLI support package
//
// Command:
// $ goa gen api_crud/design

package client

import (
	boproducts "api_crud/gen/bo_products"
	"encoding/json"
	"fmt"
	"strconv"

	goa "goa.design/goa/v3/pkg"
)

// BuildGetAllProductsPayload builds the payload for the boProducts
// getAllProducts endpoint from CLI flags.
func BuildGetAllProductsPayload(boProductsGetAllProductsOffset string, boProductsGetAllProductsLimit string, boProductsGetAllProductsField string, boProductsGetAllProductsDirection string, boProductsGetAllProductsOauth string, boProductsGetAllProductsJWTToken string) (*boproducts.GetAllProductsPayload, error) {
	var err error
	var offset int32
	{
		var v int64
		v, err = strconv.ParseInt(boProductsGetAllProductsOffset, 10, 32)
		offset = int32(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for offset, must be INT32")
		}
	}
	var limit int32
	{
		var v int64
		v, err = strconv.ParseInt(boProductsGetAllProductsLimit, 10, 32)
		limit = int32(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for limit, must be INT32")
		}
	}
	var field string
	{
		if boProductsGetAllProductsField != "" {
			field = boProductsGetAllProductsField
		}
	}
	var direction string
	{
		if boProductsGetAllProductsDirection != "" {
			direction = boProductsGetAllProductsDirection
			if !(direction == "asc" || direction == "desc") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("direction", direction, []interface{}{"asc", "desc"}))
			}
			if err != nil {
				return nil, err
			}
		}
	}
	var oauth *string
	{
		if boProductsGetAllProductsOauth != "" {
			oauth = &boProductsGetAllProductsOauth
		}
	}
	var jwtToken *string
	{
		if boProductsGetAllProductsJWTToken != "" {
			jwtToken = &boProductsGetAllProductsJWTToken
		}
	}
	v := &boproducts.GetAllProductsPayload{}
	v.Offset = offset
	v.Limit = limit
	v.Field = field
	v.Direction = direction
	v.Oauth = oauth
	v.JWTToken = jwtToken

	return v, nil
}

// BuildGetAllProductsByCategoryPayload builds the payload for the boProducts
// getAllProductsByCategory endpoint from CLI flags.
func BuildGetAllProductsByCategoryPayload(boProductsGetAllProductsByCategoryCategory string, boProductsGetAllProductsByCategoryOauth string, boProductsGetAllProductsByCategoryJWTToken string) (*boproducts.GetAllProductsByCategoryPayload, error) {
	var err error
	var category string
	{
		category = boProductsGetAllProductsByCategoryCategory
		if !(category == "men" || category == "women" || category == "hat" || category == "jacket" || category == "sneaker" || category == "nothing") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("category", category, []interface{}{"men", "women", "hat", "jacket", "sneaker", "nothing"}))
		}
		if err != nil {
			return nil, err
		}
	}
	var oauth *string
	{
		if boProductsGetAllProductsByCategoryOauth != "" {
			oauth = &boProductsGetAllProductsByCategoryOauth
		}
	}
	var jwtToken *string
	{
		if boProductsGetAllProductsByCategoryJWTToken != "" {
			jwtToken = &boProductsGetAllProductsByCategoryJWTToken
		}
	}
	v := &boproducts.GetAllProductsByCategoryPayload{}
	v.Category = category
	v.Oauth = oauth
	v.JWTToken = jwtToken

	return v, nil
}

// BuildDeleteProductPayload builds the payload for the boProducts
// deleteProduct endpoint from CLI flags.
func BuildDeleteProductPayload(boProductsDeleteProductID string, boProductsDeleteProductOauth string, boProductsDeleteProductJWTToken string) (*boproducts.DeleteProductPayload, error) {
	var err error
	var id string
	{
		id = boProductsDeleteProductID
		err = goa.MergeErrors(err, goa.ValidateFormat("id", id, goa.FormatUUID))

		if err != nil {
			return nil, err
		}
	}
	var oauth *string
	{
		if boProductsDeleteProductOauth != "" {
			oauth = &boProductsDeleteProductOauth
		}
	}
	var jwtToken *string
	{
		if boProductsDeleteProductJWTToken != "" {
			jwtToken = &boProductsDeleteProductJWTToken
		}
	}
	v := &boproducts.DeleteProductPayload{}
	v.ID = id
	v.Oauth = oauth
	v.JWTToken = jwtToken

	return v, nil
}

// BuildCreateProductPayload builds the payload for the boProducts
// createProduct endpoint from CLI flags.
func BuildCreateProductPayload(boProductsCreateProductBody string, boProductsCreateProductOauth string, boProductsCreateProductJWTToken string) (*boproducts.CreateProductPayload, error) {
	var err error
	var body CreateProductRequestBody
	{
		err = json.Unmarshal([]byte(boProductsCreateProductBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"product\": {\n         \"category\": \"men\",\n         \"cover\": \"https://i.ibb.co/ypkgK0X/blue-beanie.png\",\n         \"name\": \"Guillaume\",\n         \"price\": 69\n      }\n   }'")
		}
		if body.Product == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("product", "body"))
		}
		if body.Product != nil {
			if err2 := ValidatePayloadProductRequestBody(body.Product); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
		if err != nil {
			return nil, err
		}
	}
	var oauth *string
	{
		if boProductsCreateProductOauth != "" {
			oauth = &boProductsCreateProductOauth
		}
	}
	var jwtToken *string
	{
		if boProductsCreateProductJWTToken != "" {
			jwtToken = &boProductsCreateProductJWTToken
		}
	}
	v := &boproducts.CreateProductPayload{}
	if body.Product != nil {
		v.Product = marshalPayloadProductRequestBodyToBoproductsPayloadProduct(body.Product)
	}
	v.Oauth = oauth
	v.JWTToken = jwtToken

	return v, nil
}

// BuildUpdateProductPayload builds the payload for the boProducts
// updateProduct endpoint from CLI flags.
func BuildUpdateProductPayload(boProductsUpdateProductBody string, boProductsUpdateProductID string, boProductsUpdateProductOauth string, boProductsUpdateProductJWTToken string) (*boproducts.UpdateProductPayload, error) {
	var err error
	var body UpdateProductRequestBody
	{
		err = json.Unmarshal([]byte(boProductsUpdateProductBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"product\": {\n         \"category\": \"men\",\n         \"cover\": \"https://i.ibb.co/ypkgK0X/blue-beanie.png\",\n         \"name\": \"Guillaume\",\n         \"price\": 69\n      }\n   }'")
		}
		if body.Product == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("product", "body"))
		}
		if body.Product != nil {
			if err2 := ValidatePayloadProductRequestBody(body.Product); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
		if err != nil {
			return nil, err
		}
	}
	var id string
	{
		id = boProductsUpdateProductID
		err = goa.MergeErrors(err, goa.ValidateFormat("id", id, goa.FormatUUID))

		if err != nil {
			return nil, err
		}
	}
	var oauth *string
	{
		if boProductsUpdateProductOauth != "" {
			oauth = &boProductsUpdateProductOauth
		}
	}
	var jwtToken *string
	{
		if boProductsUpdateProductJWTToken != "" {
			jwtToken = &boProductsUpdateProductJWTToken
		}
	}
	v := &boproducts.UpdateProductPayload{}
	if body.Product != nil {
		v.Product = marshalPayloadProductRequestBodyToBoproductsPayloadProduct(body.Product)
	}
	v.ID = id
	v.Oauth = oauth
	v.JWTToken = jwtToken

	return v, nil
}

// BuildDeleteManyProductsPayload builds the payload for the boProducts
// deleteManyProducts endpoint from CLI flags.
func BuildDeleteManyProductsPayload(boProductsDeleteManyProductsBody string, boProductsDeleteManyProductsOauth string, boProductsDeleteManyProductsJWTToken string) (*boproducts.DeleteManyProductsPayload, error) {
	var err error
	var body DeleteManyProductsRequestBody
	{
		err = json.Unmarshal([]byte(boProductsDeleteManyProductsBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"tab\": [\n         \"Sed sunt dolores.\",\n         \"Cupiditate suscipit reprehenderit amet cum.\",\n         \"Ut aut provident laboriosam quisquam.\",\n         \"Atque aut vel occaecati perspiciatis doloremque totam.\"\n      ]\n   }'")
		}
		if body.Tab == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("tab", "body"))
		}
		if err != nil {
			return nil, err
		}
	}
	var oauth *string
	{
		if boProductsDeleteManyProductsOauth != "" {
			oauth = &boProductsDeleteManyProductsOauth
		}
	}
	var jwtToken *string
	{
		if boProductsDeleteManyProductsJWTToken != "" {
			jwtToken = &boProductsDeleteManyProductsJWTToken
		}
	}
	v := &boproducts.DeleteManyProductsPayload{}
	if body.Tab != nil {
		v.Tab = make([]string, len(body.Tab))
		for i, val := range body.Tab {
			v.Tab[i] = val
		}
	}
	v.Oauth = oauth
	v.JWTToken = jwtToken

	return v, nil
}

// BuildGetProductPayload builds the payload for the boProducts getProduct
// endpoint from CLI flags.
func BuildGetProductPayload(boProductsGetProductID string, boProductsGetProductOauth string, boProductsGetProductJWTToken string) (*boproducts.GetProductPayload, error) {
	var err error
	var id string
	{
		id = boProductsGetProductID
		err = goa.MergeErrors(err, goa.ValidateFormat("id", id, goa.FormatUUID))

		if err != nil {
			return nil, err
		}
	}
	var oauth *string
	{
		if boProductsGetProductOauth != "" {
			oauth = &boProductsGetProductOauth
		}
	}
	var jwtToken *string
	{
		if boProductsGetProductJWTToken != "" {
			jwtToken = &boProductsGetProductJWTToken
		}
	}
	v := &boproducts.GetProductPayload{}
	v.ID = id
	v.Oauth = oauth
	v.JWTToken = jwtToken

	return v, nil
}
