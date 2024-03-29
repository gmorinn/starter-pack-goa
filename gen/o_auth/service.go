// Code generated by goa v3.10.2, DO NOT EDIT.
//
// oAuth service
//
// Command:
// $ goa gen starter-pack-goa/design

package oauth

import (
	"context"
)

// Oauth to authentificate
type Service interface {
	// oAuth
	OAuth(context.Context, *OauthPayload) (res *OAuthResponse, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "oAuth"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"oAuth"}

// OAuthResponse is the result type of the oAuth service oAuth method.
type OAuthResponse struct {
	AccessToken string
	TokenType   string
	ExpiresIn   int64
	Success     bool
}

// OauthPayload is the payload type of the oAuth service oAuth method.
type OauthPayload struct {
	ClientID     string
	ClientSecret string
	GrantType    string
}

// Token scopes are invalid
type InvalidScopes string

// Error when a request is send before asking for oAuth
type OauthError string

// Identifiers are invalid
type Unauthorized string

type UnknownError struct {
	Err       string
	ErrorCode string
	Success   bool
}

// Error returns an error description.
func (e InvalidScopes) Error() string {
	return "Token scopes are invalid"
}

// ErrorName returns "invalid_scopes".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e InvalidScopes) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "invalid_scopes".
func (e InvalidScopes) GoaErrorName() string {
	return "invalid_scopes"
}

// Error returns an error description.
func (e OauthError) Error() string {
	return "Error when a request is send before asking for oAuth"
}

// ErrorName returns "oauth_error".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e OauthError) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "oauth_error".
func (e OauthError) GoaErrorName() string {
	return "oauth_error"
}

// Error returns an error description.
func (e Unauthorized) Error() string {
	return "Identifiers are invalid"
}

// ErrorName returns "unauthorized".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e Unauthorized) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "unauthorized".
func (e Unauthorized) GoaErrorName() string {
	return "unauthorized"
}

// Error returns an error description.
func (e *UnknownError) Error() string {
	return ""
}

// ErrorName returns "unknownError".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e *UnknownError) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "unknownError".
func (e *UnknownError) GoaErrorName() string {
	return "unknown_error"
}
