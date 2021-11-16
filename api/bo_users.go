package api

import (
	bousers "api_crud/gen/bo_users"
	db "api_crud/internal"
	"api_crud/utils"
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"goa.design/goa/v3/security"
)

// boUsers service example implementation.
// The example methods log the requests and return zero values.
type boUserssrvc struct {
	logger *log.Logger
	server *Server
}

// NewBoUsers returns the boUsers service implementation.
func NewBoUsers(logger *log.Logger, server *Server) bousers.Service {
	return &boUserssrvc{logger, server}
}

func (s *boUserssrvc) errorResponse(msg string, err error) *bousers.UnknownError {
	return &bousers.UnknownError{
		Err:       err.Error(),
		ErrorCode: msg,
	}
}

// OAuth2Auth implements the authorization logic for service "boUsers" for the
// "OAuth2" security scheme.
func (s *boUserssrvc) OAuth2Auth(ctx context.Context, token string, scheme *security.OAuth2Scheme) (context.Context, error) {
	return s.server.CheckAuth(ctx, token, scheme)
}

// JWTAuth implements the authorization logic for service "boUsers" for the
// "jwt" security scheme.
func (s *boUserssrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	return s.server.CheckJWT(ctx, token, scheme)
}

// Get All users
func (s *boUserssrvc) GetAllusers(ctx context.Context, p *bousers.GetAllusersPayload) (res *bousers.GetAllusersResult, err error) {
	err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		uS, err := q.GetAllUsers(ctx)
		if err != nil {
			return fmt.Errorf("ERROR_GET_ALL_USERS %v", err)
		}
		var allUsers []*bousers.ResBoUser
		for _, v := range uS {
			var firstname string = v.Firstname
			var lastname string = v.Lastname
			allUsers = append(allUsers, &bousers.ResBoUser{
				ID:        v.ID.String(),
				Firstname: &firstname,
				Lastname:  &lastname,
				Email:     v.Email,
				Phone:     v.Phone.String,
				Birthday:  v.Birthday.String,
			})
		}
		res = &bousers.GetAllusersResult{
			Users:   allUsers,
			Success: true,
		}
		return nil
	})
	if err != nil {
		return nil, s.errorResponse("TX_GET_ALL_USERS", err)
	}
	return res, nil
}

// Delete one User by ID
func (s *boUserssrvc) DeleteUser(ctx context.Context, p *bousers.DeleteUserPayload) (res *bousers.DeleteUserResult, err error) {
	err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		if err := q.DeleteUserByID(ctx, uuid.MustParse(p.ID)); err != nil {
			return fmt.Errorf("ERROR_DELETE_USER_BY_ID %v", err)
		}
		return nil
	})
	if err != nil {
		return nil, s.errorResponse("TX_DELETE_USER", err)
	}
	return &bousers.DeleteUserResult{Success: true}, nil
}

// Create one User
func (s *boUserssrvc) CreateUser(ctx context.Context, p *bousers.CreateUserPayload) (res *bousers.CreateUserResult, err error) {
	err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		arg := db.CreateUserParams{
			Firstname: p.User.Firstname,
			Lastname:  p.User.Lastname,
			Email:     p.User.Email,
			Phone:     utils.NullS(p.User.Phone),
			Birthday:  utils.NullS(p.User.Birthday),
		}
		createUser, err := q.CreateUser(ctx, arg)
		if err != nil {
			return fmt.Errorf("ERROR_CREATE_USER %v", err)
		}
		NewUsers, err := q.GetUserByID(ctx, createUser.ID)
		if err != nil {
			return fmt.Errorf("ERROR_GET_USER_BY_ID %v", err)
		}
		res = &bousers.CreateUserResult{
			User: &bousers.ResBoUser{
				Firstname: &NewUsers.Firstname,
				Lastname:  &NewUsers.Lastname,
				Email:     NewUsers.Email,
				Phone:     NewUsers.Phone.String,
				Birthday:  NewUsers.Birthday.String,
			},
			Success: true,
		}
		return nil
	})
	if err != nil {
		return nil, s.errorResponse("TX_CREATE_USER", err)
	}
	return res, nil
}

// Update one User
func (s *boUserssrvc) UpdateUser(ctx context.Context, p *bousers.UpdateUserPayload) (res *bousers.UpdateUserResult, err error) {
	err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		arg := db.UpdateUserParams{
			ID:        uuid.MustParse(p.ID),
			Firstname: p.User.Firstname,
			Lastname:  p.User.Lastname,
			Email:     p.User.Email,
			Phone:     utils.NullS(p.User.Phone),
			Birthday:  utils.NullS(p.User.Birthday),
		}
		if err := q.UpdateUser(ctx, arg); err != nil {
			return fmt.Errorf("ERROR_UPDATE_USER %v", err)
		}
		NewUsers, err := q.GetUserByID(ctx, uuid.MustParse(p.ID))
		if err != nil {
			return fmt.Errorf("ERROR_GET_USER_BY_ID %v", err)
		}
		res = &bousers.UpdateUserResult{
			Success: true,
			User: &bousers.ResBoUser{
				Firstname: &NewUsers.Firstname,
				Lastname:  &NewUsers.Lastname,
				Email:     NewUsers.Email,
				Phone:     NewUsers.Phone.String,
				Birthday:  NewUsers.Birthday.String,
			},
		}
		return nil
	})
	if err != nil {
		return nil, s.errorResponse("TX_UPDATE_USER", err)
	}
	return res, nil
}

// Get one User
func (s *boUserssrvc) GetUser(ctx context.Context, p *bousers.GetUserPayload) (res *bousers.GetUserResult, err error) {
	err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		u, err := q.GetUserByID(ctx, uuid.MustParse(p.ID))
		if err != nil {
			return fmt.Errorf("ERROR_GET_USER_BY_ID %v", err)
		}
		res = &bousers.GetUserResult{
			User: &bousers.ResBoUser{
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
