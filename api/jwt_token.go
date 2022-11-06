package api

import (
	"context"
	"fmt"
	"log"
	jwttoken "starter-pack-goa/gen/jwt_token"
	db "starter-pack-goa/internal"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"goa.design/goa/v3/security"
)

type jwtTokensrvc struct {
	logger *log.Logger
	server *Server
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

func (s *jwtTokensrvc) OAuth2Auth(ctx context.Context, token string, scheme *security.OAuth2Scheme) (context.Context, error) {
	return s.server.CheckAuth(ctx, token, scheme)
}

func (s *jwtTokensrvc) Signup(ctx context.Context, p *jwttoken.SignupPayload) (res *jwttoken.Sign, err error) {
	if p == nil {
		return nil, ErrNullPayload
	}
	if p.Password != p.ConfirmPassword {
		return nil, ErrInvalidPassword
	}
	isExist, err := s.server.Store.CheckEmailExist(ctx, p.Email)
	if err != nil {
		return nil, s.errorResponse("ERROR_GET_MAIL", err)
	}
	if isExist {
		return nil, s.errorResponse("EMAIL_ALREADY_EXIST", ErrEmailExist)
	}
	arg := db.SignupParams{
		Email: p.Email,
		Crypt: p.Password,
	}
	user, err := s.server.Store.Signup(ctx, arg)
	if err != nil {
		return nil, s.errorResponse("ERROR_CREATE_USER", err)
	}
	t, r, expt, err := s.generateJwtToken(user.ID, string(user.Role))
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
	if p == nil {
		return nil, ErrNullPayload
	}
	arg := db.LoginUserParams{
		Email: p.Email,
		Crypt: p.Password,
	}
	user, err := s.server.Store.LoginUser(ctx, arg)
	if err != nil {
		return nil, s.errorResponse("ERROR_LOGIN_USER", err)
	}
	t, r, expt, err := s.generateJwtToken(user.ID, string(user.Role))
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
	if p == nil {
		return nil, ErrNullPayload
	}
	token, err := jwt.Parse(p.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		b := ([]byte(s.server.Config.Security.Secret))
		return b, nil
	})
	if err != nil {
		return nil, s.errorResponse("TOKEN_ERROR", err)
	}
	if !token.Valid {
		return nil, s.errorResponse("TOKEN_IS_NOT_VALID", ErrInvalidToken)
	}
	refresh, err := s.server.Store.GetRefreshToken(ctx, p.RefreshToken)
	if err != nil {
		return nil, s.errorResponse("FIND_REFRESH_TOKEN", err)
	}
	t, r, expt, err := s.generateJwtToken(refresh.UserID, string(refresh.UserRole.Role))
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

func (s *jwtTokensrvc) generateJwtToken(ID uuid.UUID, role string) (string, string, time.Time, error) {
	// Generate access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":     ID.String(),
		"role":   role,
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
		"role":   role,
		"exp":    exp,
		"scopes": []string{"api:read", "api:write"},
	})
	r, err := refreshToken.SignedString([]byte(s.server.Config.Security.Secret))
	if err != nil {
		return "", "", time.Now(), fmt.Errorf("ERROR_GENERATE_REFRESH_JWT %v", err)
	}
	return t, r, expt, nil
}

func (server *Server) CheckJWT(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error) {

	if schema == nil {
		return ctx, ErrNullPayload
	}
	claims := make(jwt.MapClaims)
	// authorize request
	// 1. parse JWT token, token key is hardcoded to "secret" in this example
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		b := ([]byte(server.Config.Security.Secret))
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
		return ctx, ErrInvalidTokenScopes
	}

	// 3. add authInfo to context
	ctx = contextWithAuthInfo(ctx, authInfo{
		jwtToken: claims,
	})
	return ctx, nil
}
