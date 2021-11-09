// Code generated by goa v3.5.2, DO NOT EDIT.
//
// users HTTP server types
//
// Command:
// $ goa gen api_crud/design

package server

import (
	users "api_crud/gen/users"

	goa "goa.design/goa/v3/pkg"
)

// CreateUserRequestBody is the type of the "users" service "createUser"
// endpoint HTTP request body.
type CreateUserRequestBody struct {
	User *PayloadUserRequestBody `form:"user,omitempty" json:"user,omitempty" xml:"user,omitempty"`
}

// UpdateUserRequestBody is the type of the "users" service "updateUser"
// endpoint HTTP request body.
type UpdateUserRequestBody struct {
	User *PayloadUserRequestBody `form:"User,omitempty" json:"User,omitempty" xml:"User,omitempty"`
}

// GetAllusersResponseBody is the type of the "users" service "getAllusers"
// endpoint HTTP response body.
type GetAllusersResponseBody struct {
	// All users by category
	Users   []*ResUserResponseBody `form:"users" json:"users" xml:"users"`
	Success bool                   `form:"success" json:"success" xml:"success"`
}

// DeleteUserResponseBody is the type of the "users" service "deleteUser"
// endpoint HTTP response body.
type DeleteUserResponseBody struct {
	Success bool `form:"success" json:"success" xml:"success"`
}

// CreateUserResponseBody is the type of the "users" service "createUser"
// endpoint HTTP response body.
type CreateUserResponseBody struct {
	// Result is an object
	User    *ResUserResponseBody `form:"user" json:"user" xml:"user"`
	Success bool                 `form:"success" json:"success" xml:"success"`
}

// UpdateUserResponseBody is the type of the "users" service "updateUser"
// endpoint HTTP response body.
type UpdateUserResponseBody struct {
	// Result is an Object
	User    *ResUserResponseBody `form:"user" json:"user" xml:"user"`
	Success bool                 `form:"success" json:"success" xml:"success"`
}

// GetUserResponseBody is the type of the "users" service "getUser" endpoint
// HTTP response body.
type GetUserResponseBody struct {
	// Result is an object
	User    *ResUserResponseBody `form:"user" json:"user" xml:"user"`
	Success bool                 `form:"success" json:"success" xml:"success"`
}

// GetAllusersUnknownErrorResponseBody is the type of the "users" service
// "getAllusers" endpoint HTTP response body for the "unknown_error" error.
type GetAllusersUnknownErrorResponseBody struct {
	Err       string `form:"err" json:"err" xml:"err"`
	ErrorCode string `form:"error_code" json:"error_code" xml:"error_code"`
	Success   bool   `form:"success" json:"success" xml:"success"`
}

// DeleteUserUnknownErrorResponseBody is the type of the "users" service
// "deleteUser" endpoint HTTP response body for the "unknown_error" error.
type DeleteUserUnknownErrorResponseBody struct {
	Err       string `form:"err" json:"err" xml:"err"`
	ErrorCode string `form:"error_code" json:"error_code" xml:"error_code"`
	Success   bool   `form:"success" json:"success" xml:"success"`
}

// CreateUserUnknownErrorResponseBody is the type of the "users" service
// "createUser" endpoint HTTP response body for the "unknown_error" error.
type CreateUserUnknownErrorResponseBody struct {
	Err       string `form:"err" json:"err" xml:"err"`
	ErrorCode string `form:"error_code" json:"error_code" xml:"error_code"`
	Success   bool   `form:"success" json:"success" xml:"success"`
}

// UpdateUserUnknownErrorResponseBody is the type of the "users" service
// "updateUser" endpoint HTTP response body for the "unknown_error" error.
type UpdateUserUnknownErrorResponseBody struct {
	Err       string `form:"err" json:"err" xml:"err"`
	ErrorCode string `form:"error_code" json:"error_code" xml:"error_code"`
	Success   bool   `form:"success" json:"success" xml:"success"`
}

// GetUserUnknownErrorResponseBody is the type of the "users" service "getUser"
// endpoint HTTP response body for the "unknown_error" error.
type GetUserUnknownErrorResponseBody struct {
	Err       string `form:"err" json:"err" xml:"err"`
	ErrorCode string `form:"error_code" json:"error_code" xml:"error_code"`
	Success   bool   `form:"success" json:"success" xml:"success"`
}

// ResUserResponseBody is used to define fields on response body types.
type ResUserResponseBody struct {
	ID        string  `form:"id" json:"id" xml:"id"`
	Firstname *string `form:"firstname,omitempty" json:"firstname,omitempty" xml:"firstname,omitempty"`
	Lastname  *string `form:"lastname,omitempty" json:"lastname,omitempty" xml:"lastname,omitempty"`
	Email     string  `form:"email" json:"email" xml:"email"`
	Birthday  string  `form:"birthday" json:"birthday" xml:"birthday"`
	Phone     string  `form:"phone" json:"phone" xml:"phone"`
}

// PayloadUserRequestBody is used to define fields on request body types.
type PayloadUserRequestBody struct {
	Firstname *string `form:"firstname,omitempty" json:"firstname,omitempty" xml:"firstname,omitempty"`
	Lastname  *string `form:"lastname,omitempty" json:"lastname,omitempty" xml:"lastname,omitempty"`
	Email     *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	Birthday  *string `form:"birthday,omitempty" json:"birthday,omitempty" xml:"birthday,omitempty"`
	Phone     *string `form:"phone,omitempty" json:"phone,omitempty" xml:"phone,omitempty"`
}

// NewGetAllusersResponseBody builds the HTTP response body from the result of
// the "getAllusers" endpoint of the "users" service.
func NewGetAllusersResponseBody(res *users.GetAllusersResult) *GetAllusersResponseBody {
	body := &GetAllusersResponseBody{
		Success: res.Success,
	}
	if res.Users != nil {
		body.Users = make([]*ResUserResponseBody, len(res.Users))
		for i, val := range res.Users {
			body.Users[i] = marshalUsersResUserToResUserResponseBody(val)
		}
	}
	return body
}

// NewDeleteUserResponseBody builds the HTTP response body from the result of
// the "deleteUser" endpoint of the "users" service.
func NewDeleteUserResponseBody(res *users.DeleteUserResult) *DeleteUserResponseBody {
	body := &DeleteUserResponseBody{
		Success: res.Success,
	}
	return body
}

// NewCreateUserResponseBody builds the HTTP response body from the result of
// the "createUser" endpoint of the "users" service.
func NewCreateUserResponseBody(res *users.CreateUserResult) *CreateUserResponseBody {
	body := &CreateUserResponseBody{
		Success: res.Success,
	}
	if res.User != nil {
		body.User = marshalUsersResUserToResUserResponseBody(res.User)
	}
	return body
}

// NewUpdateUserResponseBody builds the HTTP response body from the result of
// the "updateUser" endpoint of the "users" service.
func NewUpdateUserResponseBody(res *users.UpdateUserResult) *UpdateUserResponseBody {
	body := &UpdateUserResponseBody{
		Success: res.Success,
	}
	if res.User != nil {
		body.User = marshalUsersResUserToResUserResponseBody(res.User)
	}
	return body
}

// NewGetUserResponseBody builds the HTTP response body from the result of the
// "getUser" endpoint of the "users" service.
func NewGetUserResponseBody(res *users.GetUserResult) *GetUserResponseBody {
	body := &GetUserResponseBody{
		Success: res.Success,
	}
	if res.User != nil {
		body.User = marshalUsersResUserToResUserResponseBody(res.User)
	}
	return body
}

// NewGetAllusersUnknownErrorResponseBody builds the HTTP response body from
// the result of the "getAllusers" endpoint of the "users" service.
func NewGetAllusersUnknownErrorResponseBody(res *users.UnknownError) *GetAllusersUnknownErrorResponseBody {
	body := &GetAllusersUnknownErrorResponseBody{
		Err:       res.Err,
		ErrorCode: res.ErrorCode,
		Success:   res.Success,
	}
	return body
}

// NewDeleteUserUnknownErrorResponseBody builds the HTTP response body from the
// result of the "deleteUser" endpoint of the "users" service.
func NewDeleteUserUnknownErrorResponseBody(res *users.UnknownError) *DeleteUserUnknownErrorResponseBody {
	body := &DeleteUserUnknownErrorResponseBody{
		Err:       res.Err,
		ErrorCode: res.ErrorCode,
		Success:   res.Success,
	}
	return body
}

// NewCreateUserUnknownErrorResponseBody builds the HTTP response body from the
// result of the "createUser" endpoint of the "users" service.
func NewCreateUserUnknownErrorResponseBody(res *users.UnknownError) *CreateUserUnknownErrorResponseBody {
	body := &CreateUserUnknownErrorResponseBody{
		Err:       res.Err,
		ErrorCode: res.ErrorCode,
		Success:   res.Success,
	}
	return body
}

// NewUpdateUserUnknownErrorResponseBody builds the HTTP response body from the
// result of the "updateUser" endpoint of the "users" service.
func NewUpdateUserUnknownErrorResponseBody(res *users.UnknownError) *UpdateUserUnknownErrorResponseBody {
	body := &UpdateUserUnknownErrorResponseBody{
		Err:       res.Err,
		ErrorCode: res.ErrorCode,
		Success:   res.Success,
	}
	return body
}

// NewGetUserUnknownErrorResponseBody builds the HTTP response body from the
// result of the "getUser" endpoint of the "users" service.
func NewGetUserUnknownErrorResponseBody(res *users.UnknownError) *GetUserUnknownErrorResponseBody {
	body := &GetUserUnknownErrorResponseBody{
		Err:       res.Err,
		ErrorCode: res.ErrorCode,
		Success:   res.Success,
	}
	return body
}

// NewGetAllusersPayload builds a users service getAllusers endpoint payload.
func NewGetAllusersPayload(oauth *string, jwtToken *string) *users.GetAllusersPayload {
	v := &users.GetAllusersPayload{}
	v.Oauth = oauth
	v.JWTToken = jwtToken

	return v
}

// NewDeleteUserPayload builds a users service deleteUser endpoint payload.
func NewDeleteUserPayload(id string, oauth *string, jwtToken *string) *users.DeleteUserPayload {
	v := &users.DeleteUserPayload{}
	v.ID = id
	v.Oauth = oauth
	v.JWTToken = jwtToken

	return v
}

// NewCreateUserPayload builds a users service createUser endpoint payload.
func NewCreateUserPayload(body *CreateUserRequestBody, oauth *string, jwtToken *string) *users.CreateUserPayload {
	v := &users.CreateUserPayload{}
	v.User = unmarshalPayloadUserRequestBodyToUsersPayloadUser(body.User)
	v.Oauth = oauth
	v.JWTToken = jwtToken

	return v
}

// NewUpdateUserPayload builds a users service updateUser endpoint payload.
func NewUpdateUserPayload(body *UpdateUserRequestBody, id string, oauth *string, jwtToken *string) *users.UpdateUserPayload {
	v := &users.UpdateUserPayload{}
	v.User = unmarshalPayloadUserRequestBodyToUsersPayloadUser(body.User)
	v.ID = id
	v.Oauth = oauth
	v.JWTToken = jwtToken

	return v
}

// NewGetUserPayload builds a users service getUser endpoint payload.
func NewGetUserPayload(id string, oauth *string, jwtToken *string) *users.GetUserPayload {
	v := &users.GetUserPayload{}
	v.ID = id
	v.Oauth = oauth
	v.JWTToken = jwtToken

	return v
}

// ValidateCreateUserRequestBody runs the validations defined on
// CreateUserRequestBody
func ValidateCreateUserRequestBody(body *CreateUserRequestBody) (err error) {
	if body.User == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("user", "body"))
	}
	if body.User != nil {
		if err2 := ValidatePayloadUserRequestBody(body.User); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateUpdateUserRequestBody runs the validations defined on
// UpdateUserRequestBody
func ValidateUpdateUserRequestBody(body *UpdateUserRequestBody) (err error) {
	if body.User == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("User", "body"))
	}
	if body.User != nil {
		if err2 := ValidatePayloadUserRequestBody(body.User); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
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
