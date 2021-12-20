// Code generated by goa v3.5.2, DO NOT EDIT.
//
// boProducts HTTP server encoders and decoders
//
// Command:
// $ goa gen api_crud/design

package server

import (
	boproducts "api_crud/gen/bo_products"
	"context"
	"io"
	"net/http"
	"strconv"
	"strings"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeGetAllProductsResponse returns an encoder for responses returned by
// the boProducts getAllProducts endpoint.
func EncodeGetAllProductsResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*boproducts.GetAllProductsResult)
		enc := encoder(ctx, w)
		body := NewGetAllProductsResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetAllProductsRequest returns a decoder for requests sent to the
// boProducts getAllProducts endpoint.
func DecodeGetAllProductsRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			offset    int32
			limit     int32
			field     string
			direction string
			oauth     *string
			jwtToken  *string
			err       error

			params = mux.Vars(r)
		)
		{
			offsetRaw := params["offset"]
			v, err2 := strconv.ParseInt(offsetRaw, 10, 32)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("offset", offsetRaw, "integer"))
			}
			offset = int32(v)
		}
		{
			limitRaw := params["limit"]
			v, err2 := strconv.ParseInt(limitRaw, 10, 32)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("limit", limitRaw, "integer"))
			}
			limit = int32(v)
		}
		fieldRaw := r.URL.Query().Get("field")
		if fieldRaw != "" {
			field = fieldRaw
		} else {
			field = "name"
		}
		directionRaw := r.URL.Query().Get("direction")
		if directionRaw != "" {
			direction = directionRaw
		} else {
			direction = "asc"
		}
		if !(direction == "asc" || direction == "desc") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("direction", direction, []interface{}{"asc", "desc"}))
		}
		oauthRaw := r.Header.Get("Authorization")
		if oauthRaw != "" {
			oauth = &oauthRaw
		}
		jwtTokenRaw := r.Header.Get("jwtToken")
		if jwtTokenRaw != "" {
			jwtToken = &jwtTokenRaw
		}
		if err != nil {
			return nil, err
		}
		payload := NewGetAllProductsPayload(offset, limit, field, direction, oauth, jwtToken)
		if payload.Oauth != nil {
			if strings.Contains(*payload.Oauth, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Oauth, " ", 2)[1]
				payload.Oauth = &cred
			}
		}
		if payload.JWTToken != nil {
			if strings.Contains(*payload.JWTToken, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.JWTToken, " ", 2)[1]
				payload.JWTToken = &cred
			}
		}

		return payload, nil
	}
}

// EncodeGetAllProductsError returns an encoder for errors returned by the
// getAllProducts boProducts endpoint.
func EncodeGetAllProductsError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unknown_error":
			res := v.(*boproducts.UnknownError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewGetAllProductsUnknownErrorResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeGetAllProductsByCategoryResponse returns an encoder for responses
// returned by the boProducts getAllProductsByCategory endpoint.
func EncodeGetAllProductsByCategoryResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*boproducts.GetAllProductsByCategoryResult)
		enc := encoder(ctx, w)
		body := NewGetAllProductsByCategoryResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetAllProductsByCategoryRequest returns a decoder for requests sent to
// the boProducts getAllProductsByCategory endpoint.
func DecodeGetAllProductsByCategoryRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			category string
			oauth    *string
			jwtToken *string
			err      error

			params = mux.Vars(r)
		)
		category = params["category"]
		if !(category == "men" || category == "women" || category == "hat" || category == "jacket" || category == "sneaker" || category == "nothing") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("category", category, []interface{}{"men", "women", "hat", "jacket", "sneaker", "nothing"}))
		}
		oauthRaw := r.Header.Get("Authorization")
		if oauthRaw != "" {
			oauth = &oauthRaw
		}
		jwtTokenRaw := r.Header.Get("jwtToken")
		if jwtTokenRaw != "" {
			jwtToken = &jwtTokenRaw
		}
		if err != nil {
			return nil, err
		}
		payload := NewGetAllProductsByCategoryPayload(category, oauth, jwtToken)
		if payload.Oauth != nil {
			if strings.Contains(*payload.Oauth, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Oauth, " ", 2)[1]
				payload.Oauth = &cred
			}
		}
		if payload.JWTToken != nil {
			if strings.Contains(*payload.JWTToken, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.JWTToken, " ", 2)[1]
				payload.JWTToken = &cred
			}
		}

		return payload, nil
	}
}

// EncodeGetAllProductsByCategoryError returns an encoder for errors returned
// by the getAllProductsByCategory boProducts endpoint.
func EncodeGetAllProductsByCategoryError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unknown_error":
			res := v.(*boproducts.UnknownError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewGetAllProductsByCategoryUnknownErrorResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeDeleteProductResponse returns an encoder for responses returned by the
// boProducts deleteProduct endpoint.
func EncodeDeleteProductResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*boproducts.DeleteProductResult)
		enc := encoder(ctx, w)
		body := NewDeleteProductResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeDeleteProductRequest returns a decoder for requests sent to the
// boProducts deleteProduct endpoint.
func DecodeDeleteProductRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id       string
			oauth    *string
			jwtToken *string
			err      error

			params = mux.Vars(r)
		)
		id = params["id"]
		err = goa.MergeErrors(err, goa.ValidateFormat("id", id, goa.FormatUUID))

		oauthRaw := r.Header.Get("Authorization")
		if oauthRaw != "" {
			oauth = &oauthRaw
		}
		jwtTokenRaw := r.Header.Get("jwtToken")
		if jwtTokenRaw != "" {
			jwtToken = &jwtTokenRaw
		}
		if err != nil {
			return nil, err
		}
		payload := NewDeleteProductPayload(id, oauth, jwtToken)
		if payload.Oauth != nil {
			if strings.Contains(*payload.Oauth, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Oauth, " ", 2)[1]
				payload.Oauth = &cred
			}
		}
		if payload.JWTToken != nil {
			if strings.Contains(*payload.JWTToken, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.JWTToken, " ", 2)[1]
				payload.JWTToken = &cred
			}
		}

		return payload, nil
	}
}

// EncodeDeleteProductError returns an encoder for errors returned by the
// deleteProduct boProducts endpoint.
func EncodeDeleteProductError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unknown_error":
			res := v.(*boproducts.UnknownError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewDeleteProductUnknownErrorResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeCreateProductResponse returns an encoder for responses returned by the
// boProducts createProduct endpoint.
func EncodeCreateProductResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*boproducts.CreateProductResult)
		enc := encoder(ctx, w)
		body := NewCreateProductResponseBody(res)
		w.WriteHeader(http.StatusCreated)
		return enc.Encode(body)
	}
}

// DecodeCreateProductRequest returns a decoder for requests sent to the
// boProducts createProduct endpoint.
func DecodeCreateProductRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body CreateProductRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateCreateProductRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			oauth    *string
			jwtToken *string
		)
		oauthRaw := r.Header.Get("Authorization")
		if oauthRaw != "" {
			oauth = &oauthRaw
		}
		jwtTokenRaw := r.Header.Get("jwtToken")
		if jwtTokenRaw != "" {
			jwtToken = &jwtTokenRaw
		}
		payload := NewCreateProductPayload(&body, oauth, jwtToken)
		if payload.Oauth != nil {
			if strings.Contains(*payload.Oauth, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Oauth, " ", 2)[1]
				payload.Oauth = &cred
			}
		}
		if payload.JWTToken != nil {
			if strings.Contains(*payload.JWTToken, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.JWTToken, " ", 2)[1]
				payload.JWTToken = &cred
			}
		}

		return payload, nil
	}
}

// EncodeCreateProductError returns an encoder for errors returned by the
// createProduct boProducts endpoint.
func EncodeCreateProductError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unknown_error":
			res := v.(*boproducts.UnknownError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewCreateProductUnknownErrorResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeUpdateProductResponse returns an encoder for responses returned by the
// boProducts updateProduct endpoint.
func EncodeUpdateProductResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*boproducts.UpdateProductResult)
		enc := encoder(ctx, w)
		body := NewUpdateProductResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeUpdateProductRequest returns a decoder for requests sent to the
// boProducts updateProduct endpoint.
func DecodeUpdateProductRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body UpdateProductRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateUpdateProductRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			id       string
			oauth    *string
			jwtToken *string

			params = mux.Vars(r)
		)
		id = params["id"]
		err = goa.MergeErrors(err, goa.ValidateFormat("id", id, goa.FormatUUID))

		oauthRaw := r.Header.Get("Authorization")
		if oauthRaw != "" {
			oauth = &oauthRaw
		}
		jwtTokenRaw := r.Header.Get("jwtToken")
		if jwtTokenRaw != "" {
			jwtToken = &jwtTokenRaw
		}
		if err != nil {
			return nil, err
		}
		payload := NewUpdateProductPayload(&body, id, oauth, jwtToken)
		if payload.Oauth != nil {
			if strings.Contains(*payload.Oauth, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Oauth, " ", 2)[1]
				payload.Oauth = &cred
			}
		}
		if payload.JWTToken != nil {
			if strings.Contains(*payload.JWTToken, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.JWTToken, " ", 2)[1]
				payload.JWTToken = &cred
			}
		}

		return payload, nil
	}
}

// EncodeUpdateProductError returns an encoder for errors returned by the
// updateProduct boProducts endpoint.
func EncodeUpdateProductError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unknown_error":
			res := v.(*boproducts.UnknownError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewUpdateProductUnknownErrorResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeDeleteManyProductsResponse returns an encoder for responses returned
// by the boProducts deleteManyProducts endpoint.
func EncodeDeleteManyProductsResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*boproducts.DeleteManyProductsResult)
		enc := encoder(ctx, w)
		body := NewDeleteManyProductsResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeDeleteManyProductsRequest returns a decoder for requests sent to the
// boProducts deleteManyProducts endpoint.
func DecodeDeleteManyProductsRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body DeleteManyProductsRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateDeleteManyProductsRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			oauth    *string
			jwtToken *string
		)
		oauthRaw := r.Header.Get("Authorization")
		if oauthRaw != "" {
			oauth = &oauthRaw
		}
		jwtTokenRaw := r.Header.Get("jwtToken")
		if jwtTokenRaw != "" {
			jwtToken = &jwtTokenRaw
		}
		payload := NewDeleteManyProductsPayload(&body, oauth, jwtToken)
		if payload.Oauth != nil {
			if strings.Contains(*payload.Oauth, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Oauth, " ", 2)[1]
				payload.Oauth = &cred
			}
		}
		if payload.JWTToken != nil {
			if strings.Contains(*payload.JWTToken, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.JWTToken, " ", 2)[1]
				payload.JWTToken = &cred
			}
		}

		return payload, nil
	}
}

// EncodeDeleteManyProductsError returns an encoder for errors returned by the
// deleteManyProducts boProducts endpoint.
func EncodeDeleteManyProductsError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unknown_error":
			res := v.(*boproducts.UnknownError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewDeleteManyProductsUnknownErrorResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeGetProductResponse returns an encoder for responses returned by the
// boProducts getProduct endpoint.
func EncodeGetProductResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*boproducts.GetProductResult)
		enc := encoder(ctx, w)
		body := NewGetProductResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetProductRequest returns a decoder for requests sent to the
// boProducts getProduct endpoint.
func DecodeGetProductRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id       string
			oauth    *string
			jwtToken *string
			err      error

			params = mux.Vars(r)
		)
		id = params["id"]
		err = goa.MergeErrors(err, goa.ValidateFormat("id", id, goa.FormatUUID))

		oauthRaw := r.Header.Get("Authorization")
		if oauthRaw != "" {
			oauth = &oauthRaw
		}
		jwtTokenRaw := r.Header.Get("jwtToken")
		if jwtTokenRaw != "" {
			jwtToken = &jwtTokenRaw
		}
		if err != nil {
			return nil, err
		}
		payload := NewGetProductPayload(id, oauth, jwtToken)
		if payload.Oauth != nil {
			if strings.Contains(*payload.Oauth, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Oauth, " ", 2)[1]
				payload.Oauth = &cred
			}
		}
		if payload.JWTToken != nil {
			if strings.Contains(*payload.JWTToken, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.JWTToken, " ", 2)[1]
				payload.JWTToken = &cred
			}
		}

		return payload, nil
	}
}

// EncodeGetProductError returns an encoder for errors returned by the
// getProduct boProducts endpoint.
func EncodeGetProductError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unknown_error":
			res := v.(*boproducts.UnknownError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewGetProductUnknownErrorResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// marshalBoproductsResBoProductToResBoProductResponseBody builds a value of
// type *ResBoProductResponseBody from a value of type *boproducts.ResBoProduct.
func marshalBoproductsResBoProductToResBoProductResponseBody(v *boproducts.ResBoProduct) *ResBoProductResponseBody {
	res := &ResBoProductResponseBody{
		ID:       v.ID,
		Name:     v.Name,
		Price:    v.Price,
		Cover:    v.Cover,
		Category: v.Category,
	}

	return res
}

// unmarshalPayloadProductRequestBodyToBoproductsPayloadProduct builds a value
// of type *boproducts.PayloadProduct from a value of type
// *PayloadProductRequestBody.
func unmarshalPayloadProductRequestBodyToBoproductsPayloadProduct(v *PayloadProductRequestBody) *boproducts.PayloadProduct {
	res := &boproducts.PayloadProduct{
		Name:     *v.Name,
		Price:    *v.Price,
		Cover:    *v.Cover,
		Category: *v.Category,
	}

	return res
}
