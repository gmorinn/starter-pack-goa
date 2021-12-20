// Code generated by goa v3.5.2, DO NOT EDIT.
//
// jwtToken endpoints
//
// Command:
// $ goa gen api_crud/design

package jwttoken

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "jwtToken" service endpoints.
type Endpoints struct {
	Signup        goa.Endpoint
	Signin        goa.Endpoint
	Refresh       goa.Endpoint
	AuthProviders goa.Endpoint
}

// NewEndpoints wraps the methods of the "jwtToken" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		Signup:        NewSignupEndpoint(s, a.OAuth2Auth),
		Signin:        NewSigninEndpoint(s, a.OAuth2Auth),
		Refresh:       NewRefreshEndpoint(s, a.OAuth2Auth),
		AuthProviders: NewAuthProvidersEndpoint(s, a.OAuth2Auth),
	}
}

// Use applies the given middleware to all the "jwtToken" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Signup = m(e.Signup)
	e.Signin = m(e.Signin)
	e.Refresh = m(e.Refresh)
	e.AuthProviders = m(e.AuthProviders)
}

// NewSignupEndpoint returns an endpoint function that calls the method
// "signup" of service "jwtToken".
func NewSignupEndpoint(s Service, authOAuth2Fn security.AuthOAuth2Func) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*SignupPayload)
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
		return s.Signup(ctx, p)
	}
}

// NewSigninEndpoint returns an endpoint function that calls the method
// "signin" of service "jwtToken".
func NewSigninEndpoint(s Service, authOAuth2Fn security.AuthOAuth2Func) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*SigninPayload)
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
		return s.Signin(ctx, p)
	}
}

// NewRefreshEndpoint returns an endpoint function that calls the method
// "refresh" of service "jwtToken".
func NewRefreshEndpoint(s Service, authOAuth2Fn security.AuthOAuth2Func) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*RefreshPayload)
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
		return s.Refresh(ctx, p)
	}
}

// NewAuthProvidersEndpoint returns an endpoint function that calls the method
// "auth-providers" of service "jwtToken".
func NewAuthProvidersEndpoint(s Service, authOAuth2Fn security.AuthOAuth2Func) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*AuthProvidersPayload)
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
		return s.AuthProviders(ctx, p)
	}
}
