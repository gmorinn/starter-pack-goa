// Code generated by goa v3.5.2, DO NOT EDIT.
//
// users client HTTP transport
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

// Client lists the users service endpoint HTTP clients.
type Client struct {
	// GetUser Doer is the HTTP client used to make requests to the getUser
	// endpoint.
	GetUserDoer goahttp.Doer

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

// NewClient instantiates HTTP clients for all the users service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		GetUserDoer:         doer,
		CORSDoer:            doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// GetUser returns an endpoint that makes HTTP requests to the users service
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
			return nil, goahttp.ErrRequestError("users", "getUser", err)
		}
		return decodeResponse(resp)
	}
}
