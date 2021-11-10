// Code generated by goa v3.5.2, DO NOT EDIT.
//
// users HTTP client types
//
// Command:
// $ goa gen api_crud/design

package client

import (
	users "api_crud/gen/users"

	goa "goa.design/goa/v3/pkg"
)

// CreateUserRequestBody is the type of the "users" service "createUser"
// endpoint HTTP request body.
type CreateUserRequestBody struct {
	User *PayloadUserRequestBody `form:"user" json:"user" xml:"user"`
}

// UpdateUserRequestBody is the type of the "users" service "updateUser"
// endpoint HTTP request body.
type UpdateUserRequestBody struct {
	User *PayloadUserRequestBody `form:"User" json:"User" xml:"User"`
}

// GetAllusersResponseBody is the type of the "users" service "getAllusers"
// endpoint HTTP response body.
type GetAllusersResponseBody struct {
	// All users by category
	Users   []*ResUserResponseBody `form:"users,omitempty" json:"users,omitempty" xml:"users,omitempty"`
	Success *bool                  `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
}

// DeleteUserResponseBody is the type of the "users" service "deleteUser"
// endpoint HTTP response body.
type DeleteUserResponseBody struct {
	Success *bool `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
}

// CreateUserResponseBody is the type of the "users" service "createUser"
// endpoint HTTP response body.
type CreateUserResponseBody struct {
	// Result is an object
	User    *ResUserResponseBody `form:"user,omitempty" json:"user,omitempty" xml:"user,omitempty"`
	Success *bool                `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
}

// UpdateUserResponseBody is the type of the "users" service "updateUser"
// endpoint HTTP response body.
type UpdateUserResponseBody struct {
	// Result is an Object
	User    *ResUserResponseBody `form:"user,omitempty" json:"user,omitempty" xml:"user,omitempty"`
	Success *bool                `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
}

// GetUserResponseBody is the type of the "users" service "getUser" endpoint
// HTTP response body.
type GetUserResponseBody struct {
	// Result is an object
	User    *ResUserResponseBody `form:"user,omitempty" json:"user,omitempty" xml:"user,omitempty"`
	Success *bool                `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
}

// GetAllusersUnknownErrorResponseBody is the type of the "users" service
// "getAllusers" endpoint HTTP response body for the "unknown_error" error.
type GetAllusersUnknownErrorResponseBody struct {
	Err       *string `form:"err,omitempty" json:"err,omitempty" xml:"err,omitempty"`
	ErrorCode *string `form:"error_code,omitempty" json:"error_code,omitempty" xml:"error_code,omitempty"`
	Success   *bool   `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
}

// DeleteUserUnknownErrorResponseBody is the type of the "users" service
// "deleteUser" endpoint HTTP response body for the "unknown_error" error.
type DeleteUserUnknownErrorResponseBody struct {
	Err       *string `form:"err,omitempty" json:"err,omitempty" xml:"err,omitempty"`
	ErrorCode *string `form:"error_code,omitempty" json:"error_code,omitempty" xml:"error_code,omitempty"`
	Success   *bool   `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
}

// CreateUserUnknownErrorResponseBody is the type of the "users" service
// "createUser" endpoint HTTP response body for the "unknown_error" error.
type CreateUserUnknownErrorResponseBody struct {
	Err       *string `form:"err,omitempty" json:"err,omitempty" xml:"err,omitempty"`
	ErrorCode *string `form:"error_code,omitempty" json:"error_code,omitempty" xml:"error_code,omitempty"`
	Success   *bool   `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
}

// UpdateUserUnknownErrorResponseBody is the type of the "users" service
// "updateUser" endpoint HTTP response body for the "unknown_error" error.
type UpdateUserUnknownErrorResponseBody struct {
	Err       *string `form:"err,omitempty" json:"err,omitempty" xml:"err,omitempty"`
	ErrorCode *string `form:"error_code,omitempty" json:"error_code,omitempty" xml:"error_code,omitempty"`
	Success   *bool   `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
}

// GetUserUnknownErrorResponseBody is the type of the "users" service "getUser"
// endpoint HTTP response body for the "unknown_error" error.
type GetUserUnknownErrorResponseBody struct {
	Err       *string `form:"err,omitempty" json:"err,omitempty" xml:"err,omitempty"`
	ErrorCode *string `form:"error_code,omitempty" json:"error_code,omitempty" xml:"error_code,omitempty"`
	Success   *bool   `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
}

// ResUserResponseBody is used to define fields on response body types.
type ResUserResponseBody struct {
	ID        *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Firstname *string `form:"firstname,omitempty" json:"firstname,omitempty" xml:"firstname,omitempty"`
	Lastname  *string `form:"lastname,omitempty" json:"lastname,omitempty" xml:"lastname,omitempty"`
	Email     *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	Birthday  *string `form:"birthday,omitempty" json:"birthday,omitempty" xml:"birthday,omitempty"`
	Phone     *string `form:"phone,omitempty" json:"phone,omitempty" xml:"phone,omitempty"`
}

// PayloadUserRequestBody is used to define fields on request body types.
type PayloadUserRequestBody struct {
	Firstname *string `form:"firstname,omitempty" json:"firstname,omitempty" xml:"firstname,omitempty"`
	Lastname  *string `form:"lastname,omitempty" json:"lastname,omitempty" xml:"lastname,omitempty"`
	Email     *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	Birthday  string  `form:"birthday" json:"birthday" xml:"birthday"`
	Phone     string  `form:"phone" json:"phone" xml:"phone"`
}

// NewCreateUserRequestBody builds the HTTP request body from the payload of
// the "createUser" endpoint of the "users" service.
func NewCreateUserRequestBody(p *users.CreateUserPayload) *CreateUserRequestBody {
	body := &CreateUserRequestBody{}
	if p.User != nil {
		body.User = marshalUsersPayloadUserToPayloadUserRequestBody(p.User)
	}
	return body
}

// NewUpdateUserRequestBody builds the HTTP request body from the payload of
// the "updateUser" endpoint of the "users" service.
func NewUpdateUserRequestBody(p *users.UpdateUserPayload) *UpdateUserRequestBody {
	body := &UpdateUserRequestBody{}
	if p.User != nil {
		body.User = marshalUsersPayloadUserToPayloadUserRequestBody(p.User)
	}
	return body
}

// NewGetAllusersResultOK builds a "users" service "getAllusers" endpoint
// result from a HTTP "OK" response.
func NewGetAllusersResultOK(body *GetAllusersResponseBody) *users.GetAllusersResult {
	v := &users.GetAllusersResult{
		Success: *body.Success,
	}
	v.Users = make([]*users.ResUser, len(body.Users))
	for i, val := range body.Users {
		v.Users[i] = unmarshalResUserResponseBodyToUsersResUser(val)
	}

	return v
}

// NewGetAllusersUnknownError builds a users service getAllusers endpoint
// unknown_error error.
func NewGetAllusersUnknownError(body *GetAllusersUnknownErrorResponseBody) *users.UnknownError {
	v := &users.UnknownError{
		Err:       *body.Err,
		ErrorCode: *body.ErrorCode,
		Success:   *body.Success,
	}

	return v
}

// NewDeleteUserResultOK builds a "users" service "deleteUser" endpoint result
// from a HTTP "OK" response.
func NewDeleteUserResultOK(body *DeleteUserResponseBody) *users.DeleteUserResult {
	v := &users.DeleteUserResult{
		Success: *body.Success,
	}

	return v
}

// NewDeleteUserUnknownError builds a users service deleteUser endpoint
// unknown_error error.
func NewDeleteUserUnknownError(body *DeleteUserUnknownErrorResponseBody) *users.UnknownError {
	v := &users.UnknownError{
		Err:       *body.Err,
		ErrorCode: *body.ErrorCode,
		Success:   *body.Success,
	}

	return v
}

// NewCreateUserResultCreated builds a "users" service "createUser" endpoint
// result from a HTTP "Created" response.
func NewCreateUserResultCreated(body *CreateUserResponseBody) *users.CreateUserResult {
	v := &users.CreateUserResult{
		Success: *body.Success,
	}
	v.User = unmarshalResUserResponseBodyToUsersResUser(body.User)

	return v
}

// NewCreateUserUnknownError builds a users service createUser endpoint
// unknown_error error.
func NewCreateUserUnknownError(body *CreateUserUnknownErrorResponseBody) *users.UnknownError {
	v := &users.UnknownError{
		Err:       *body.Err,
		ErrorCode: *body.ErrorCode,
		Success:   *body.Success,
	}

	return v
}

// NewUpdateUserResultOK builds a "users" service "updateUser" endpoint result
// from a HTTP "OK" response.
func NewUpdateUserResultOK(body *UpdateUserResponseBody) *users.UpdateUserResult {
	v := &users.UpdateUserResult{
		Success: *body.Success,
	}
	v.User = unmarshalResUserResponseBodyToUsersResUser(body.User)

	return v
}

// NewUpdateUserUnknownError builds a users service updateUser endpoint
// unknown_error error.
func NewUpdateUserUnknownError(body *UpdateUserUnknownErrorResponseBody) *users.UnknownError {
	v := &users.UnknownError{
		Err:       *body.Err,
		ErrorCode: *body.ErrorCode,
		Success:   *body.Success,
	}

	return v
}

// NewGetUserResultOK builds a "users" service "getUser" endpoint result from a
// HTTP "OK" response.
func NewGetUserResultOK(body *GetUserResponseBody) *users.GetUserResult {
	v := &users.GetUserResult{
		Success: *body.Success,
	}
	v.User = unmarshalResUserResponseBodyToUsersResUser(body.User)

	return v
}

// NewGetUserUnknownError builds a users service getUser endpoint unknown_error
// error.
func NewGetUserUnknownError(body *GetUserUnknownErrorResponseBody) *users.UnknownError {
	v := &users.UnknownError{
		Err:       *body.Err,
		ErrorCode: *body.ErrorCode,
		Success:   *body.Success,
	}

	return v
}

// ValidateGetAllusersResponseBody runs the validations defined on
// GetAllusersResponseBody
func ValidateGetAllusersResponseBody(body *GetAllusersResponseBody) (err error) {
	if body.Users == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("users", "body"))
	}
	if body.Success == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("success", "body"))
	}
	for _, e := range body.Users {
		if e != nil {
			if err2 := ValidateResUserResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateDeleteUserResponseBody runs the validations defined on
// DeleteUserResponseBody
func ValidateDeleteUserResponseBody(body *DeleteUserResponseBody) (err error) {
	if body.Success == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("success", "body"))
	}
	return
}

// ValidateCreateUserResponseBody runs the validations defined on
// CreateUserResponseBody
func ValidateCreateUserResponseBody(body *CreateUserResponseBody) (err error) {
	if body.User == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("user", "body"))
	}
	if body.Success == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("success", "body"))
	}
	if body.User != nil {
		if err2 := ValidateResUserResponseBody(body.User); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateUpdateUserResponseBody runs the validations defined on
// UpdateUserResponseBody
func ValidateUpdateUserResponseBody(body *UpdateUserResponseBody) (err error) {
	if body.User == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("user", "body"))
	}
	if body.Success == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("success", "body"))
	}
	if body.User != nil {
		if err2 := ValidateResUserResponseBody(body.User); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateGetUserResponseBody runs the validations defined on
// GetUserResponseBody
func ValidateGetUserResponseBody(body *GetUserResponseBody) (err error) {
	if body.User == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("user", "body"))
	}
	if body.Success == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("success", "body"))
	}
	if body.User != nil {
		if err2 := ValidateResUserResponseBody(body.User); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateGetAllusersUnknownErrorResponseBody runs the validations defined on
// getAllusers_unknown_error_response_body
func ValidateGetAllusersUnknownErrorResponseBody(body *GetAllusersUnknownErrorResponseBody) (err error) {
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

// ValidateDeleteUserUnknownErrorResponseBody runs the validations defined on
// deleteUser_unknown_error_response_body
func ValidateDeleteUserUnknownErrorResponseBody(body *DeleteUserUnknownErrorResponseBody) (err error) {
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

// ValidateCreateUserUnknownErrorResponseBody runs the validations defined on
// createUser_unknown_error_response_body
func ValidateCreateUserUnknownErrorResponseBody(body *CreateUserUnknownErrorResponseBody) (err error) {
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

// ValidateUpdateUserUnknownErrorResponseBody runs the validations defined on
// updateUser_unknown_error_response_body
func ValidateUpdateUserUnknownErrorResponseBody(body *UpdateUserUnknownErrorResponseBody) (err error) {
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

// ValidateGetUserUnknownErrorResponseBody runs the validations defined on
// getUser_unknown_error_response_body
func ValidateGetUserUnknownErrorResponseBody(body *GetUserUnknownErrorResponseBody) (err error) {
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

// ValidateResUserResponseBody runs the validations defined on
// resUserResponseBody
func ValidateResUserResponseBody(body *ResUserResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Email == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	if body.Email != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.email", *body.Email, goa.FormatEmail))
	}
	return
}

// ValidatePayloadUserRequestBody runs the validations defined on
// payloadUserRequestBody
func ValidatePayloadUserRequestBody(body *PayloadUserRequestBody) (err error) {
	if body.Email != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.email", *body.Email, goa.FormatEmail))
	}
	return
}