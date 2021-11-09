package api

import (
	users "api_crud/gen/users"
	db "api_crud/internal"
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
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
}

// JWTAuth implements the authorization logic for service "users" for the "jwt"
// security scheme.
func (s *userssrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	return s.server.CheckJWT(ctx, token, scheme)
}

// Get All users
func (s *userssrvc) GetAllusers(ctx context.Context, p *users.GetAllusersPayload) (res *users.GetAllusersResult, err error) {
	res = &users.GetAllusersResult{}
	s.logger.Print("users.getAllusers")
	return
}

// Delete one User by ID
func (s *userssrvc) DeleteUser(ctx context.Context, p *users.DeleteUserPayload) (res *users.DeleteUserResult, err error) {
	err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		if err := q.DeleteUserByID(ctx, uuid.MustParse(p.ID)); err != nil {
			return fmt.Errorf("DELETE_USER_BY_ID %v", err)
		}
		return nil
	})
	if err != nil {
		return nil, s.errorResponse("TX_DELETE_USER", err)
	}
	return &users.DeleteUserResult{
		Success: true,
	}, nil
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
	err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		u, err := q.GetUserByID(ctx, uuid.MustParse(p.ID))
		if err != nil {
			return fmt.Errorf("GET_USER_BY_ID %v", err)
		}
		res = &users.GetUserResult{
			User: &users.ResUser{
				ID:        u.ID.String(),
				Firstname: &u.Firstname,
				Lastname:  &u.Lastname,
				Email:     u.Email,
				Phone:     u.Phone.String,
				Birthday:  u.Birthday.String,
			},
			Success: true,
		}
		return nil
	})
	if err != nil {
		return nil, s.errorResponse("TX_GET_USER_BY_ID", err)
	}
	return res, nil
}
