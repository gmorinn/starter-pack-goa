// Code generated by goa v3.5.2, DO NOT EDIT.
//
// boUsers HTTP client encoders and decoders
//
// Command:
// $ goa gen api_crud/design

package client

import (
	bousers "api_crud/gen/bo_users"
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	goahttp "goa.design/goa/v3/http"
)

// BuildGetAllusersRequest instantiates a HTTP request object with method and
// path set to call the "boUsers" service "getAllusers" endpoint
func (c *Client) BuildGetAllusersRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		offset int32
		limit  int32
	)
	{
		p, ok := v.(*bousers.GetAllusersPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("boUsers", "getAllusers", "*bousers.GetAllusersPayload", v)
		}
		offset = p.Offset
		limit = p.Limit
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetAllusersBoUsersPath(offset, limit)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("boUsers", "getAllusers", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeGetAllusersRequest returns an encoder for requests sent to the boUsers
// getAllusers server.
func EncodeGetAllusersRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*bousers.GetAllusersPayload)
		if !ok {
			return goahttp.ErrInvalidType("boUsers", "getAllusers", "*bousers.GetAllusersPayload", v)
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
		values := req.URL.Query()
		values.Add("field", p.Field)
		values.Add("direction", p.Direction)
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeGetAllusersResponse returns a decoder for responses returned by the
// boUsers getAllusers endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeGetAllusersResponse may return the following errors:
//	- "unknown_error" (type *bousers.UnknownError): http.StatusInternalServerError
//	- error: internal error
func DecodeGetAllusersResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				body GetAllusersResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("boUsers", "getAllusers", err)
			}
			err = ValidateGetAllusersResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("boUsers", "getAllusers", err)
			}
			res := NewGetAllusersResultOK(&body)
			return res, nil
		case http.StatusInternalServerError:
			var (
				body GetAllusersUnknownErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("boUsers", "getAllusers", err)
			}
			err = ValidateGetAllusersUnknownErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("boUsers", "getAllusers", err)
			}
			return nil, NewGetAllusersUnknownError(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("boUsers", "getAllusers", resp.StatusCode, string(body))
		}
	}
}

// BuildDeleteUserRequest instantiates a HTTP request object with method and
// path set to call the "boUsers" service "deleteUser" endpoint
func (c *Client) BuildDeleteUserRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*bousers.DeleteUserPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("boUsers", "deleteUser", "*bousers.DeleteUserPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DeleteUserBoUsersPath(id)}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("boUsers", "deleteUser", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeDeleteUserRequest returns an encoder for requests sent to the boUsers
// deleteUser server.
func EncodeDeleteUserRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*bousers.DeleteUserPayload)
		if !ok {
			return goahttp.ErrInvalidType("boUsers", "deleteUser", "*bousers.DeleteUserPayload", v)
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

// DecodeDeleteUserResponse returns a decoder for responses returned by the
// boUsers deleteUser endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeDeleteUserResponse may return the following errors:
//	- "unknown_error" (type *bousers.UnknownError): http.StatusInternalServerError
//	- error: internal error
func DecodeDeleteUserResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				body DeleteUserResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("boUsers", "deleteUser", err)
			}
			err = ValidateDeleteUserResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("boUsers", "deleteUser", err)
			}
			res := NewDeleteUserResultOK(&body)
			return res, nil
		case http.StatusInternalServerError:
			var (
				body DeleteUserUnknownErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("boUsers", "deleteUser", err)
			}
			err = ValidateDeleteUserUnknownErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("boUsers", "deleteUser", err)
			}
			return nil, NewDeleteUserUnknownError(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("boUsers", "deleteUser", resp.StatusCode, string(body))
		}
	}
}

// BuildCreateUserRequest instantiates a HTTP request object with method and
// path set to call the "boUsers" service "createUser" endpoint
func (c *Client) BuildCreateUserRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateUserBoUsersPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("boUsers", "createUser", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateUserRequest returns an encoder for requests sent to the boUsers
// createUser server.
func EncodeCreateUserRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*bousers.CreateUserPayload)
		if !ok {
			return goahttp.ErrInvalidType("boUsers", "createUser", "*bousers.CreateUserPayload", v)
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
		body := NewCreateUserRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("boUsers", "createUser", err)
		}
		return nil
	}
}

// DecodeCreateUserResponse returns a decoder for responses returned by the
// boUsers createUser endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeCreateUserResponse may return the following errors:
//	- "unknown_error" (type *bousers.UnknownError): http.StatusInternalServerError
//	- error: internal error
func DecodeCreateUserResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				body CreateUserResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("boUsers", "createUser", err)
			}
			err = ValidateCreateUserResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("boUsers", "createUser", err)
			}
			res := NewCreateUserResultCreated(&body)
			return res, nil
		case http.StatusInternalServerError:
			var (
				body CreateUserUnknownErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("boUsers", "createUser", err)
			}
			err = ValidateCreateUserUnknownErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("boUsers", "createUser", err)
			}
			return nil, NewCreateUserUnknownError(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("boUsers", "createUser", resp.StatusCode, string(body))
		}
	}
}

// BuildUpdateUserRequest instantiates a HTTP request object with method and
// path set to call the "boUsers" service "updateUser" endpoint
func (c *Client) BuildUpdateUserRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*bousers.UpdateUserPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("boUsers", "updateUser", "*bousers.UpdateUserPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UpdateUserBoUsersPath(id)}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("boUsers", "updateUser", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUpdateUserRequest returns an encoder for requests sent to the boUsers
// updateUser server.
func EncodeUpdateUserRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*bousers.UpdateUserPayload)
		if !ok {
			return goahttp.ErrInvalidType("boUsers", "updateUser", "*bousers.UpdateUserPayload", v)
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
		body := NewUpdateUserRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("boUsers", "updateUser", err)
		}
		return nil
	}
}

// DecodeUpdateUserResponse returns a decoder for responses returned by the
// boUsers updateUser endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeUpdateUserResponse may return the following errors:
//	- "unknown_error" (type *bousers.UnknownError): http.StatusInternalServerError
//	- error: internal error
func DecodeUpdateUserResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				body UpdateUserResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("boUsers", "updateUser", err)
			}
			err = ValidateUpdateUserResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("boUsers", "updateUser", err)
			}
			res := NewUpdateUserResultOK(&body)
			return res, nil
		case http.StatusInternalServerError:
			var (
				body UpdateUserUnknownErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("boUsers", "updateUser", err)
			}
			err = ValidateUpdateUserUnknownErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("boUsers", "updateUser", err)
			}
			return nil, NewUpdateUserUnknownError(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("boUsers", "updateUser", resp.StatusCode, string(body))
		}
	}
}

// BuildGetUserRequest instantiates a HTTP request object with method and path
// set to call the "boUsers" service "getUser" endpoint
func (c *Client) BuildGetUserRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*bousers.GetUserPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("boUsers", "getUser", "*bousers.GetUserPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetUserBoUsersPath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("boUsers", "getUser", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeGetUserRequest returns an encoder for requests sent to the boUsers
// getUser server.
func EncodeGetUserRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*bousers.GetUserPayload)
		if !ok {
			return goahttp.ErrInvalidType("boUsers", "getUser", "*bousers.GetUserPayload", v)
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

// DecodeGetUserResponse returns a decoder for responses returned by the
// boUsers getUser endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeGetUserResponse may return the following errors:
//	- "unknown_error" (type *bousers.UnknownError): http.StatusInternalServerError
//	- error: internal error
func DecodeGetUserResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				body GetUserResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("boUsers", "getUser", err)
			}
			err = ValidateGetUserResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("boUsers", "getUser", err)
			}
			res := NewGetUserResultOK(&body)
			return res, nil
		case http.StatusInternalServerError:
			var (
				body GetUserUnknownErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("boUsers", "getUser", err)
			}
			err = ValidateGetUserUnknownErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("boUsers", "getUser", err)
			}
			return nil, NewGetUserUnknownError(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("boUsers", "getUser", resp.StatusCode, string(body))
		}
	}
}

// BuildDeleteManyUsersRequest instantiates a HTTP request object with method
// and path set to call the "boUsers" service "deleteManyUsers" endpoint
func (c *Client) BuildDeleteManyUsersRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DeleteManyUsersBoUsersPath()}
	req, err := http.NewRequest("PATCH", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("boUsers", "deleteManyUsers", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeDeleteManyUsersRequest returns an encoder for requests sent to the
// boUsers deleteManyUsers server.
func EncodeDeleteManyUsersRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*bousers.DeleteManyUsersPayload)
		if !ok {
			return goahttp.ErrInvalidType("boUsers", "deleteManyUsers", "*bousers.DeleteManyUsersPayload", v)
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
		body := NewDeleteManyUsersRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("boUsers", "deleteManyUsers", err)
		}
		return nil
	}
}

// DecodeDeleteManyUsersResponse returns a decoder for responses returned by
// the boUsers deleteManyUsers endpoint. restoreBody controls whether the
// response body should be restored after having been read.
// DecodeDeleteManyUsersResponse may return the following errors:
//	- "unknown_error" (type *bousers.UnknownError): http.StatusInternalServerError
//	- error: internal error
func DecodeDeleteManyUsersResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				body DeleteManyUsersResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("boUsers", "deleteManyUsers", err)
			}
			err = ValidateDeleteManyUsersResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("boUsers", "deleteManyUsers", err)
			}
			res := NewDeleteManyUsersResultOK(&body)
			return res, nil
		case http.StatusInternalServerError:
			var (
				body DeleteManyUsersUnknownErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("boUsers", "deleteManyUsers", err)
			}
			err = ValidateDeleteManyUsersUnknownErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("boUsers", "deleteManyUsers", err)
			}
			return nil, NewDeleteManyUsersUnknownError(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("boUsers", "deleteManyUsers", resp.StatusCode, string(body))
		}
	}
}

// BuildNewPasswordRequest instantiates a HTTP request object with method and
// path set to call the "boUsers" service "newPassword" endpoint
func (c *Client) BuildNewPasswordRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*bousers.NewPasswordPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("boUsers", "newPassword", "*bousers.NewPasswordPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: NewPasswordBoUsersPath(id)}
	req, err := http.NewRequest("PATCH", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("boUsers", "newPassword", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeNewPasswordRequest returns an encoder for requests sent to the boUsers
// newPassword server.
func EncodeNewPasswordRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*bousers.NewPasswordPayload)
		if !ok {
			return goahttp.ErrInvalidType("boUsers", "newPassword", "*bousers.NewPasswordPayload", v)
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
		body := NewNewPasswordRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("boUsers", "newPassword", err)
		}
		return nil
	}
}

// DecodeNewPasswordResponse returns a decoder for responses returned by the
// boUsers newPassword endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeNewPasswordResponse may return the following errors:
//	- "unknown_error" (type *bousers.UnknownError): http.StatusInternalServerError
//	- error: internal error
func DecodeNewPasswordResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				body NewPasswordResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("boUsers", "newPassword", err)
			}
			err = ValidateNewPasswordResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("boUsers", "newPassword", err)
			}
			res := NewNewPasswordResultOK(&body)
			return res, nil
		case http.StatusInternalServerError:
			var (
				body NewPasswordUnknownErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("boUsers", "newPassword", err)
			}
			err = ValidateNewPasswordUnknownErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("boUsers", "newPassword", err)
			}
			return nil, NewNewPasswordUnknownError(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("boUsers", "newPassword", resp.StatusCode, string(body))
		}
	}
}

// unmarshalResBoUserResponseBodyToBousersResBoUser builds a value of type
// *bousers.ResBoUser from a value of type *ResBoUserResponseBody.
func unmarshalResBoUserResponseBodyToBousersResBoUser(v *ResBoUserResponseBody) *bousers.ResBoUser {
	res := &bousers.ResBoUser{
		ID:        *v.ID,
		Firstname: v.Firstname,
		Lastname:  v.Lastname,
		Email:     *v.Email,
	}
	if v.Birthday != nil {
		res.Birthday = *v.Birthday
	}
	if v.Phone != nil {
		res.Phone = *v.Phone
	}
	if v.Role != nil {
		res.Role = *v.Role
	}
	if v.Birthday == nil {
		res.Birthday = ""
	}
	if v.Phone == nil {
		res.Phone = ""
	}
	if v.Role == nil {
		res.Role = "user"
	}

	return res
}

// marshalBousersPayloadUserToPayloadUserRequestBody builds a value of type
// *PayloadUserRequestBody from a value of type *bousers.PayloadUser.
func marshalBousersPayloadUserToPayloadUserRequestBody(v *bousers.PayloadUser) *PayloadUserRequestBody {
	res := &PayloadUserRequestBody{
		Firstname: v.Firstname,
		Lastname:  v.Lastname,
		Email:     v.Email,
		Birthday:  v.Birthday,
		Role:      v.Role,
		Phone:     v.Phone,
	}
	{
		var zero string
		if res.Birthday == zero {
			res.Birthday = ""
		}
	}
	{
		var zero string
		if res.Role == zero {
			res.Role = "user"
		}
	}
	{
		var zero string
		if res.Phone == zero {
			res.Phone = ""
		}
	}

	return res
}

// marshalPayloadUserRequestBodyToBousersPayloadUser builds a value of type
// *bousers.PayloadUser from a value of type *PayloadUserRequestBody.
func marshalPayloadUserRequestBodyToBousersPayloadUser(v *PayloadUserRequestBody) *bousers.PayloadUser {
	res := &bousers.PayloadUser{
		Firstname: v.Firstname,
		Lastname:  v.Lastname,
		Email:     v.Email,
		Birthday:  v.Birthday,
		Role:      v.Role,
		Phone:     v.Phone,
	}
	{
		var zero string
		if res.Birthday == zero {
			res.Birthday = ""
		}
	}
	{
		var zero string
		if res.Role == zero {
			res.Role = "user"
		}
	}
	{
		var zero string
		if res.Phone == zero {
			res.Phone = ""
		}
	}

	return res
}
