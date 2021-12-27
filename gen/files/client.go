// Code generated by goa v3.5.2, DO NOT EDIT.
//
// files client
//
// Command:
// $ goa gen api_crud/design

package files

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "files" service client.
type Client struct {
	ImportFileEndpoint goa.Endpoint
}

// NewClient initializes a "files" service client given the endpoints.
func NewClient(importFile goa.Endpoint) *Client {
	return &Client{
		ImportFileEndpoint: importFile,
	}
}

// ImportFile calls the "importFile" endpoint of the "files" service.
func (c *Client) ImportFile(ctx context.Context, p *ImportFilePayload) (res *ImportFileResult, err error) {
	var ires interface{}
	ires, err = c.ImportFileEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*ImportFileResult), nil
}