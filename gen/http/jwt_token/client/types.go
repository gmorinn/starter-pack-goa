// Code generated by goa v3.5.2, DO NOT EDIT.
//
// jwtToken HTTP client types
//
// Command:
// $ goa gen api_crud/design

package client

import (
	jwttoken "api_crud/gen/jwt_token"

	goa "goa.design/goa/v3/pkg"
)

// SignupRequestBody is the type of the "jwtToken" service "signup" endpoint
// HTTP request body.
type SignupRequestBody struct {
	Firstname string `form:"firstname" json:"firstname" xml:"firstname"`
	Lastname  string `form:"lastname" json:"lastname" xml:"lastname"`
	// Minimum 8 charactères / Chiffre Obligatoire
	Password string `form:"password" json:"password" xml:"password"`
	// Minimum 8 charactères / Chiffre Obligatoire
	ConfirmPassword string `form:"confirm_password" json:"confirm_password" xml:"confirm_password"`
	Email           string `form:"email" json:"email" xml:"email"`
	Birthday        string `form:"birthday" json:"birthday" xml:"birthday"`
	Phone           string `form:"phone" json:"phone" xml:"phone"`
}

// SigninRequestBody is the type of the "jwtToken" service "signin" endpoint
// HTTP request body.
type SigninRequestBody struct {
	Email string `form:"email" json:"email" xml:"email"`
	// Minimum 8 charactères / Chiffre Obligatoire
	Password string `form:"password" json:"password" xml:"password"`
}

// RefreshRequestBody is the type of the "jwtToken" service "refresh" endpoint
// HTTP request body.
type RefreshRequestBody struct {
	RefreshToken string `form:"refresh_token" json:"refresh_token" xml:"refresh_token"`
}

// EmailExistRequestBody is the type of the "jwtToken" service "email-exist"
// endpoint HTTP request body.
type EmailExistRequestBody struct {
	Email string `form:"email" json:"email" xml:"email"`
}

// AuthProvidersRequestBody is the type of the "jwtToken" service
// "auth-providers" endpoint HTTP request body.
type AuthProvidersRequestBody struct {
	Firstname        string `form:"firstname" json:"firstname" xml:"firstname"`
	Lastname         string `form:"lastname" json:"lastname" xml:"lastname"`
	Email            string `form:"email" json:"email" xml:"email"`
	FirebaseIDToken  string `form:"firebase_id_token" json:"firebase_id_token" xml:"firebase_id_token"`
	FirebaseUID      string `form:"firebase_uid" json:"firebase_uid" xml:"firebase_uid"`
	FirebaseProvider string `form:"firebase_provider" json:"firebase_provider" xml:"firebase_provider"`
}

// SignupResponseBody is the type of the "jwtToken" service "signup" endpoint
// HTTP response body.
type SignupResponseBody struct {
	AccessToken  *string `form:"access_token,omitempty" json:"access_token,omitempty" xml:"access_token,omitempty"`
	RefreshToken *string `form:"refresh_token,omitempty" json:"refresh_token,omitempty" xml:"refresh_token,omitempty"`
	Success      *bool   `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
}

// SigninResponseBody is the type of the "jwtToken" service "signin" endpoint
// HTTP response body.
type SigninResponseBody struct {
	AccessToken  *string `form:"access_token,omitempty" json:"access_token,omitempty" xml:"access_token,omitempty"`
	RefreshToken *string `form:"refresh_token,omitempty" json:"refresh_token,omitempty" xml:"refresh_token,omitempty"`
	Success      *bool   `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
}

// RefreshResponseBody is the type of the "jwtToken" service "refresh" endpoint
// HTTP response body.
type RefreshResponseBody struct {
	AccessToken  *string `form:"access_token,omitempty" json:"access_token,omitempty" xml:"access_token,omitempty"`
	RefreshToken *string `form:"refresh_token,omitempty" json:"refresh_token,omitempty" xml:"refresh_token,omitempty"`
	Success      *bool   `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
}

// EmailExistResponseBody is the type of the "jwtToken" service "email-exist"
// endpoint HTTP response body.
type EmailExistResponseBody struct {
	Success *bool `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
	Exist   *bool `form:"exist,omitempty" json:"exist,omitempty" xml:"exist,omitempty"`
}

// AuthProvidersCreatedResponseBody is the type of the "jwtToken" service
// "auth-providers" endpoint HTTP response body.
type AuthProvidersCreatedResponseBody struct {
	AccessToken  *string `form:"access_token,omitempty" json:"access_token,omitempty" xml:"access_token,omitempty"`
	RefreshToken *string `form:"refresh_token,omitempty" json:"refresh_token,omitempty" xml:"refresh_token,omitempty"`
	Success      *bool   `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
}

// SignupEmailAlreadyExistResponseBody is the type of the "jwtToken" service
// "signup" endpoint HTTP response body for the "email_already_exist" error.
type SignupEmailAlreadyExistResponseBody struct {
	Err     *string `form:"err,omitempty" json:"err,omitempty" xml:"err,omitempty"`
	Success *bool   `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
}

// SignupUnknownErrorResponseBody is the type of the "jwtToken" service
// "signup" endpoint HTTP response body for the "unknown_error" error.
type SignupUnknownErrorResponseBody struct {
	Err       *string `form:"err,omitempty" json:"err,omitempty" xml:"err,omitempty"`
	ErrorCode *string `form:"error_code,omitempty" json:"error_code,omitempty" xml:"error_code,omitempty"`
	Success   *bool   `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
}

// SignupInvalidScopesResponseBody is the type of the "jwtToken" service
// "signup" endpoint HTTP response body for the "invalid_scopes" error.
type SignupInvalidScopesResponseBody string

// SignupUnauthorizedResponseBody is the type of the "jwtToken" service
// "signup" endpoint HTTP response body for the "unauthorized" error.
type SignupUnauthorizedResponseBody string

// SigninEmailAlreadyExistResponseBody is the type of the "jwtToken" service
// "signin" endpoint HTTP response body for the "email_already_exist" error.
type SigninEmailAlreadyExistResponseBody struct {
	Err     *string `form:"err,omitempty" json:"err,omitempty" xml:"err,omitempty"`
	Success *bool   `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
}

// SigninUnknownErrorResponseBody is the type of the "jwtToken" service
// "signin" endpoint HTTP response body for the "unknown_error" error.
type SigninUnknownErrorResponseBody struct {
	Err       *string `form:"err,omitempty" json:"err,omitempty" xml:"err,omitempty"`
	ErrorCode *string `form:"error_code,omitempty" json:"error_code,omitempty" xml:"error_code,omitempty"`
	Success   *bool   `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
}

// SigninInvalidScopesResponseBody is the type of the "jwtToken" service
// "signin" endpoint HTTP response body for the "invalid_scopes" error.
type SigninInvalidScopesResponseBody string

// SigninUnauthorizedResponseBody is the type of the "jwtToken" service
// "signin" endpoint HTTP response body for the "unauthorized" error.
type SigninUnauthorizedResponseBody string

// RefreshEmailAlreadyExistResponseBody is the type of the "jwtToken" service
// "refresh" endpoint HTTP response body for the "email_already_exist" error.
type RefreshEmailAlreadyExistResponseBody struct {
	Err     *string `form:"err,omitempty" json:"err,omitempty" xml:"err,omitempty"`
	Success *bool   `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
}

// RefreshUnknownErrorResponseBody is the type of the "jwtToken" service
// "refresh" endpoint HTTP response body for the "unknown_error" error.
type RefreshUnknownErrorResponseBody struct {
	Err       *string `form:"err,omitempty" json:"err,omitempty" xml:"err,omitempty"`
	ErrorCode *string `form:"error_code,omitempty" json:"error_code,omitempty" xml:"error_code,omitempty"`
	Success   *bool   `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
}

// RefreshInvalidScopesResponseBody is the type of the "jwtToken" service
// "refresh" endpoint HTTP response body for the "invalid_scopes" error.
type RefreshInvalidScopesResponseBody string

// RefreshUnauthorizedResponseBody is the type of the "jwtToken" service
// "refresh" endpoint HTTP response body for the "unauthorized" error.
type RefreshUnauthorizedResponseBody string

// EmailExistEmailAlreadyExistResponseBody is the type of the "jwtToken"
// service "email-exist" endpoint HTTP response body for the
// "email_already_exist" error.
type EmailExistEmailAlreadyExistResponseBody struct {
	Err     *string `form:"err,omitempty" json:"err,omitempty" xml:"err,omitempty"`
	Success *bool   `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
}

// EmailExistUnknownErrorResponseBody is the type of the "jwtToken" service
// "email-exist" endpoint HTTP response body for the "unknown_error" error.
type EmailExistUnknownErrorResponseBody struct {
	Err       *string `form:"err,omitempty" json:"err,omitempty" xml:"err,omitempty"`
	ErrorCode *string `form:"error_code,omitempty" json:"error_code,omitempty" xml:"error_code,omitempty"`
	Success   *bool   `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
}

// EmailExistInvalidScopesResponseBody is the type of the "jwtToken" service
// "email-exist" endpoint HTTP response body for the "invalid_scopes" error.
type EmailExistInvalidScopesResponseBody string

// EmailExistUnauthorizedResponseBody is the type of the "jwtToken" service
// "email-exist" endpoint HTTP response body for the "unauthorized" error.
type EmailExistUnauthorizedResponseBody string

// AuthProvidersEmailAlreadyExistResponseBody is the type of the "jwtToken"
// service "auth-providers" endpoint HTTP response body for the
// "email_already_exist" error.
type AuthProvidersEmailAlreadyExistResponseBody struct {
	Err     *string `form:"err,omitempty" json:"err,omitempty" xml:"err,omitempty"`
	Success *bool   `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
}

// AuthProvidersUnknownErrorResponseBody is the type of the "jwtToken" service
// "auth-providers" endpoint HTTP response body for the "unknown_error" error.
type AuthProvidersUnknownErrorResponseBody struct {
	Err       *string `form:"err,omitempty" json:"err,omitempty" xml:"err,omitempty"`
	ErrorCode *string `form:"error_code,omitempty" json:"error_code,omitempty" xml:"error_code,omitempty"`
	Success   *bool   `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
}

// AuthProvidersInvalidScopesResponseBody is the type of the "jwtToken" service
// "auth-providers" endpoint HTTP response body for the "invalid_scopes" error.
type AuthProvidersInvalidScopesResponseBody string

// AuthProvidersUnauthorizedResponseBody is the type of the "jwtToken" service
// "auth-providers" endpoint HTTP response body for the "unauthorized" error.
type AuthProvidersUnauthorizedResponseBody string

// AuthProvidersBadRequestResponseBody is used to define fields on response
// body types.
type AuthProvidersBadRequestResponseBody struct {
	AccessToken  *string `form:"access_token,omitempty" json:"access_token,omitempty" xml:"access_token,omitempty"`
	RefreshToken *string `form:"refresh_token,omitempty" json:"refresh_token,omitempty" xml:"refresh_token,omitempty"`
	Success      *bool   `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
}

// NewSignupRequestBody builds the HTTP request body from the payload of the
// "signup" endpoint of the "jwtToken" service.
func NewSignupRequestBody(p *jwttoken.SignupPayload) *SignupRequestBody {
	body := &SignupRequestBody{
		Firstname:       p.Firstname,
		Lastname:        p.Lastname,
		Password:        p.Password,
		ConfirmPassword: p.ConfirmPassword,
		Email:           p.Email,
		Birthday:        p.Birthday,
		Phone:           p.Phone,
	}
	{
		var zero string
		if body.Birthday == zero {
			body.Birthday = ""
		}
	}
	{
		var zero string
		if body.Phone == zero {
			body.Phone = ""
		}
	}
	return body
}

// NewSigninRequestBody builds the HTTP request body from the payload of the
// "signin" endpoint of the "jwtToken" service.
func NewSigninRequestBody(p *jwttoken.SigninPayload) *SigninRequestBody {
	body := &SigninRequestBody{
		Email:    p.Email,
		Password: p.Password,
	}
	return body
}

// NewRefreshRequestBody builds the HTTP request body from the payload of the
// "refresh" endpoint of the "jwtToken" service.
func NewRefreshRequestBody(p *jwttoken.RefreshPayload) *RefreshRequestBody {
	body := &RefreshRequestBody{
		RefreshToken: p.RefreshToken,
	}
	return body
}

// NewEmailExistRequestBody builds the HTTP request body from the payload of
// the "email-exist" endpoint of the "jwtToken" service.
func NewEmailExistRequestBody(p *jwttoken.EmailExistPayload) *EmailExistRequestBody {
	body := &EmailExistRequestBody{
		Email: p.Email,
	}
	return body
}

// NewAuthProvidersRequestBody builds the HTTP request body from the payload of
// the "auth-providers" endpoint of the "jwtToken" service.
func NewAuthProvidersRequestBody(p *jwttoken.AuthProvidersPayload) *AuthProvidersRequestBody {
	body := &AuthProvidersRequestBody{
		Firstname:        p.Firstname,
		Lastname:         p.Lastname,
		Email:            p.Email,
		FirebaseIDToken:  p.FirebaseIDToken,
		FirebaseUID:      p.FirebaseUID,
		FirebaseProvider: p.FirebaseProvider,
	}
	return body
}

// NewSignupSignOK builds a "jwtToken" service "signup" endpoint result from a
// HTTP "OK" response.
func NewSignupSignOK(body *SignupResponseBody) *jwttoken.Sign {
	v := &jwttoken.Sign{
		AccessToken:  *body.AccessToken,
		RefreshToken: *body.RefreshToken,
		Success:      *body.Success,
	}

	return v
}

// NewSignupEmailAlreadyExist builds a jwtToken service signup endpoint
// email_already_exist error.
func NewSignupEmailAlreadyExist(body *SignupEmailAlreadyExistResponseBody) *jwttoken.EmailAlreadyExist {
	v := &jwttoken.EmailAlreadyExist{
		Err:     *body.Err,
		Success: *body.Success,
	}

	return v
}

// NewSignupUnknownError builds a jwtToken service signup endpoint
// unknown_error error.
func NewSignupUnknownError(body *SignupUnknownErrorResponseBody) *jwttoken.UnknownError {
	v := &jwttoken.UnknownError{
		Err:       *body.Err,
		ErrorCode: *body.ErrorCode,
		Success:   *body.Success,
	}

	return v
}

// NewSignupInvalidScopes builds a jwtToken service signup endpoint
// invalid_scopes error.
func NewSignupInvalidScopes(body SignupInvalidScopesResponseBody) jwttoken.InvalidScopes {
	v := jwttoken.InvalidScopes(body)

	return v
}

// NewSignupUnauthorized builds a jwtToken service signup endpoint unauthorized
// error.
func NewSignupUnauthorized(body SignupUnauthorizedResponseBody) jwttoken.Unauthorized {
	v := jwttoken.Unauthorized(body)

	return v
}

// NewSigninSignOK builds a "jwtToken" service "signin" endpoint result from a
// HTTP "OK" response.
func NewSigninSignOK(body *SigninResponseBody) *jwttoken.Sign {
	v := &jwttoken.Sign{
		AccessToken:  *body.AccessToken,
		RefreshToken: *body.RefreshToken,
		Success:      *body.Success,
	}

	return v
}

// NewSigninEmailAlreadyExist builds a jwtToken service signin endpoint
// email_already_exist error.
func NewSigninEmailAlreadyExist(body *SigninEmailAlreadyExistResponseBody) *jwttoken.EmailAlreadyExist {
	v := &jwttoken.EmailAlreadyExist{
		Err:     *body.Err,
		Success: *body.Success,
	}

	return v
}

// NewSigninUnknownError builds a jwtToken service signin endpoint
// unknown_error error.
func NewSigninUnknownError(body *SigninUnknownErrorResponseBody) *jwttoken.UnknownError {
	v := &jwttoken.UnknownError{
		Err:       *body.Err,
		ErrorCode: *body.ErrorCode,
		Success:   *body.Success,
	}

	return v
}

// NewSigninInvalidScopes builds a jwtToken service signin endpoint
// invalid_scopes error.
func NewSigninInvalidScopes(body SigninInvalidScopesResponseBody) jwttoken.InvalidScopes {
	v := jwttoken.InvalidScopes(body)

	return v
}

// NewSigninUnauthorized builds a jwtToken service signin endpoint unauthorized
// error.
func NewSigninUnauthorized(body SigninUnauthorizedResponseBody) jwttoken.Unauthorized {
	v := jwttoken.Unauthorized(body)

	return v
}

// NewRefreshSignOK builds a "jwtToken" service "refresh" endpoint result from
// a HTTP "OK" response.
func NewRefreshSignOK(body *RefreshResponseBody) *jwttoken.Sign {
	v := &jwttoken.Sign{
		AccessToken:  *body.AccessToken,
		RefreshToken: *body.RefreshToken,
		Success:      *body.Success,
	}

	return v
}

// NewRefreshEmailAlreadyExist builds a jwtToken service refresh endpoint
// email_already_exist error.
func NewRefreshEmailAlreadyExist(body *RefreshEmailAlreadyExistResponseBody) *jwttoken.EmailAlreadyExist {
	v := &jwttoken.EmailAlreadyExist{
		Err:     *body.Err,
		Success: *body.Success,
	}

	return v
}

// NewRefreshUnknownError builds a jwtToken service refresh endpoint
// unknown_error error.
func NewRefreshUnknownError(body *RefreshUnknownErrorResponseBody) *jwttoken.UnknownError {
	v := &jwttoken.UnknownError{
		Err:       *body.Err,
		ErrorCode: *body.ErrorCode,
		Success:   *body.Success,
	}

	return v
}

// NewRefreshInvalidScopes builds a jwtToken service refresh endpoint
// invalid_scopes error.
func NewRefreshInvalidScopes(body RefreshInvalidScopesResponseBody) jwttoken.InvalidScopes {
	v := jwttoken.InvalidScopes(body)

	return v
}

// NewRefreshUnauthorized builds a jwtToken service refresh endpoint
// unauthorized error.
func NewRefreshUnauthorized(body RefreshUnauthorizedResponseBody) jwttoken.Unauthorized {
	v := jwttoken.Unauthorized(body)

	return v
}

// NewEmailExistResultOK builds a "jwtToken" service "email-exist" endpoint
// result from a HTTP "OK" response.
func NewEmailExistResultOK(body *EmailExistResponseBody) *jwttoken.EmailExistResult {
	v := &jwttoken.EmailExistResult{
		Success: *body.Success,
		Exist:   *body.Exist,
	}

	return v
}

// NewEmailExistEmailAlreadyExist builds a jwtToken service email-exist
// endpoint email_already_exist error.
func NewEmailExistEmailAlreadyExist(body *EmailExistEmailAlreadyExistResponseBody) *jwttoken.EmailAlreadyExist {
	v := &jwttoken.EmailAlreadyExist{
		Err:     *body.Err,
		Success: *body.Success,
	}

	return v
}

// NewEmailExistUnknownError builds a jwtToken service email-exist endpoint
// unknown_error error.
func NewEmailExistUnknownError(body *EmailExistUnknownErrorResponseBody) *jwttoken.UnknownError {
	v := &jwttoken.UnknownError{
		Err:       *body.Err,
		ErrorCode: *body.ErrorCode,
		Success:   *body.Success,
	}

	return v
}

// NewEmailExistInvalidScopes builds a jwtToken service email-exist endpoint
// invalid_scopes error.
func NewEmailExistInvalidScopes(body EmailExistInvalidScopesResponseBody) jwttoken.InvalidScopes {
	v := jwttoken.InvalidScopes(body)

	return v
}

// NewEmailExistUnauthorized builds a jwtToken service email-exist endpoint
// unauthorized error.
func NewEmailExistUnauthorized(body EmailExistUnauthorizedResponseBody) jwttoken.Unauthorized {
	v := jwttoken.Unauthorized(body)

	return v
}

// NewAuthProvidersSignCreated builds a "jwtToken" service "auth-providers"
// endpoint result from a HTTP "Created" response.
func NewAuthProvidersSignCreated(body *AuthProvidersCreatedResponseBody) *jwttoken.Sign {
	v := &jwttoken.Sign{
		AccessToken:  *body.AccessToken,
		RefreshToken: *body.RefreshToken,
		Success:      *body.Success,
	}

	return v
}

// NewAuthProvidersEmailAlreadyExist builds a jwtToken service auth-providers
// endpoint email_already_exist error.
func NewAuthProvidersEmailAlreadyExist(body *AuthProvidersEmailAlreadyExistResponseBody) *jwttoken.EmailAlreadyExist {
	v := &jwttoken.EmailAlreadyExist{
		Err:     *body.Err,
		Success: *body.Success,
	}

	return v
}

// NewAuthProvidersUnknownError builds a jwtToken service auth-providers
// endpoint unknown_error error.
func NewAuthProvidersUnknownError(body *AuthProvidersUnknownErrorResponseBody) *jwttoken.UnknownError {
	v := &jwttoken.UnknownError{
		Err:       *body.Err,
		ErrorCode: *body.ErrorCode,
		Success:   *body.Success,
	}

	return v
}

// NewAuthProvidersInvalidScopes builds a jwtToken service auth-providers
// endpoint invalid_scopes error.
func NewAuthProvidersInvalidScopes(body AuthProvidersInvalidScopesResponseBody) jwttoken.InvalidScopes {
	v := jwttoken.InvalidScopes(body)

	return v
}

// NewAuthProvidersUnauthorized builds a jwtToken service auth-providers
// endpoint unauthorized error.
func NewAuthProvidersUnauthorized(body AuthProvidersUnauthorizedResponseBody) jwttoken.Unauthorized {
	v := jwttoken.Unauthorized(body)

	return v
}

// ValidateSignupResponseBody runs the validations defined on SignupResponseBody
func ValidateSignupResponseBody(body *SignupResponseBody) (err error) {
	if body.AccessToken == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("access_token", "body"))
	}
	if body.RefreshToken == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("refresh_token", "body"))
	}
	if body.Success == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("success", "body"))
	}
	return
}

// ValidateSigninResponseBody runs the validations defined on SigninResponseBody
func ValidateSigninResponseBody(body *SigninResponseBody) (err error) {
	if body.AccessToken == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("access_token", "body"))
	}
	if body.RefreshToken == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("refresh_token", "body"))
	}
	if body.Success == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("success", "body"))
	}
	return
}

// ValidateRefreshResponseBody runs the validations defined on
// RefreshResponseBody
func ValidateRefreshResponseBody(body *RefreshResponseBody) (err error) {
	if body.AccessToken == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("access_token", "body"))
	}
	if body.RefreshToken == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("refresh_token", "body"))
	}
	if body.Success == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("success", "body"))
	}
	return
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

// ValidateAuthProvidersCreatedResponseBody runs the validations defined on
// Auth-ProvidersCreatedResponseBody
func ValidateAuthProvidersCreatedResponseBody(body *AuthProvidersCreatedResponseBody) (err error) {
	if body.AccessToken == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("access_token", "body"))
	}
	if body.RefreshToken == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("refresh_token", "body"))
	}
	if body.Success == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("success", "body"))
	}
	return
}

// ValidateSignupEmailAlreadyExistResponseBody runs the validations defined on
// signup_email_already_exist_response_body
func ValidateSignupEmailAlreadyExistResponseBody(body *SignupEmailAlreadyExistResponseBody) (err error) {
	if body.Err == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("err", "body"))
	}
	if body.Success == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("success", "body"))
	}
	return
}

// ValidateSignupUnknownErrorResponseBody runs the validations defined on
// signup_unknown_error_response_body
func ValidateSignupUnknownErrorResponseBody(body *SignupUnknownErrorResponseBody) (err error) {
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

// ValidateSigninEmailAlreadyExistResponseBody runs the validations defined on
// signin_email_already_exist_response_body
func ValidateSigninEmailAlreadyExistResponseBody(body *SigninEmailAlreadyExistResponseBody) (err error) {
	if body.Err == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("err", "body"))
	}
	if body.Success == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("success", "body"))
	}
	return
}

// ValidateSigninUnknownErrorResponseBody runs the validations defined on
// signin_unknown_error_response_body
func ValidateSigninUnknownErrorResponseBody(body *SigninUnknownErrorResponseBody) (err error) {
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

// ValidateRefreshEmailAlreadyExistResponseBody runs the validations defined on
// refresh_email_already_exist_response_body
func ValidateRefreshEmailAlreadyExistResponseBody(body *RefreshEmailAlreadyExistResponseBody) (err error) {
	if body.Err == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("err", "body"))
	}
	if body.Success == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("success", "body"))
	}
	return
}

// ValidateRefreshUnknownErrorResponseBody runs the validations defined on
// refresh_unknown_error_response_body
func ValidateRefreshUnknownErrorResponseBody(body *RefreshUnknownErrorResponseBody) (err error) {
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

// ValidateEmailExistEmailAlreadyExistResponseBody runs the validations defined
// on email-exist_email_already_exist_response_body
func ValidateEmailExistEmailAlreadyExistResponseBody(body *EmailExistEmailAlreadyExistResponseBody) (err error) {
	if body.Err == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("err", "body"))
	}
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

// ValidateAuthProvidersEmailAlreadyExistResponseBody runs the validations
// defined on auth-providers_email_already_exist_response_body
func ValidateAuthProvidersEmailAlreadyExistResponseBody(body *AuthProvidersEmailAlreadyExistResponseBody) (err error) {
	if body.Err == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("err", "body"))
	}
	if body.Success == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("success", "body"))
	}
	return
}

// ValidateAuthProvidersUnknownErrorResponseBody runs the validations defined
// on auth-providers_unknown_error_response_body
func ValidateAuthProvidersUnknownErrorResponseBody(body *AuthProvidersUnknownErrorResponseBody) (err error) {
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

// ValidateAuthProvidersBadRequestResponseBody runs the validations defined on
// Auth-ProvidersBad RequestResponseBody
func ValidateAuthProvidersBadRequestResponseBody(body *AuthProvidersBadRequestResponseBody) (err error) {
	if body.AccessToken == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("access_token", "body"))
	}
	if body.RefreshToken == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("refresh_token", "body"))
	}
	if body.Success == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("success", "body"))
	}
	return
}
