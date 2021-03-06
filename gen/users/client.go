// Code generated by goa v3.5.2, DO NOT EDIT.
//
// users client
//
// Command:
// $ goa gen api_crud/design

package users

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "users" service client.
type Client struct {
	GetUserEndpoint goa.Endpoint
}

// NewClient initializes a "users" service client given the endpoints.
func NewClient(getUser goa.Endpoint) *Client {
	return &Client{
		GetUserEndpoint: getUser,
	}
}

// GetUser calls the "getUser" endpoint of the "users" service.
func (c *Client) GetUser(ctx context.Context, p *GetUserPayload) (res *GetUserResult, err error) {
	var ires interface{}
	ires, err = c.GetUserEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*GetUserResult), nil
}
