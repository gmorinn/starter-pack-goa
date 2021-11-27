// Code generated by goa v3.5.2, DO NOT EDIT.
//
// jwtToken HTTP server types
//
// Command:
// $ goa gen api_crud/design

package server

import (
	jwttoken "api_crud/gen/jwt_token"
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// SignupRequestBody is the type of the "jwtToken" service "signup" endpoint
// HTTP request body.
type SignupRequestBody struct {
	Firstname *string `form:"firstname,omitempty" json:"firstname,omitempty" xml:"firstname,omitempty"`
	Lastname  *string `form:"lastname,omitempty" json:"lastname,omitempty" xml:"lastname,omitempty"`
	// Minimum 8 charactères / Chiffre Obligatoire
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
	// Minimum 8 charactères / Chiffre Obligatoire
	ConfirmPassword *string `form:"confirm_password,omitempty" json:"confirm_password,omitempty" xml:"confirm_password,omitempty"`
	Email           *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	Birthday        *string `form:"birthday,omitempty" json:"birthday,omitempty" xml:"birthday,omitempty"`
	Phone           *string `form:"phone,omitempty" json:"phone,omitempty" xml:"phone,omitempty"`
}

// SigninRequestBody is the type of the "jwtToken" service "signin" endpoint
// HTTP request body.
type SigninRequestBody struct {
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Minimum 8 charactères / Chiffre Obligatoire
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
}

// RefreshRequestBody is the type of the "jwtToken" service "refresh" endpoint
// HTTP request body.
type RefreshRequestBody struct {
	RefreshToken *string `form:"refresh_token,omitempty" json:"refresh_token,omitempty" xml:"refresh_token,omitempty"`
}

// EmailExistRequestBody is the type of the "jwtToken" service "email-exist"
// endpoint HTTP request body.
type EmailExistRequestBody struct {
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
}

// AuthProvidersRequestBody is the type of the "jwtToken" service
// "auth-providers" endpoint HTTP request body.
type AuthProvidersRequestBody struct {
	Firstname        *string `form:"firstname,omitempty" json:"firstname,omitempty" xml:"firstname,omitempty"`
	Lastname         *string `form:"lastname,omitempty" json:"lastname,omitempty" xml:"lastname,omitempty"`
	Email            *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	FirebaseIDToken  *string `form:"firebase_id_token,omitempty" json:"firebase_id_token,omitempty" xml:"firebase_id_token,omitempty"`
	FirebaseUID      *string `form:"firebase_uid,omitempty" json:"firebase_uid,omitempty" xml:"firebase_uid,omitempty"`
	FirebaseProvider *string `form:"firebase_provider,omitempty" json:"firebase_provider,omitempty" xml:"firebase_provider,omitempty"`
}

// SignupResponseBody is the type of the "jwtToken" service "signup" endpoint
// HTTP response body.
type SignupResponseBody struct {
	AccessToken  string `form:"access_token" json:"access_token" xml:"access_token"`
	RefreshToken string `form:"refresh_token" json:"refresh_token" xml:"refresh_token"`
	Success      bool   `form:"success" json:"success" xml:"success"`
}

// SigninResponseBody is the type of the "jwtToken" service "signin" endpoint
// HTTP response body.
type SigninResponseBody struct {
	AccessToken  string `form:"access_token" json:"access_token" xml:"access_token"`
	RefreshToken string `form:"refresh_token" json:"refresh_token" xml:"refresh_token"`
	Success      bool   `form:"success" json:"success" xml:"success"`
}

// RefreshResponseBody is the type of the "jwtToken" service "refresh" endpoint
// HTTP response body.
type RefreshResponseBody struct {
	AccessToken  string `form:"access_token" json:"access_token" xml:"access_token"`
	RefreshToken string `form:"refresh_token" json:"refresh_token" xml:"refresh_token"`
	Success      bool   `form:"success" json:"success" xml:"success"`
}

// EmailExistResponseBody is the type of the "jwtToken" service "email-exist"
// endpoint HTTP response body.
type EmailExistResponseBody struct {
	Success bool `form:"success" json:"success" xml:"success"`
	Exist   bool `form:"exist" json:"exist" xml:"exist"`
}

// AuthProvidersCreatedResponseBody is the type of the "jwtToken" service
// "auth-providers" endpoint HTTP response body.
type AuthProvidersCreatedResponseBody struct {
	AccessToken  string `form:"access_token" json:"access_token" xml:"access_token"`
	RefreshToken string `form:"refresh_token" json:"refresh_token" xml:"refresh_token"`
	Success      bool   `form:"success" json:"success" xml:"success"`
}

// SignupEmailAlreadyExistResponseBody is the type of the "jwtToken" service
// "signup" endpoint HTTP response body for the "email_already_exist" error.
type SignupEmailAlreadyExistResponseBody struct {
	Message string `form:"message" json:"message" xml:"message"`
	Success bool   `form:"success" json:"success" xml:"success"`
}

// SignupUnknownErrorResponseBody is the type of the "jwtToken" service
// "signup" endpoint HTTP response body for the "unknown_error" error.
type SignupUnknownErrorResponseBody struct {
	Err       string `form:"err" json:"err" xml:"err"`
	ErrorCode string `form:"error_code" json:"error_code" xml:"error_code"`
	Success   bool   `form:"success" json:"success" xml:"success"`
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
	Message string `form:"message" json:"message" xml:"message"`
	Success bool   `form:"success" json:"success" xml:"success"`
}

// SigninUnknownErrorResponseBody is the type of the "jwtToken" service
// "signin" endpoint HTTP response body for the "unknown_error" error.
type SigninUnknownErrorResponseBody struct {
	Err       string `form:"err" json:"err" xml:"err"`
	ErrorCode string `form:"error_code" json:"error_code" xml:"error_code"`
	Success   bool   `form:"success" json:"success" xml:"success"`
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
	Message string `form:"message" json:"message" xml:"message"`
	Success bool   `form:"success" json:"success" xml:"success"`
}

// RefreshUnknownErrorResponseBody is the type of the "jwtToken" service
// "refresh" endpoint HTTP response body for the "unknown_error" error.
type RefreshUnknownErrorResponseBody struct {
	Err       string `form:"err" json:"err" xml:"err"`
	ErrorCode string `form:"error_code" json:"error_code" xml:"error_code"`
	Success   bool   `form:"success" json:"success" xml:"success"`
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
	Message string `form:"message" json:"message" xml:"message"`
	Success bool   `form:"success" json:"success" xml:"success"`
}

// EmailExistUnknownErrorResponseBody is the type of the "jwtToken" service
// "email-exist" endpoint HTTP response body for the "unknown_error" error.
type EmailExistUnknownErrorResponseBody struct {
	Err       string `form:"err" json:"err" xml:"err"`
	ErrorCode string `form:"error_code" json:"error_code" xml:"error_code"`
	Success   bool   `form:"success" json:"success" xml:"success"`
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
	Message string `form:"message" json:"message" xml:"message"`
	Success bool   `form:"success" json:"success" xml:"success"`
}

// AuthProvidersUnknownErrorResponseBody is the type of the "jwtToken" service
// "auth-providers" endpoint HTTP response body for the "unknown_error" error.
type AuthProvidersUnknownErrorResponseBody struct {
	Err       string `form:"err" json:"err" xml:"err"`
	ErrorCode string `form:"error_code" json:"error_code" xml:"error_code"`
	Success   bool   `form:"success" json:"success" xml:"success"`
}

// AuthProvidersInvalidScopesResponseBody is the type of the "jwtToken" service
// "auth-providers" endpoint HTTP response body for the "invalid_scopes" error.
type AuthProvidersInvalidScopesResponseBody string

// AuthProvidersUnauthorizedResponseBody is the type of the "jwtToken" service
// "auth-providers" endpoint HTTP response body for the "unauthorized" error.
type AuthProvidersUnauthorizedResponseBody string

// NewSignupResponseBody builds the HTTP response body from the result of the
// "signup" endpoint of the "jwtToken" service.
func NewSignupResponseBody(res *jwttoken.Sign) *SignupResponseBody {
	body := &SignupResponseBody{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
		Success:      res.Success,
	}
	return body
}

// NewSigninResponseBody builds the HTTP response body from the result of the
// "signin" endpoint of the "jwtToken" service.
func NewSigninResponseBody(res *jwttoken.Sign) *SigninResponseBody {
	body := &SigninResponseBody{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
		Success:      res.Success,
	}
	return body
}

// NewRefreshResponseBody builds the HTTP response body from the result of the
// "refresh" endpoint of the "jwtToken" service.
func NewRefreshResponseBody(res *jwttoken.Sign) *RefreshResponseBody {
	body := &RefreshResponseBody{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
		Success:      res.Success,
	}
	return body
}

// NewEmailExistResponseBody builds the HTTP response body from the result of
// the "email-exist" endpoint of the "jwtToken" service.
func NewEmailExistResponseBody(res *jwttoken.EmailExistResult) *EmailExistResponseBody {
	body := &EmailExistResponseBody{
		Success: res.Success,
		Exist:   res.Exist,
	}
	return body
}

// NewAuthProvidersCreatedResponseBody builds the HTTP response body from the
// result of the "auth-providers" endpoint of the "jwtToken" service.
func NewAuthProvidersCreatedResponseBody(res *jwttoken.Sign) *AuthProvidersCreatedResponseBody {
	body := &AuthProvidersCreatedResponseBody{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
		Success:      res.Success,
	}
	return body
}

// NewSignupEmailAlreadyExistResponseBody builds the HTTP response body from
// the result of the "signup" endpoint of the "jwtToken" service.
func NewSignupEmailAlreadyExistResponseBody(res *jwttoken.EmailAlreadyExist) *SignupEmailAlreadyExistResponseBody {
	body := &SignupEmailAlreadyExistResponseBody{
		Message: res.Message,
		Success: res.Success,
	}
	return body
}

// NewSignupUnknownErrorResponseBody builds the HTTP response body from the
// result of the "signup" endpoint of the "jwtToken" service.
func NewSignupUnknownErrorResponseBody(res *jwttoken.UnknownError) *SignupUnknownErrorResponseBody {
	body := &SignupUnknownErrorResponseBody{
		Err:       res.Err,
		ErrorCode: res.ErrorCode,
		Success:   res.Success,
	}
	return body
}

// NewSignupInvalidScopesResponseBody builds the HTTP response body from the
// result of the "signup" endpoint of the "jwtToken" service.
func NewSignupInvalidScopesResponseBody(res jwttoken.InvalidScopes) SignupInvalidScopesResponseBody {
	body := SignupInvalidScopesResponseBody(res)
	return body
}

// NewSignupUnauthorizedResponseBody builds the HTTP response body from the
// result of the "signup" endpoint of the "jwtToken" service.
func NewSignupUnauthorizedResponseBody(res jwttoken.Unauthorized) SignupUnauthorizedResponseBody {
	body := SignupUnauthorizedResponseBody(res)
	return body
}

// NewSigninEmailAlreadyExistResponseBody builds the HTTP response body from
// the result of the "signin" endpoint of the "jwtToken" service.
func NewSigninEmailAlreadyExistResponseBody(res *jwttoken.EmailAlreadyExist) *SigninEmailAlreadyExistResponseBody {
	body := &SigninEmailAlreadyExistResponseBody{
		Message: res.Message,
		Success: res.Success,
	}
	return body
}

// NewSigninUnknownErrorResponseBody builds the HTTP response body from the
// result of the "signin" endpoint of the "jwtToken" service.
func NewSigninUnknownErrorResponseBody(res *jwttoken.UnknownError) *SigninUnknownErrorResponseBody {
	body := &SigninUnknownErrorResponseBody{
		Err:       res.Err,
		ErrorCode: res.ErrorCode,
		Success:   res.Success,
	}
	return body
}

// NewSigninInvalidScopesResponseBody builds the HTTP response body from the
// result of the "signin" endpoint of the "jwtToken" service.
func NewSigninInvalidScopesResponseBody(res jwttoken.InvalidScopes) SigninInvalidScopesResponseBody {
	body := SigninInvalidScopesResponseBody(res)
	return body
}

// NewSigninUnauthorizedResponseBody builds the HTTP response body from the
// result of the "signin" endpoint of the "jwtToken" service.
func NewSigninUnauthorizedResponseBody(res jwttoken.Unauthorized) SigninUnauthorizedResponseBody {
	body := SigninUnauthorizedResponseBody(res)
	return body
}

// NewRefreshEmailAlreadyExistResponseBody builds the HTTP response body from
// the result of the "refresh" endpoint of the "jwtToken" service.
func NewRefreshEmailAlreadyExistResponseBody(res *jwttoken.EmailAlreadyExist) *RefreshEmailAlreadyExistResponseBody {
	body := &RefreshEmailAlreadyExistResponseBody{
		Message: res.Message,
		Success: res.Success,
	}
	return body
}

// NewRefreshUnknownErrorResponseBody builds the HTTP response body from the
// result of the "refresh" endpoint of the "jwtToken" service.
func NewRefreshUnknownErrorResponseBody(res *jwttoken.UnknownError) *RefreshUnknownErrorResponseBody {
	body := &RefreshUnknownErrorResponseBody{
		Err:       res.Err,
		ErrorCode: res.ErrorCode,
		Success:   res.Success,
	}
	return body
}

// NewRefreshInvalidScopesResponseBody builds the HTTP response body from the
// result of the "refresh" endpoint of the "jwtToken" service.
func NewRefreshInvalidScopesResponseBody(res jwttoken.InvalidScopes) RefreshInvalidScopesResponseBody {
	body := RefreshInvalidScopesResponseBody(res)
	return body
}

// NewRefreshUnauthorizedResponseBody builds the HTTP response body from the
// result of the "refresh" endpoint of the "jwtToken" service.
func NewRefreshUnauthorizedResponseBody(res jwttoken.Unauthorized) RefreshUnauthorizedResponseBody {
	body := RefreshUnauthorizedResponseBody(res)
	return body
}

// NewEmailExistEmailAlreadyExistResponseBody builds the HTTP response body
// from the result of the "email-exist" endpoint of the "jwtToken" service.
func NewEmailExistEmailAlreadyExistResponseBody(res *jwttoken.EmailAlreadyExist) *EmailExistEmailAlreadyExistResponseBody {
	body := &EmailExistEmailAlreadyExistResponseBody{
		Message: res.Message,
		Success: res.Success,
	}
	return body
}

// NewEmailExistUnknownErrorResponseBody builds the HTTP response body from the
// result of the "email-exist" endpoint of the "jwtToken" service.
func NewEmailExistUnknownErrorResponseBody(res *jwttoken.UnknownError) *EmailExistUnknownErrorResponseBody {
	body := &EmailExistUnknownErrorResponseBody{
		Err:       res.Err,
		ErrorCode: res.ErrorCode,
		Success:   res.Success,
	}
	return body
}

// NewEmailExistInvalidScopesResponseBody builds the HTTP response body from
// the result of the "email-exist" endpoint of the "jwtToken" service.
func NewEmailExistInvalidScopesResponseBody(res jwttoken.InvalidScopes) EmailExistInvalidScopesResponseBody {
	body := EmailExistInvalidScopesResponseBody(res)
	return body
}

// NewEmailExistUnauthorizedResponseBody builds the HTTP response body from the
// result of the "email-exist" endpoint of the "jwtToken" service.
func NewEmailExistUnauthorizedResponseBody(res jwttoken.Unauthorized) EmailExistUnauthorizedResponseBody {
	body := EmailExistUnauthorizedResponseBody(res)
	return body
}

// NewAuthProvidersEmailAlreadyExistResponseBody builds the HTTP response body
// from the result of the "auth-providers" endpoint of the "jwtToken" service.
func NewAuthProvidersEmailAlreadyExistResponseBody(res *jwttoken.EmailAlreadyExist) *AuthProvidersEmailAlreadyExistResponseBody {
	body := &AuthProvidersEmailAlreadyExistResponseBody{
		Message: res.Message,
		Success: res.Success,
	}
	return body
}

// NewAuthProvidersUnknownErrorResponseBody builds the HTTP response body from
// the result of the "auth-providers" endpoint of the "jwtToken" service.
func NewAuthProvidersUnknownErrorResponseBody(res *jwttoken.UnknownError) *AuthProvidersUnknownErrorResponseBody {
	body := &AuthProvidersUnknownErrorResponseBody{
		Err:       res.Err,
		ErrorCode: res.ErrorCode,
		Success:   res.Success,
	}
	return body
}

// NewAuthProvidersInvalidScopesResponseBody builds the HTTP response body from
// the result of the "auth-providers" endpoint of the "jwtToken" service.
func NewAuthProvidersInvalidScopesResponseBody(res jwttoken.InvalidScopes) AuthProvidersInvalidScopesResponseBody {
	body := AuthProvidersInvalidScopesResponseBody(res)
	return body
}

// NewAuthProvidersUnauthorizedResponseBody builds the HTTP response body from
// the result of the "auth-providers" endpoint of the "jwtToken" service.
func NewAuthProvidersUnauthorizedResponseBody(res jwttoken.Unauthorized) AuthProvidersUnauthorizedResponseBody {
	body := AuthProvidersUnauthorizedResponseBody(res)
	return body
}

// NewSignupPayload builds a jwtToken service signup endpoint payload.
func NewSignupPayload(body *SignupRequestBody, oauth *string) *jwttoken.SignupPayload {
	v := &jwttoken.SignupPayload{
		Firstname:       *body.Firstname,
		Lastname:        *body.Lastname,
		Password:        *body.Password,
		ConfirmPassword: *body.ConfirmPassword,
		Email:           *body.Email,
	}
	if body.Birthday != nil {
		v.Birthday = *body.Birthday
	}
	if body.Phone != nil {
		v.Phone = *body.Phone
	}
	if body.Birthday == nil {
		v.Birthday = ""
	}
	if body.Phone == nil {
		v.Phone = ""
	}
	v.Oauth = oauth

	return v
}

// NewSigninPayload builds a jwtToken service signin endpoint payload.
func NewSigninPayload(body *SigninRequestBody, oauth *string) *jwttoken.SigninPayload {
	v := &jwttoken.SigninPayload{
		Email:    *body.Email,
		Password: *body.Password,
	}
	v.Oauth = oauth

	return v
}

// NewRefreshPayload builds a jwtToken service refresh endpoint payload.
func NewRefreshPayload(body *RefreshRequestBody, oauth *string) *jwttoken.RefreshPayload {
	v := &jwttoken.RefreshPayload{
		RefreshToken: *body.RefreshToken,
	}
	v.Oauth = oauth

	return v
}

// NewEmailExistPayload builds a jwtToken service email-exist endpoint payload.
func NewEmailExistPayload(body *EmailExistRequestBody, oauth *string) *jwttoken.EmailExistPayload {
	v := &jwttoken.EmailExistPayload{
		Email: *body.Email,
	}
	v.Oauth = oauth

	return v
}

// NewAuthProvidersPayload builds a jwtToken service auth-providers endpoint
// payload.
func NewAuthProvidersPayload(body *AuthProvidersRequestBody, oauth *string) *jwttoken.AuthProvidersPayload {
	v := &jwttoken.AuthProvidersPayload{
		Firstname:        *body.Firstname,
		Lastname:         *body.Lastname,
		Email:            *body.Email,
		FirebaseIDToken:  *body.FirebaseIDToken,
		FirebaseUID:      *body.FirebaseUID,
		FirebaseProvider: *body.FirebaseProvider,
	}
	v.Oauth = oauth

	return v
}

// ValidateSignupRequestBody runs the validations defined on SignupRequestBody
func ValidateSignupRequestBody(body *SignupRequestBody) (err error) {
	if body.Firstname == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("firstname", "body"))
	}
	if body.Lastname == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("lastname", "body"))
	}
	if body.Password == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("password", "body"))
	}
	if body.Email == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email", "body"))
	}
	if body.ConfirmPassword == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("confirm_password", "body"))
	}
	if body.Firstname != nil {
		if utf8.RuneCountInString(*body.Firstname) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.firstname", *body.Firstname, utf8.RuneCountInString(*body.Firstname), 3, true))
		}
	}
	if body.Firstname != nil {
		if utf8.RuneCountInString(*body.Firstname) > 15 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.firstname", *body.Firstname, utf8.RuneCountInString(*body.Firstname), 15, false))
		}
	}
	if body.Lastname != nil {
		if utf8.RuneCountInString(*body.Lastname) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.lastname", *body.Lastname, utf8.RuneCountInString(*body.Lastname), 3, true))
		}
	}
	if body.Password != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.password", *body.Password, "\\d"))
	}
	if body.Password != nil {
		if utf8.RuneCountInString(*body.Password) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.password", *body.Password, utf8.RuneCountInString(*body.Password), 8, true))
		}
	}
	if body.ConfirmPassword != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.confirm_password", *body.ConfirmPassword, "\\d"))
	}
	if body.ConfirmPassword != nil {
		if utf8.RuneCountInString(*body.ConfirmPassword) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.confirm_password", *body.ConfirmPassword, utf8.RuneCountInString(*body.ConfirmPassword), 8, true))
		}
	}
	if body.Email != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.email", *body.Email, goa.FormatEmail))
	}
	return
}

// ValidateSigninRequestBody runs the validations defined on SigninRequestBody
func ValidateSigninRequestBody(body *SigninRequestBody) (err error) {
	if body.Password == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("password", "body"))
	}
	if body.Email == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email", "body"))
	}
	if body.Email != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.email", *body.Email, goa.FormatEmail))
	}
	if body.Password != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.password", *body.Password, "\\d"))
	}
	if body.Password != nil {
		if utf8.RuneCountInString(*body.Password) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.password", *body.Password, utf8.RuneCountInString(*body.Password), 8, true))
		}
	}
	return
}

// ValidateRefreshRequestBody runs the validations defined on RefreshRequestBody
func ValidateRefreshRequestBody(body *RefreshRequestBody) (err error) {
	if body.RefreshToken == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("refresh_token", "body"))
	}
	return
}

// ValidateEmailExistRequestBody runs the validations defined on
// Email-ExistRequestBody
func ValidateEmailExistRequestBody(body *EmailExistRequestBody) (err error) {
	if body.Email == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email", "body"))
	}
	if body.Email != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.email", *body.Email, goa.FormatEmail))
	}
	return
}

// ValidateAuthProvidersRequestBody runs the validations defined on
// Auth-ProvidersRequestBody
func ValidateAuthProvidersRequestBody(body *AuthProvidersRequestBody) (err error) {
	if body.Email == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email", "body"))
	}
	if body.Firstname == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("firstname", "body"))
	}
	if body.Lastname == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("lastname", "body"))
	}
	if body.FirebaseIDToken == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("firebase_id_token", "body"))
	}
	if body.FirebaseUID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("firebase_uid", "body"))
	}
	if body.FirebaseProvider == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("firebase_provider", "body"))
	}
	if body.Firstname != nil {
		if utf8.RuneCountInString(*body.Firstname) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.firstname", *body.Firstname, utf8.RuneCountInString(*body.Firstname), 3, true))
		}
	}
	if body.Firstname != nil {
		if utf8.RuneCountInString(*body.Firstname) > 15 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.firstname", *body.Firstname, utf8.RuneCountInString(*body.Firstname), 15, false))
		}
	}
	if body.Lastname != nil {
		if utf8.RuneCountInString(*body.Lastname) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.lastname", *body.Lastname, utf8.RuneCountInString(*body.Lastname), 3, true))
		}
	}
	if body.Email != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.email", *body.Email, goa.FormatEmail))
	}
	if body.FirebaseIDToken != nil {
		if utf8.RuneCountInString(*body.FirebaseIDToken) < 400 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.firebase_id_token", *body.FirebaseIDToken, utf8.RuneCountInString(*body.FirebaseIDToken), 400, true))
		}
	}
	if body.FirebaseUID != nil {
		if utf8.RuneCountInString(*body.FirebaseUID) < 15 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.firebase_uid", *body.FirebaseUID, utf8.RuneCountInString(*body.FirebaseUID), 15, true))
		}
	}
	return
}
