// Code generated by goa v3.5.2, DO NOT EDIT.
//
// auth HTTP client types
//
// Command:
// $ goa gen api_crud/design

package client

import (
	auth "api_crud/gen/auth"

	goa "goa.design/goa/v3/pkg"
)

// EmailExistRequestBody is the type of the "auth" service "email-exist"
// endpoint HTTP request body.
type EmailExistRequestBody struct {
	Email string `form:"email" json:"email" xml:"email"`
}

// SendConfirmationRequestBody is the type of the "auth" service
// "send-confirmation" endpoint HTTP request body.
type SendConfirmationRequestBody struct {
	Email string `form:"email" json:"email" xml:"email"`
}

// ResetPasswordRequestBody is the type of the "auth" service "reset-password"
// endpoint HTTP request body.
type ResetPasswordRequestBody struct {
	Email string `form:"email" json:"email" xml:"email"`
	Code  string `form:"code" json:"code" xml:"code"`
	// Minimum 8 charactères / Chiffre Obligatoire
	Password string `form:"password" json:"password" xml:"password"`
	// Minimum 8 charactères / Chiffre Obligatoire
	ConfirmPassword string `form:"confirm_password" json:"confirm_password" xml:"confirm_password"`
}

// EmailExistResponseBody is the type of the "auth" service "email-exist"
// endpoint HTTP response body.
type EmailExistResponseBody struct {
	Success *bool `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
	Exist   *bool `form:"exist,omitempty" json:"exist,omitempty" xml:"exist,omitempty"`
}

// SendConfirmationResponseBody is the type of the "auth" service
// "send-confirmation" endpoint HTTP response body.
type SendConfirmationResponseBody struct {
	Success *bool `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
	Exist   *bool `form:"exist,omitempty" json:"exist,omitempty" xml:"exist,omitempty"`
}

// ResetPasswordResponseBody is the type of the "auth" service "reset-password"
// endpoint HTTP response body.
type ResetPasswordResponseBody struct {
	Success *bool `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
}

// EmailExistUnknownErrorResponseBody is the type of the "auth" service
// "email-exist" endpoint HTTP response body for the "unknown_error" error.
type EmailExistUnknownErrorResponseBody struct {
	Err       *string `form:"err,omitempty" json:"err,omitempty" xml:"err,omitempty"`
	ErrorCode *string `form:"error_code,omitempty" json:"error_code,omitempty" xml:"error_code,omitempty"`
	Success   *bool   `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
}

// SendConfirmationUnknownErrorResponseBody is the type of the "auth" service
// "send-confirmation" endpoint HTTP response body for the "unknown_error"
// error.
type SendConfirmationUnknownErrorResponseBody struct {
	Err       *string `form:"err,omitempty" json:"err,omitempty" xml:"err,omitempty"`
	ErrorCode *string `form:"error_code,omitempty" json:"error_code,omitempty" xml:"error_code,omitempty"`
	Success   *bool   `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
}

// ResetPasswordUnknownErrorResponseBody is the type of the "auth" service
// "reset-password" endpoint HTTP response body for the "unknown_error" error.
type ResetPasswordUnknownErrorResponseBody struct {
	Err       *string `form:"err,omitempty" json:"err,omitempty" xml:"err,omitempty"`
	ErrorCode *string `form:"error_code,omitempty" json:"error_code,omitempty" xml:"error_code,omitempty"`
	Success   *bool   `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
}

// NewEmailExistRequestBody builds the HTTP request body from the payload of
// the "email-exist" endpoint of the "auth" service.
func NewEmailExistRequestBody(p *auth.EmailExistPayload) *EmailExistRequestBody {
	body := &EmailExistRequestBody{
		Email: p.Email,
	}
	return body
}

// NewSendConfirmationRequestBody builds the HTTP request body from the payload
// of the "send-confirmation" endpoint of the "auth" service.
func NewSendConfirmationRequestBody(p *auth.SendConfirmationPayload) *SendConfirmationRequestBody {
	body := &SendConfirmationRequestBody{
		Email: p.Email,
	}
	return body
}

// NewResetPasswordRequestBody builds the HTTP request body from the payload of
// the "reset-password" endpoint of the "auth" service.
func NewResetPasswordRequestBody(p *auth.ResetPasswordPayload) *ResetPasswordRequestBody {
	body := &ResetPasswordRequestBody{
		Email:           p.Email,
		Code:            p.Code,
		Password:        p.Password,
		ConfirmPassword: p.ConfirmPassword,
	}
	return body
}

// NewEmailExistResultOK builds a "auth" service "email-exist" endpoint result
// from a HTTP "OK" response.
func NewEmailExistResultOK(body *EmailExistResponseBody) *auth.EmailExistResult {
	v := &auth.EmailExistResult{
		Success: *body.Success,
		Exist:   *body.Exist,
	}

	return v
}

// NewEmailExistUnknownError builds a auth service email-exist endpoint
// unknown_error error.
func NewEmailExistUnknownError(body *EmailExistUnknownErrorResponseBody) *auth.UnknownError {
	v := &auth.UnknownError{
		Err:       *body.Err,
		ErrorCode: *body.ErrorCode,
		Success:   *body.Success,
	}

	return v
}

// NewSendConfirmationResultOK builds a "auth" service "send-confirmation"
// endpoint result from a HTTP "OK" response.
func NewSendConfirmationResultOK(body *SendConfirmationResponseBody) *auth.SendConfirmationResult {
	v := &auth.SendConfirmationResult{
		Success: *body.Success,
		Exist:   *body.Exist,
	}

	return v
}

// NewSendConfirmationUnknownError builds a auth service send-confirmation
// endpoint unknown_error error.
func NewSendConfirmationUnknownError(body *SendConfirmationUnknownErrorResponseBody) *auth.UnknownError {
	v := &auth.UnknownError{
		Err:       *body.Err,
		ErrorCode: *body.ErrorCode,
		Success:   *body.Success,
	}

	return v
}

// NewResetPasswordResultOK builds a "auth" service "reset-password" endpoint
// result from a HTTP "OK" response.
func NewResetPasswordResultOK(body *ResetPasswordResponseBody) *auth.ResetPasswordResult {
	v := &auth.ResetPasswordResult{
		Success: *body.Success,
	}

	return v
}

// NewResetPasswordUnknownError builds a auth service reset-password endpoint
// unknown_error error.
func NewResetPasswordUnknownError(body *ResetPasswordUnknownErrorResponseBody) *auth.UnknownError {
	v := &auth.UnknownError{
		Err:       *body.Err,
		ErrorCode: *body.ErrorCode,
		Success:   *body.Success,
	}

	return v
}

// ValidateEmailExistResponseBody runs the validations defined on
// Email-ExistResponseBody
func ValidateEmailExistResponseBody(body *EmailExistResponseBody) (err error) {
	if body.Exist == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("exist", "body"))
	}
	if body.Success == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("success", "body"))
	}
	return
}

// ValidateSendConfirmationResponseBody runs the validations defined on
// Send-ConfirmationResponseBody
func ValidateSendConfirmationResponseBody(body *SendConfirmationResponseBody) (err error) {
	if body.Exist == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("exist", "body"))
	}
	if body.Success == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("success", "body"))
	}
	return
}

// ValidateResetPasswordResponseBody runs the validations defined on
// Reset-PasswordResponseBody
func ValidateResetPasswordResponseBody(body *ResetPasswordResponseBody) (err error) {
	if body.Success == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("success", "body"))
	}
	return
}

// ValidateEmailExistUnknownErrorResponseBody runs the validations defined on
// email-exist_unknown_error_response_body
func ValidateEmailExistUnknownErrorResponseBody(body *EmailExistUnknownErrorResponseBody) (err error) {
	if body.Err == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("err", "body"))
	}
	if body.Success == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("success", "body"))
	}
	if body.ErrorCode == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("error_code", "body"))
	}
	return
}

// ValidateSendConfirmationUnknownErrorResponseBody runs the validations
// defined on send-confirmation_unknown_error_response_body
func ValidateSendConfirmationUnknownErrorResponseBody(body *SendConfirmationUnknownErrorResponseBody) (err error) {
	if body.Err == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("err", "body"))
	}
	if body.Success == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("success", "body"))
	}
	if body.ErrorCode == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("error_code", "body"))
	}
	return
}

// ValidateResetPasswordUnknownErrorResponseBody runs the validations defined
// on reset-password_unknown_error_response_body
func ValidateResetPasswordUnknownErrorResponseBody(body *ResetPasswordUnknownErrorResponseBody) (err error) {
	if body.Err == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("err", "body"))
	}
	if body.Success == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("success", "body"))
	}
	if body.ErrorCode == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("error_code", "body"))
	}
	return
}
