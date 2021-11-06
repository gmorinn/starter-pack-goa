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
	DeleteProduct            goa.Endpoint
	CreateProduct            goa.Endpoint
	UpdateProduct            goa.Endpoint
	GetProduct               goa.Endpoint
}

// NewEndpoints wraps the methods of the "products" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		GetAllProductsByCategory: NewGetAllProductsByCategoryEndpoint(s, a.OAuth2Auth, a.JWTAuth),
		DeleteProduct:            NewDeleteProductEndpoint(s, a.OAuth2Auth, a.JWTAuth),
		CreateProduct:            NewCreateProductEndpoint(s, a.OAuth2Auth, a.JWTAuth),
		UpdateProduct:            NewUpdateProductEndpoint(s, a.OAuth2Auth, a.JWTAuth),
		GetProduct:               NewGetProductEndpoint(s, a.OAuth2Auth, a.JWTAuth),
	}
}

// Use applies the given middleware to all the "products" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.GetAllProductsByCategory = m(e.GetAllProductsByCategory)
	e.DeleteProduct = m(e.DeleteProduct)
	e.CreateProduct = m(e.CreateProduct)
	e.UpdateProduct = m(e.UpdateProduct)
	e.GetProduct = m(e.GetProduct)
}

// NewGetAllProductsByCategoryEndpoint returns an endpoint function that calls
// the method "getAllProductsByCategory" of service "products".
func NewGetAllProductsByCategoryEndpoint(s Service, authOAuth2Fn security.AuthOAuth2Func, authJWTFn security.AuthJWTFunc) goa.Endpoint {
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
		if err == nil {
			sc := security.JWTScheme{
				Name:           "jwt",
				Scopes:         []string{"api:read", "api:write"},
				RequiredScopes: []string{},
			}
			var token string
			if p.JWTToken != nil {
				token = *p.JWTToken
			}
			ctx, err = authJWTFn(ctx, token, &sc)
		}
		if err != nil {
			return nil, err
		}
		return s.GetAllProductsByCategory(ctx, p)
	}
}

// NewDeleteProductEndpoint returns an endpoint function that calls the method
// "deleteProduct" of service "products".
func NewDeleteProductEndpoint(s Service, authOAuth2Fn security.AuthOAuth2Func, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*DeleteProductPayload)
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
		if err == nil {
			sc := security.JWTScheme{
				Name:           "jwt",
				Scopes:         []string{"api:read", "api:write"},
				RequiredScopes: []string{},
			}
			var token string
			if p.JWTToken != nil {
				token = *p.JWTToken
			}
			ctx, err = authJWTFn(ctx, token, &sc)
		}
		if err != nil {
			return nil, err
		}
		return s.DeleteProduct(ctx, p)
	}
}

// NewCreateProductEndpoint returns an endpoint function that calls the method
// "createProduct" of service "products".
func NewCreateProductEndpoint(s Service, authOAuth2Fn security.AuthOAuth2Func, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*CreateProductPayload)
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
		if err == nil {
			sc := security.JWTScheme{
				Name:           "jwt",
				Scopes:         []string{"api:read", "api:write"},
				RequiredScopes: []string{},
			}
			var token string
			if p.JWTToken != nil {
				token = *p.JWTToken
			}
			ctx, err = authJWTFn(ctx, token, &sc)
		}
		if err != nil {
			return nil, err
		}
		return s.CreateProduct(ctx, p)
	}
}

// NewUpdateProductEndpoint returns an endpoint function that calls the method
// "updateProduct" of service "products".
func NewUpdateProductEndpoint(s Service, authOAuth2Fn security.AuthOAuth2Func, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*UpdateProductPayload)
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
		if err == nil {
			sc := security.JWTScheme{
				Name:           "jwt",
				Scopes:         []string{"api:read", "api:write"},
				RequiredScopes: []string{},
			}
			var token string
			if p.JWTToken != nil {
				token = *p.JWTToken
			}
			ctx, err = authJWTFn(ctx, token, &sc)
		}
		if err != nil {
			return nil, err
		}
		return s.UpdateProduct(ctx, p)
	}
}

// NewGetProductEndpoint returns an endpoint function that calls the method
// "getProduct" of service "products".
func NewGetProductEndpoint(s Service, authOAuth2Fn security.AuthOAuth2Func, authJWTFn security.AuthJWTFunc) goa.Endpoint {
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
		if err == nil {
			sc := security.JWTScheme{
				Name:           "jwt",
				Scopes:         []string{"api:read", "api:write"},
				RequiredScopes: []string{},
			}
			var token string
			if p.JWTToken != nil {
				token = *p.JWTToken
			}
			ctx, err = authJWTFn(ctx, token, &sc)
		}
		if err != nil {
			return nil, err
		}
		return s.GetProduct(ctx, p)
	}
}
