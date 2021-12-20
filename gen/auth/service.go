// Code generated by goa v3.5.2, DO NOT EDIT.
//
// auth service
//
// Command:
// $ goa gen api_crud/design

package auth

import (
	"context"

	"goa.design/goa/v3/security"
)

// Forget password / reset password / send Email Code
type Service interface {
	// Check if email exist in database
	EmailExist(context.Context, *EmailExistPayload) (res *EmailExistResult, err error)
	// Check if email exist in database and send code by email to reset password
	SendConfirmation(context.Context, *SendConfirmationPayload) (res *SendConfirmationResult, err error)
	// Reset forget password of the user with the correct confirm code
	ResetPassword(context.Context, *ResetPasswordPayload) (res *ResetPasswordResult, err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// OAuth2Auth implements the authorization logic for the OAuth2 security scheme.
	OAuth2Auth(ctx context.Context, token string, schema *security.OAuth2Scheme) (context.Context, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "auth"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [3]string{"email-exist", "send-confirmation", "reset-password"}

// EmailExistPayload is the payload type of the auth service email-exist method.
type EmailExistPayload struct {
	Email string
	// Use to generate Oauth with /authorization
	Oauth *string
}

// EmailExistResult is the result type of the auth service email-exist method.
type EmailExistResult struct {
	Success bool
	Exist   bool
}

// SendConfirmationPayload is the payload type of the auth service
// send-confirmation method.
type SendConfirmationPayload struct {
	Email string
	// Use to generate Oauth with /authorization
	Oauth *string
}

// SendConfirmationResult is the result type of the auth service
// send-confirmation method.
type SendConfirmationResult struct {
	Success bool
	Exist   bool
}

// ResetPasswordPayload is the payload type of the auth service reset-password
// method.
type ResetPasswordPayload struct {
	Email string
	Code  string
	// Minimum 8 charactères / Chiffre Obligatoire
	Password string
	// Minimum 8 charactères / Chiffre Obligatoire
	ConfirmPassword string
	// Use to generate Oauth with /authorization
	Oauth *string
}

// ResetPasswordResult is the result type of the auth service reset-password
// method.
type ResetPasswordResult struct {
	Success bool
}

type UnknownError struct {
	Err       string
	ErrorCode string
	Success   bool
}

// Error returns an error description.
func (e *UnknownError) Error() string {
	return ""
}

// ErrorName returns "unknownError".
func (e *UnknownError) ErrorName() string {
	return "unknown_error"
}
