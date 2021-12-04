// Code generated by goa v3.5.2, DO NOT EDIT.
//
// boUsers client HTTP transport
//
// Command:
// $ goa gen api_crud/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the boUsers service endpoint HTTP clients.
type Client struct {
	// GetAllusers Doer is the HTTP client used to make requests to the getAllusers
	// endpoint.
	GetAllusersDoer goahttp.Doer

	// DeleteUser Doer is the HTTP client used to make requests to the deleteUser
	// endpoint.
	DeleteUserDoer goahttp.Doer

	// CreateUser Doer is the HTTP client used to make requests to the createUser
	// endpoint.
	CreateUserDoer goahttp.Doer

	// UpdateUser Doer is the HTTP client used to make requests to the updateUser
	// endpoint.
	UpdateUserDoer goahttp.Doer

	// GetUser Doer is the HTTP client used to make requests to the getUser
	// endpoint.
	GetUserDoer goahttp.Doer

	// DeleteManyUsers Doer is the HTTP client used to make requests to the
	// deleteManyUsers endpoint.
	DeleteManyUsersDoer goahttp.Doer

	// NewPassword Doer is the HTTP client used to make requests to the newPassword
	// endpoint.
	NewPasswordDoer goahttp.Doer

	// CORS Doer is the HTTP client used to make requests to the  endpoint.
	CORSDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the boUsers service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		GetAllusersDoer:     doer,
		DeleteUserDoer:      doer,
		CreateUserDoer:      doer,
		UpdateUserDoer:      doer,
		GetUserDoer:         doer,
		DeleteManyUsersDoer: doer,
		NewPasswordDoer:     doer,
		CORSDoer:            doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// GetAllusers returns an endpoint that makes HTTP requests to the boUsers
// service getAllusers server.
func (c *Client) GetAllusers() goa.Endpoint {
	var (
		encodeRequest  = EncodeGetAllusersRequest(c.encoder)
		decodeResponse = DecodeGetAllusersResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildGetAllusersRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetAllusersDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("boUsers", "getAllusers", err)
		}
		return decodeResponse(resp)
	}
}

// DeleteUser returns an endpoint that makes HTTP requests to the boUsers
// service deleteUser server.
func (c *Client) DeleteUser() goa.Endpoint {
	var (
		encodeRequest  = EncodeDeleteUserRequest(c.encoder)
		decodeResponse = DecodeDeleteUserResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildDeleteUserRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.DeleteUserDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("boUsers", "deleteUser", err)
		}
		return decodeResponse(resp)
	}
}

// CreateUser returns an endpoint that makes HTTP requests to the boUsers
// service createUser server.
func (c *Client) CreateUser() goa.Endpoint {
	var (
		encodeRequest  = EncodeCreateUserRequest(c.encoder)
		decodeResponse = DecodeCreateUserResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildCreateUserRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.CreateUserDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("boUsers", "createUser", err)
		}
		return decodeResponse(resp)
	}
}

// UpdateUser returns an endpoint that makes HTTP requests to the boUsers
// service updateUser server.
func (c *Client) UpdateUser() goa.Endpoint {
	var (
		encodeRequest  = EncodeUpdateUserRequest(c.encoder)
		decodeResponse = DecodeUpdateUserResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildUpdateUserRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.UpdateUserDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("boUsers", "updateUser", err)
		}
		return decodeResponse(resp)
	}
}

// GetUser returns an endpoint that makes HTTP requests to the boUsers service
// getUser server.
func (c *Client) GetUser() goa.Endpoint {
	var (
		encodeRequest  = EncodeGetUserRequest(c.encoder)
		decodeResponse = DecodeGetUserResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildGetUserRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetUserDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("boUsers", "getUser", err)
		}
		return decodeResponse(resp)
	}
}

// DeleteManyUsers returns an endpoint that makes HTTP requests to the boUsers
// service deleteManyUsers server.
func (c *Client) DeleteManyUsers() goa.Endpoint {
	var (
		encodeRequest  = EncodeDeleteManyUsersRequest(c.encoder)
		decodeResponse = DecodeDeleteManyUsersResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildDeleteManyUsersRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.DeleteManyUsersDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("boUsers", "deleteManyUsers", err)
		}
		return decodeResponse(resp)
	}
}

// NewPassword returns an endpoint that makes HTTP requests to the boUsers
// service newPassword server.
func (c *Client) NewPassword() goa.Endpoint {
	var (
		encodeRequest  = EncodeNewPasswordRequest(c.encoder)
		decodeResponse = DecodeNewPasswordResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildNewPasswordRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.NewPasswordDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("boUsers", "newPassword", err)
		}
		return decodeResponse(resp)
	}
}