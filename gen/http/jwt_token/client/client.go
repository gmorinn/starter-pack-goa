// Code generated by goa v3.5.2, DO NOT EDIT.
//
// jwtToken client HTTP transport
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

// Client lists the jwtToken service endpoint HTTP clients.
type Client struct {
	// Signup Doer is the HTTP client used to make requests to the signup endpoint.
	SignupDoer goahttp.Doer

	// Signin Doer is the HTTP client used to make requests to the signin endpoint.
	SigninDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the jwtToken service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		SignupDoer:          doer,
		SigninDoer:          doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Signup returns an endpoint that makes HTTP requests to the jwtToken service
// signup server.
func (c *Client) Signup() goa.Endpoint {
	var (
		encodeRequest  = EncodeSignupRequest(c.encoder)
		decodeResponse = DecodeSignupResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildSignupRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.SignupDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("jwtToken", "signup", err)
		}
		return decodeResponse(resp)
	}
}

// Signin returns an endpoint that makes HTTP requests to the jwtToken service
// signin server.
func (c *Client) Signin() goa.Endpoint {
	var (
		encodeRequest  = EncodeSigninRequest(c.encoder)
		decodeResponse = DecodeSigninResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildSigninRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.SigninDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("jwtToken", "signin", err)
		}
		return decodeResponse(resp)
	}
}
