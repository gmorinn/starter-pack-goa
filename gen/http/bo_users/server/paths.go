// Code generated by goa v3.5.2, DO NOT EDIT.
//
// HTTP request path constructors for the boUsers service.
//
// Command:
// $ goa gen api_crud/design

package server

import (
	"fmt"
)

// GetAllusersBoUsersPath returns the URL path to the boUsers service getAllusers HTTP endpoint.
func GetAllusersBoUsersPath() string {
	return "/v1/bo/users"
}

// DeleteUserBoUsersPath returns the URL path to the boUsers service deleteUser HTTP endpoint.
func DeleteUserBoUsersPath(id string) string {
	return fmt.Sprintf("/v1/bo/user/remove/%v", id)
}

// CreateUserBoUsersPath returns the URL path to the boUsers service createUser HTTP endpoint.
func CreateUserBoUsersPath() string {
	return "/v1/bo/user/add"
}

// UpdateUserBoUsersPath returns the URL path to the boUsers service updateUser HTTP endpoint.
func UpdateUserBoUsersPath(id string) string {
	return fmt.Sprintf("/v1/bo/user/%v", id)
}

// GetUserBoUsersPath returns the URL path to the boUsers service getUser HTTP endpoint.
func GetUserBoUsersPath(id string) string {
	return fmt.Sprintf("/v1/bo/user/%v", id)
}

// DeleteManyUsersBoUsersPath returns the URL path to the boUsers service deleteManyUsers HTTP endpoint.
func DeleteManyUsersBoUsersPath() string {
	return "/v1/bo/users/remove"
}

// NewPasswordBoUsersPath returns the URL path to the boUsers service newPassword HTTP endpoint.
func NewPasswordBoUsersPath(id string) string {
	return fmt.Sprintf("/v1/bo/user/change/password/%v", id)
}
