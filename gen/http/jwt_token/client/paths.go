// Code generated by goa v3.5.2, DO NOT EDIT.
//
// HTTP request path constructors for the jwtToken service.
//
// Command:
// $ goa gen api_crud/design

package client

// SignupJWTTokenPath returns the URL path to the jwtToken service signup HTTP endpoint.
func SignupJWTTokenPath() string {
	return "/signup"
}

// SigninJWTTokenPath returns the URL path to the jwtToken service signin HTTP endpoint.
func SigninJWTTokenPath() string {
	return "/signin"
}

// SigninBoJWTTokenPath returns the URL path to the jwtToken service signin Bo HTTP endpoint.
func SigninBoJWTTokenPath() string {
	return "/bo/signin"
}

// RefreshJWTTokenPath returns the URL path to the jwtToken service refresh HTTP endpoint.
func RefreshJWTTokenPath() string {
	return "/resfresh"
}

// AuthProvidersJWTTokenPath returns the URL path to the jwtToken service auth-providers HTTP endpoint.
func AuthProvidersJWTTokenPath() string {
	return "/sign-providers"
}
