// Code generated by goa v3.5.2, DO NOT EDIT.
//
// products HTTP client encoders and decoders
//
// Command:
// $ goa gen api_crud/design

package client

import (
	products "api_crud/gen/products"
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	goahttp "goa.design/goa/v3/http"
)

// BuildGetAllProductsByCategoryRequest instantiates a HTTP request object with
// method and path set to call the "products" service
// "getAllProductsByCategory" endpoint
func (c *Client) BuildGetAllProductsByCategoryRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		category string
	)
	{
		p, ok := v.(*products.GetAllProductsByCategoryPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("products", "getAllProductsByCategory", "*products.GetAllProductsByCategoryPayload", v)
		}
		category = p.Category
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetAllProductsByCategoryProductsPath(category)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("products", "getAllProductsByCategory", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeGetAllProductsByCategoryRequest returns an encoder for requests sent
// to the products getAllProductsByCategory server.
func EncodeGetAllProductsByCategoryRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*products.GetAllProductsByCategoryPayload)
		if !ok {
			return goahttp.ErrInvalidType("products", "getAllProductsByCategory", "*products.GetAllProductsByCategoryPayload", v)
		}
		if p.Oauth != nil {
			head := *p.Oauth
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		if p.JWTToken != nil {
			head := *p.JWTToken
			req.Header.Set("jwtToken", head)
		}
		return nil
	}
}

// DecodeGetAllProductsByCategoryResponse returns a decoder for responses
// returned by the products getAllProductsByCategory endpoint. restoreBody
// controls whether the response body should be restored after having been read.
// DecodeGetAllProductsByCategoryResponse may return the following errors:
//	- "unknown_error" (type *products.UnknownError): http.StatusInternalServerError
//	- error: internal error
func DecodeGetAllProductsByCategoryResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetAllProductsByCategoryResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("products", "getAllProductsByCategory", err)
			}
			err = ValidateGetAllProductsByCategoryResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("products", "getAllProductsByCategory", err)
			}
			res := NewGetAllProductsByCategoryResultOK(&body)
			return res, nil
		case http.StatusInternalServerError:
			var (
				body GetAllProductsByCategoryUnknownErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("products", "getAllProductsByCategory", err)
			}
			err = ValidateGetAllProductsByCategoryUnknownErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("products", "getAllProductsByCategory", err)
			}
			return nil, NewGetAllProductsByCategoryUnknownError(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("products", "getAllProductsByCategory", resp.StatusCode, string(body))
		}
	}
}

// BuildDeleteProductRequest instantiates a HTTP request object with method and
// path set to call the "products" service "deleteProduct" endpoint
func (c *Client) BuildDeleteProductRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*products.DeleteProductPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("products", "deleteProduct", "*products.DeleteProductPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DeleteProductProductsPath(id)}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("products", "deleteProduct", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeDeleteProductRequest returns an encoder for requests sent to the
// products deleteProduct server.
func EncodeDeleteProductRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*products.DeleteProductPayload)
		if !ok {
			return goahttp.ErrInvalidType("products", "deleteProduct", "*products.DeleteProductPayload", v)
		}
		if p.Oauth != nil {
			head := *p.Oauth
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		if p.JWTToken != nil {
			head := *p.JWTToken
			req.Header.Set("jwtToken", head)
		}
		return nil
	}
}

// DecodeDeleteProductResponse returns a decoder for responses returned by the
// products deleteProduct endpoint. restoreBody controls whether the response
// body should be restored after having been read.
// DecodeDeleteProductResponse may return the following errors:
//	- "unknown_error" (type *products.UnknownError): http.StatusInternalServerError
//	- error: internal error
func DecodeDeleteProductResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body DeleteProductResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("products", "deleteProduct", err)
			}
			err = ValidateDeleteProductResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("products", "deleteProduct", err)
			}
			res := NewDeleteProductResultOK(&body)
			return res, nil
		case http.StatusInternalServerError:
			var (
				body DeleteProductUnknownErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("products", "deleteProduct", err)
			}
			err = ValidateDeleteProductUnknownErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("products", "deleteProduct", err)
			}
			return nil, NewDeleteProductUnknownError(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("products", "deleteProduct", resp.StatusCode, string(body))
		}
	}
}

// BuildCreateProductRequest instantiates a HTTP request object with method and
// path set to call the "products" service "createProduct" endpoint
func (c *Client) BuildCreateProductRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateProductProductsPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("products", "createProduct", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateProductRequest returns an encoder for requests sent to the
// products createProduct server.
func EncodeCreateProductRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*products.CreateProductPayload)
		if !ok {
			return goahttp.ErrInvalidType("products", "createProduct", "*products.CreateProductPayload", v)
		}
		if p.Oauth != nil {
			head := *p.Oauth
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		if p.JWTToken != nil {
			head := *p.JWTToken
			req.Header.Set("jwtToken", head)
		}
		body := NewCreateProductRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("products", "createProduct", err)
		}
		return nil
	}
}

// DecodeCreateProductResponse returns a decoder for responses returned by the
// products createProduct endpoint. restoreBody controls whether the response
// body should be restored after having been read.
// DecodeCreateProductResponse may return the following errors:
//	- "unknown_error" (type *products.UnknownError): http.StatusInternalServerError
//	- error: internal error
func DecodeCreateProductResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusCreated:
			var (
				body CreateProductResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("products", "createProduct", err)
			}
			err = ValidateCreateProductResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("products", "createProduct", err)
			}
			res := NewCreateProductResultCreated(&body)
			return res, nil
		case http.StatusInternalServerError:
			var (
				body CreateProductUnknownErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("products", "createProduct", err)
			}
			err = ValidateCreateProductUnknownErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("products", "createProduct", err)
			}
			return nil, NewCreateProductUnknownError(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("products", "createProduct", resp.StatusCode, string(body))
		}
	}
}

// BuildUpdateProductRequest instantiates a HTTP request object with method and
// path set to call the "products" service "updateProduct" endpoint
func (c *Client) BuildUpdateProductRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*products.UpdateProductPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("products", "updateProduct", "*products.UpdateProductPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UpdateProductProductsPath(id)}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("products", "updateProduct", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUpdateProductRequest returns an encoder for requests sent to the
// products updateProduct server.
func EncodeUpdateProductRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*products.UpdateProductPayload)
		if !ok {
			return goahttp.ErrInvalidType("products", "updateProduct", "*products.UpdateProductPayload", v)
		}
		if p.Oauth != nil {
			head := *p.Oauth
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		if p.JWTToken != nil {
			head := *p.JWTToken
			req.Header.Set("jwtToken", head)
		}
		body := NewUpdateProductRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("products", "updateProduct", err)
		}
		return nil
	}
}

// DecodeUpdateProductResponse returns a decoder for responses returned by the
// products updateProduct endpoint. restoreBody controls whether the response
// body should be restored after having been read.
// DecodeUpdateProductResponse may return the following errors:
//	- "unknown_error" (type *products.UnknownError): http.StatusInternalServerError
//	- error: internal error
func DecodeUpdateProductResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body UpdateProductResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("products", "updateProduct", err)
			}
			err = ValidateUpdateProductResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("products", "updateProduct", err)
			}
			res := NewUpdateProductResultOK(&body)
			return res, nil
		case http.StatusInternalServerError:
			var (
				body UpdateProductUnknownErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("products", "updateProduct", err)
			}
			err = ValidateUpdateProductUnknownErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("products", "updateProduct", err)
			}
			return nil, NewUpdateProductUnknownError(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("products", "updateProduct", resp.StatusCode, string(body))
		}
	}
}

// BuildGetProductRequest instantiates a HTTP request object with method and
// path set to call the "products" service "getProduct" endpoint
func (c *Client) BuildGetProductRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*products.GetProductPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("products", "getProduct", "*products.GetProductPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetProductProductsPath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("products", "getProduct", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeGetProductRequest returns an encoder for requests sent to the products
// getProduct server.
func EncodeGetProductRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*products.GetProductPayload)
		if !ok {
			return goahttp.ErrInvalidType("products", "getProduct", "*products.GetProductPayload", v)
		}
		if p.Oauth != nil {
			head := *p.Oauth
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		if p.JWTToken != nil {
			head := *p.JWTToken
			req.Header.Set("jwtToken", head)
		}
		return nil
	}
}

// DecodeGetProductResponse returns a decoder for responses returned by the
// products getProduct endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeGetProductResponse may return the following errors:
//	- "unknown_error" (type *products.UnknownError): http.StatusInternalServerError
//	- error: internal error
func DecodeGetProductResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetProductResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("products", "getProduct", err)
			}
			err = ValidateGetProductResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("products", "getProduct", err)
			}
			res := NewGetProductResultOK(&body)
			return res, nil
		case http.StatusInternalServerError:
			var (
				body GetProductUnknownErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("products", "getProduct", err)
			}
			err = ValidateGetProductUnknownErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("products", "getProduct", err)
			}
			return nil, NewGetProductUnknownError(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("products", "getProduct", resp.StatusCode, string(body))
		}
	}
}

// unmarshalResProductResponseBodyToProductsResProduct builds a value of type
// *products.ResProduct from a value of type *ResProductResponseBody.
func unmarshalResProductResponseBodyToProductsResProduct(v *ResProductResponseBody) *products.ResProduct {
	res := &products.ResProduct{
		ID:       *v.ID,
		Name:     *v.Name,
		Price:    *v.Price,
		Cover:    *v.Cover,
		Category: *v.Category,
	}

	return res
}

// marshalProductsPayloadProductToPayloadProductRequestBody builds a value of
// type *PayloadProductRequestBody from a value of type
// *products.PayloadProduct.
func marshalProductsPayloadProductToPayloadProductRequestBody(v *products.PayloadProduct) *PayloadProductRequestBody {
	res := &PayloadProductRequestBody{
		Name:     v.Name,
		Price:    v.Price,
		Cover:    v.Cover,
		Category: v.Category,
	}

	return res
}

// marshalPayloadProductRequestBodyToProductsPayloadProduct builds a value of
// type *products.PayloadProduct from a value of type
// *PayloadProductRequestBody.
func marshalPayloadProductRequestBodyToProductsPayloadProduct(v *PayloadProductRequestBody) *products.PayloadProduct {
	res := &products.PayloadProduct{
		Name:     v.Name,
		Price:    v.Price,
		Cover:    v.Cover,
		Category: v.Category,
	}

	return res
}
