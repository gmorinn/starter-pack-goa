package api

import (
	"context"
	"fmt"
	"log"
	auth "starter-pack-goa/gen/auth"
	db "starter-pack-goa/internal"
	"starter-pack-goa/utils"

	"goa.design/goa/v3/security"
)

// auth service example implementation.
// The example methods log the requests and return zero values.
type authsrvc struct {
	logger *log.Logger
	server *Server
}

// NewAuth returns the auth service implementation.
func NewAuth(logger *log.Logger, server *Server) auth.Service {
	return &authsrvc{logger, server}
}

func (s *authsrvc) errorResponse(msg string, err error) *auth.UnknownError {
	return &auth.UnknownError{
		Err:       err.Error(),
		ErrorCode: msg,
	}
}

// OAuth2Auth implements the authorization logic for service "auth" for the
// "OAuth2" security scheme.
func (s *authsrvc) OAuth2Auth(ctx context.Context, token string, scheme *security.OAuth2Scheme) (context.Context, error) {
	return s.server.CheckAuth(ctx, token, scheme)
}

// Check if email exist in database
func (s *authsrvc) EmailExist(ctx context.Context, p *auth.EmailExistPayload) (res *auth.EmailExistResult, err error) {
	if p == nil {
		return nil, s.errorResponse("PAYLOAD_NULL", ErrNullPayload)
	}
	var isExist bool = false
	err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		isExist, err = q.CheckEmailExist(ctx, p.Email)
		if err != nil {
			return fmt.Errorf("error check email: %v", err)
		}
		return nil
	})
	if err != nil {
		return nil, s.errorResponse("TX_EMAIL_EXIST", err)
	}
	return &auth.EmailExistResult{
		Success: true,
		Exist:   isExist,
	}, nil
}

// Check if email exist in database and send code in email
func (s *authsrvc) SendConfirmation(ctx context.Context, p *auth.SendConfirmationPayload) (res *auth.SendConfirmationResult, err error) {
	if p == nil {
		return nil, s.errorResponse("PAYLOAD_NULL", ErrNullPayload)
	}
	var isExist bool = false
	err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		isExist, err = q.CheckEmailExist(ctx, p.Email)
		if err != nil {
			return fmt.Errorf("error check email: %v", err)
		}
		if !isExist {
			return nil
		}
		user, err := q.FindUserByEmail(ctx, p.Email)
		if err != nil {
			return fmt.Errorf("error to find the email: %v", err)
		}
		if err := s.server.sendEmail(ctx, user.ID.String(), user.Email); err != nil {
			return fmt.Errorf("error for send the email: %v", err)
		}
		return nil
	})
	if err != nil {
		return nil, s.errorResponse("TX_UPDATE_EMAIL_EXIST", err)
	}
	return &auth.SendConfirmationResult{
		Success: true,
		Exist:   isExist,
	}, nil
}

// Reset forget password of the user with the correct confirm code
func (s *authsrvc) ResetPassword(ctx context.Context, p *auth.ResetPasswordPayload) (res *auth.ResetPasswordResult, err error) {
	if p == nil {
		return nil, s.errorResponse("PAYLOAD_NULL", ErrNullPayload)
	}
	err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		if p.Password != p.ConfirmPassword {
			return fmt.Errorf("not the same password")
		}
		arg := db.ExistUserByEmailAndConfirmCodeParams{
			Email:               p.Email,
			PasswordConfirmCode: utils.NullS(p.Code),
		}
		isExist, err := q.ExistUserByEmailAndConfirmCode(ctx, arg)
		if err != nil {
			return fmt.Errorf("error to search the email and code: %v", err)
		}
		if !isExist {
			return fmt.Errorf("CODE_OR_EMAIL_NOT_EXIST")
		}
		args := db.UpdatePasswordUserWithconfirmCodeParams{
			Email:               p.Email,
			PasswordConfirmCode: utils.NullS(p.Code),
			Crypt:               p.Password,
		}
		if err := q.UpdatePasswordUserWithconfirmCode(ctx, args); err != nil {
			return fmt.Errorf("error to set the code and password: %v", err)
		}
		return nil
	})

	if err != nil {
		return nil, s.errorResponse("TX_RESET_PASSWORD", err)
	}

	return &auth.ResetPasswordResult{
		Success: true,
	}, nil
}
