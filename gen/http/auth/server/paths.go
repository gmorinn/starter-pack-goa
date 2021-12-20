// Code generated by goa v3.5.2, DO NOT EDIT.
//
// HTTP request path constructors for the auth service.
//
// Command:
// $ goa gen api_crud/design

package server

// EmailExistAuthPath returns the URL path to the auth service email-exist HTTP endpoint.
func EmailExistAuthPath() string {
	return "/v1/email-exist"
}

// SendConfirmationAuthPath returns the URL path to the auth service send-confirmation HTTP endpoint.
func SendConfirmationAuthPath() string {
	return "/v1/lost"
}

// ResetPasswordAuthPath returns the URL path to the auth service reset-password HTTP endpoint.
func ResetPasswordAuthPath() string {
	return "/v1/reset-password"
}
