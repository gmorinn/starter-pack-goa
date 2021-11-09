package api

import (
	users "api_crud/gen/users"
	"context"
	"log"

	"goa.design/goa/v3/security"
)

// users service example implementation.
// The example methods log the requests and return zero values.
type userssrvc struct {
	logger *log.Logger
	server *Server
}

// NewUsers returns the users service implementation.
func NewUsers(logger *log.Logger, server *Server) users.Service {
	return &userssrvc{logger, server}
}

func (s *userssrvc) errorResponse(msg string, err error) *users.UnknownError {
	return &users.UnknownError{
		Err:       err.Error(),
		ErrorCode: msg,
	}
}

// OAuth2Auth implements the authorization logic for service "users" for the
// "OAuth2" security scheme.
func (s *userssrvc) OAuth2Auth(ctx context.Context, token string, scheme *security.OAuth2Scheme) (context.Context, error) {
	return s.server.CheckAuth(ctx, token, scheme)
	// TBD: add authorization logic.
	//
	// In case of authorization failure this function should return
	// one of the generated error structs, e.g.:
	//
	//    return ctx, myservice.MakeUnauthorizedError("invalid token")
	//
	// Alternatively this function may return an instance of
	// goa.ServiceError with a Name field value that matches one of
	// the design error names, e.g:
	//
	//    return ctx, goa.PermanentError("unauthorized", "invalid token")
	//

}

// JWTAuth implements the authorization logic for service "users" for the "jwt"
// security scheme.
func (s *userssrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	return s.server.CheckJWT(ctx, token, scheme)
	// TBD: add authorization logic.
	//
	// In case of authorization failure this function should return
	// one of the generated error structs, e.g.:
	//
	//    return ctx, myservice.MakeUnauthorizedError("invalid token")
	//
	// Alternatively this function may return an instance of
	// goa.ServiceError with a Name field value that matches one of
	// the design error names, e.g:
	//
	//    return ctx, goa.PermanentError("unauthorized", "invalid token")
	//

}

// Get All users
func (s *userssrvc) GetAllusers(ctx context.Context, p *users.GetAllusersPayload) (res *users.GetAllusersResult, err error) {
	res = &users.GetAllusersResult{}
	s.logger.Print("users.getAllusers")
	return
}

// Delete one User by ID
func (s *userssrvc) DeleteUser(ctx context.Context, p *users.DeleteUserPayload) (res *users.DeleteUserResult, err error) {
	res = &users.DeleteUserResult{}
	s.logger.Print("users.deleteUser")
	return
}

// Create one User
func (s *userssrvc) CreateUser(ctx context.Context, p *users.CreateUserPayload) (res *users.CreateUserResult, err error) {
	res = &users.CreateUserResult{}
	s.logger.Print("users.createUser")
	return
}

// Update one User
func (s *userssrvc) UpdateUser(ctx context.Context, p *users.UpdateUserPayload) (res *users.UpdateUserResult, err error) {
	res = &users.UpdateUserResult{}
	s.logger.Print("users.updateUser")
	return
}

// Get one User
func (s *userssrvc) GetUser(ctx context.Context, p *users.GetUserPayload) (res *users.GetUserResult, err error) {
	res = &users.GetUserResult{}
	s.logger.Print("users.getUser")
	return
}
