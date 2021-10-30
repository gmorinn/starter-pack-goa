package api

import (
	"api_crud/gen/book"
	jwttoken "api_crud/gen/jwt_token"
	db "api_crud/internal"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"goa.design/goa/v3/security"
	"gopkg.in/oauth2.v3/utils/uuid"
)

type jwtTokensrvc struct {
	logger *log.Logger
	server *Server
}

func errorEmail() *jwttoken.EmailAlreadyExist {
	return &jwttoken.EmailAlreadyExist{
		Message: "EMAIL_ALREADY_EXIST",
	}
}

func NewJWTToken(logger *log.Logger, server *Server) jwttoken.Service {
	return &jwtTokensrvc{logger, server}
}

func (s *jwtTokensrvc) errorResponse(msg string, err error) *jwttoken.UnknownError {
	return &jwttoken.UnknownError{
		Err:       err.Error(),
		ErrorCode: msg,
	}
}

var (
	ErrInvalidToken       error = book.Unauthorized("invalid token")
	ErrInvalidTokenScopes error = jwttoken.InvalidScopes("invalid scopes in token")
)

func (s *booksrvc) JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error) {

	claims := make(jwt.MapClaims)

	// authorize request
	// 1. parse JWT token, token key is hardcoded to "secret" in this example
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		b := ([]byte(s.server.Config.Security.Secret))
		return b, nil
	})
	if err != nil {
		return ctx, ErrInvalidToken
	}

	// 2. validate provided "scopes" claim
	if claims["scopes"] == nil {
		return ctx, ErrInvalidTokenScopes
	}
	if claims["id"] == nil {
		return ctx, ErrInvalidTokenScopes
	}
	if claims["exp"] == nil {
		return ctx, ErrInvalidTokenScopes
	}
	scopes, ok := claims["scopes"].([]interface{})
	if !ok {
		return ctx, ErrInvalidTokenScopes
	}
	scopesInToken := make([]string, len(scopes))
	for _, scp := range scopes {
		scopesInToken = append(scopesInToken, scp.(string))
	}
	if err := schema.Validate(scopesInToken); err != nil {
		return ctx, jwttoken.InvalidScopes(err.Error())
	}

	// 3. add authInfo to context
	ctx = contextWithAuthInfo(ctx, authInfo{
		jwtToken: claims,
	})
	return ctx, nil
}

func (s *jwtTokensrvc) Signup(ctx context.Context, p *jwttoken.SignupPayload) (res *jwttoken.Sign, err error) {

	isExist, err := s.server.Store.ExistUserByEmail(ctx, p.Email)
	if err != nil {
		return nil, s.errorResponse("ERROR_GET_MAIL", err)
	}
	if isExist {
		return nil, errorEmail()
	}

	arg := db.SignupParams{
		Firstname: p.Firstname,
		Lastname:  p.Lastname,
		Email:     p.Email,
		Crypt:     p.Password,
	}
	user, err := s.server.Store.Signup(ctx, arg)
	if err != nil {
		return nil, s.errorResponse("ERROR_CREATE_USER", err)
	}

	t, r, expt, err := s.generateJwtToken(uuid.UUID(user.ID))
	if err != nil {
		return nil, s.errorResponse("ERROR_TOKEN", err)
	}

	if err := s.server.StoreRefresh(ctx, r, expt, user.ID); err != nil {
		return nil, s.errorResponse("ERROR_REFRESH_TOKEN", err)
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
		return nil, s.errorResponse("ERROR_LOGIN_USER", err)
	}

	t, r, expt, err := s.generateJwtToken(uuid.UUID(user.ID))
	if err != nil {
		return nil, s.errorResponse("ERROR_TOKEN", err)
	}

	if err := s.server.StoreRefresh(ctx, r, expt, user.ID); err != nil {
		return nil, s.errorResponse("ERROR_REFRESH_TOKEN", err)
	}

	response := jwttoken.Sign{
		AccessToken:  t,
		RefreshToken: r,
		Success:      true,
	}
	return &response, nil
}

func (s *jwtTokensrvc) Refresh(ctx context.Context, p *jwttoken.RefreshPayload) (res *jwttoken.Sign, err error) {
	token, err := jwt.Parse(p.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		b := ([]byte(s.server.Config.Security.Secret))
		return b, nil
	})

	if err != nil {
		return nil, s.errorResponse("TOKEN_ERROR", err)
	}

	if !token.Valid {
		return nil, s.errorResponse("TOKEN_IS_NOT_VALID", err)
	}

	refresh, err := s.server.Store.GetRefreshToken(ctx, p.RefreshToken)
	if err != nil {
		return nil, s.errorResponse("FIND_REFRESH_TOKEN", err)
	}

	t, r, expt, err := s.generateJwtToken(uuid.UUID(refresh.UserID))
	if err != nil {
		return nil, s.errorResponse("ERROR_TOKEN", err)
	}

	if err := s.server.StoreRefresh(ctx, r, expt, refresh.UserID); err != nil {
		return nil, s.errorResponse("ERROR_REFRESH_TOKEN", err)
	}

	response := jwttoken.Sign{
		AccessToken:  t,
		RefreshToken: r,
		Success:      true,
	}
	return &response, nil
}

func (s *jwtTokensrvc) generateJwtToken(ID uuid.UUID) (string, string, time.Time, error) {
	// Generate access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":     ID.String(),
		"exp":    time.Now().Add(time.Duration((time.Hour * 24) * time.Duration(s.server.Config.Security.AccessTokenDuration))).Unix(),
		"scopes": []string{"api:read", "api:write"},
	})
	t, err := accessToken.SignedString([]byte(s.server.Config.Security.Secret))
	if err != nil {
		return "", "", time.Now(), fmt.Errorf("ERROR_GENERATE_ACCESS_JWT %v", err)
	}

	expt := time.Now().Add(time.Duration((time.Hour * 24) * time.Duration(s.server.Config.Security.RefreshTokenDuration)))
	exp := expt.Unix()

	// Generate refresh token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":     ID.String(),
		"exp":    exp,
		"scopes": []string{"api:read", "api:write"},
	})
	r, err := refreshToken.SignedString([]byte(s.server.Config.Security.Secret))
	if err != nil {
		return "", "", time.Now(), fmt.Errorf("ERROR_GENERATE_REFRESH_JWT %v", err)
	}

	return t, r, expt, nil
}