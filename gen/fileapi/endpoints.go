// Code generated by goa v3.10.2, DO NOT EDIT.
//
// fileapi endpoints
//
// Command:
// $ goa gen starter-pack-goa/design

package fileapi

import (
	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "fileapi" service endpoints.
type Endpoints struct {
}

// NewEndpoints wraps the methods of the "fileapi" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{}
}

// Use applies the given middleware to all the "fileapi" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
}
