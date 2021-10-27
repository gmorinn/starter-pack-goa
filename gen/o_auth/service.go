// Code generated by goa v3.5.2, DO NOT EDIT.
//
// oAuth service
//
// Command:
// $ goa gen api_crud/design

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

// OauthPayload is the payload type of the oAuth service oAuth method.
type OauthPayload struct {
	ClientID     string
	ClientSecret string
	GrantType    string
}

// OAuthResponse is the result type of the oAuth service oAuth method.
type OAuthResponse struct {
	AccessToken string
	TokenType   string
	ExpiresIn   int64
	Success     bool
}

type UnknownError struct {
	Err       string
	ErrorCode string
	Success   bool
}

// Token scopes are invalid
type InvalidScopes string

// Identifiers are invalid
type Unauthorized string

// Error returns an error description.
func (e *UnknownError) Error() string {
	return ""
}

// ErrorName returns "unknownError".
func (e *UnknownError) ErrorName() string {
	return "unknown_error"
}

// Error returns an error description.
func (e InvalidScopes) Error() string {
	return "Token scopes are invalid"
}

// ErrorName returns "invalid_scopes".
func (e InvalidScopes) ErrorName() string {
	return "invalid_scopes"
}

// Error returns an error description.
func (e Unauthorized) Error() string {
	return "Identifiers are invalid"
}

// ErrorName returns "unauthorized".
func (e Unauthorized) ErrorName() string {
	return "unauthorized"
}
