// Code generated by goa v3.5.2, DO NOT EDIT.
//
// products endpoints
//
// Command:
// $ goa gen api_crud/design

package products

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "products" service endpoints.
type Endpoints struct {
	GetAllProductsByCategory goa.Endpoint
	GetAllProducts           goa.Endpoint
	GetProduct               goa.Endpoint
}

// NewEndpoints wraps the methods of the "products" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		GetAllProductsByCategory: NewGetAllProductsByCategoryEndpoint(s, a.OAuth2Auth),
		GetAllProducts:           NewGetAllProductsEndpoint(s, a.OAuth2Auth),
		GetProduct:               NewGetProductEndpoint(s, a.OAuth2Auth),
	}
}

// Use applies the given middleware to all the "products" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.GetAllProductsByCategory = m(e.GetAllProductsByCategory)
	e.GetAllProducts = m(e.GetAllProducts)
	e.GetProduct = m(e.GetProduct)
}

// NewGetAllProductsByCategoryEndpoint returns an endpoint function that calls
// the method "getAllProductsByCategory" of service "products".
func NewGetAllProductsByCategoryEndpoint(s Service, authOAuth2Fn security.AuthOAuth2Func) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*GetAllProductsByCategoryPayload)
		var err error
		sc := security.OAuth2Scheme{
			Name:           "OAuth2",
			Scopes:         []string{"api:read"},
			RequiredScopes: []string{},
			Flows: []*security.OAuthFlow{
				&security.OAuthFlow{
					Type:       "client_credentials",
					TokenURL:   "/authorization",
					RefreshURL: "/refresh",
				},
			},
		}
		var token string
		if p.Oauth != nil {
			token = *p.Oauth
		}
		ctx, err = authOAuth2Fn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return s.GetAllProductsByCategory(ctx, p)
	}
}

// NewGetAllProductsEndpoint returns an endpoint function that calls the method
// "getAllProducts" of service "products".
func NewGetAllProductsEndpoint(s Service, authOAuth2Fn security.AuthOAuth2Func) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*GetAllProductsPayload)
		var err error
		sc := security.OAuth2Scheme{
			Name:           "OAuth2",
			Scopes:         []string{"api:read"},
			RequiredScopes: []string{},
			Flows: []*security.OAuthFlow{
				&security.OAuthFlow{
					Type:       "client_credentials",
					TokenURL:   "/authorization",
					RefreshURL: "/refresh",
				},
			},
		}
		var token string
		if p.Oauth != nil {
			token = *p.Oauth
		}
		ctx, err = authOAuth2Fn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return s.GetAllProducts(ctx, p)
	}
}

// NewGetProductEndpoint returns an endpoint function that calls the method
// "getProduct" of service "products".
func NewGetProductEndpoint(s Service, authOAuth2Fn security.AuthOAuth2Func) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*GetProductPayload)
		var err error
		sc := security.OAuth2Scheme{
			Name:           "OAuth2",
			Scopes:         []string{"api:read"},
			RequiredScopes: []string{},
			Flows: []*security.OAuthFlow{
				&security.OAuthFlow{
					Type:       "client_credentials",
					TokenURL:   "/authorization",
					RefreshURL: "/refresh",
				},
			},
		}
		var token string
		if p.Oauth != nil {
			token = *p.Oauth
		}
		ctx, err = authOAuth2Fn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return s.GetProduct(ctx, p)
	}
}
