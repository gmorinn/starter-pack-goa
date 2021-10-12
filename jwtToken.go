package basic

import (
	"api_crud/api"
	jwttoken "api_crud/gen/jwt_token"
	"api_crud/internal/db"
	"context"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtTokensrvc struct {
	logger *log.Logger
	server *api.Server
}

func ErrorEmail() *jwttoken.EmailAlreadyExist {
	return &jwttoken.EmailAlreadyExist{
		Message: "EMAIL_ALREADY_EXIST",
	}
}

func NewJwtToken(logger *log.Logger, server *api.Server) jwttoken.Service {
	return &jwtTokensrvc{logger, server}
}

func (s *jwtTokensrvc) Signup(ctx context.Context, p *jwttoken.SignupPayload) (res *jwttoken.Sign, err error) {

	isExist, err := s.server.Store.ExistUserByEmail(ctx, p.Email)
	if err != nil {
		return nil, ErrorResponse("ERROR_GET_MAIL", err)
	}
	if isExist {
		return nil, ErrorEmail()
	}

	arg := db.SignupParams{
		Firstname: p.Firstname,
		Lastname:  p.Lastname,
		Email:     p.Email,
		Crypt:     p.Password,
	}
	user, err := s.server.Store.Signup(ctx, arg)
	if err != nil {
		return nil, ErrorResponse("ERROR_CREATE_USER", err)
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":     user.ID.String(),
		"exp":    time.Now().Add(time.Duration((time.Hour * 24) * time.Duration(s.server.Config.Security.AccessTokenDuration))).Unix(),
		"scopes": []string{"api:read", "api:write"},
	})
	t, err := accessToken.SignedString(s.server.Config.Security.Secret)
	if err != nil {
		return nil, ErrorResponse("ERROR_GENERATE_ACCESS_JWT", err)
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":     user.ID.String(),
		"exp":    time.Now().Add(time.Duration((time.Hour * 24) * time.Duration(s.server.Config.Security.RefreshTokenDuration))).Unix(),
		"scopes": []string{"api:read", "api:write"},
	})
	r, err := refreshToken.SignedString(s.server.Config.Security.Secret)
	if err != nil {
		return nil, ErrorResponse("ERROR_GENERATE_REFRESH_JWT", err)
	}

	response := jwttoken.Sign{
		AccessToken:  t,
		RefreshToken: r,
		Success:      true,
	}
	return &response, nil
}

func (s *jwtTokensrvc) Signin(ctx context.Context, p *jwttoken.SigninPayload) (res *jwttoken.Sign, err error) {
	// Request Login
	arg := db.LoginUserParams{
		Email: p.Email,
		Crypt: p.Password,
	}
	user, err := s.server.Store.LoginUser(ctx, arg)
	if err != nil {
		return nil, ErrorResponse("ERROR_LOGIN_USER", err)
	}

	// Generate access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":     user.ID.String(),
		"exp":    time.Now().Add(time.Duration((time.Hour * 24) * time.Duration(s.server.Config.Security.AccessTokenDuration))).Unix(),
		"scopes": []string{"api:read", "api:write"},
	})
	t, err := accessToken.SignedString([]byte(s.server.Config.Security.Secret))
	if err != nil {
		return nil, ErrorResponse("ERROR_GENERATE_ACCESS_JWT", err)
	}

	// Generate refresh token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":     user.ID.String(),
		"exp":    time.Now().Add(time.Duration((time.Hour * 24) * time.Duration(s.server.Config.Security.RefreshTokenDuration))).Unix(),
		"scopes": []string{"api:read", "api:write"},
	})
	r, err := refreshToken.SignedString([]byte(s.server.Config.Security.Secret))
	if err != nil {
		return nil, ErrorResponse("ERROR_GENERATE_REFRESH_JWT", err)
	}

	response := jwttoken.Sign{
		AccessToken:  t,
		RefreshToken: r,
		Success:      true,
	}
	return &response, nil
}
