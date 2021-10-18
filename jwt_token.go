package basic

import (
	"api_crud/api"
	jwttoken "api_crud/gen/jwt_token"
	"api_crud/internal/db"
	"context"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"goa.design/goa/v3/security"
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

func NewJWTToken(logger *log.Logger, server *api.Server) jwttoken.Service {
	return &jwtTokensrvc{logger, server}
}

func (s *jwtTokensrvc) ErrorResponse(msg string, err error) *jwttoken.UnknownError {
	return &jwttoken.UnknownError{
		Err:       err.Error(),
		ErrorCode: msg,
	}
}

var (
	ErrInvalidToken       error = jwttoken.Unauthorized("invalid token")
	ErrInvalidTokenScopes error = jwttoken.InvalidScopes("invalid scopes in token")
)

func (s *crudsrvc) JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error) {

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
		claims: claims,
	})
	return ctx, nil
}

func (s *jwtTokensrvc) Signup(ctx context.Context, p *jwttoken.SignupPayload) (res *jwttoken.Sign, err error) {

	isExist, err := s.server.Store.ExistUserByEmail(ctx, p.Email)
	if err != nil {
		return nil, s.ErrorResponse("ERROR_GET_MAIL", err)
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
		return nil, s.ErrorResponse("ERROR_CREATE_USER", err)
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":     user.ID.String(),
		"exp":    time.Now().Add(time.Duration((time.Hour * 24) * time.Duration(s.server.Config.Security.AccessTokenDuration))).Unix(),
		"scopes": []string{"api:read", "api:write"},
	})
	t, err := accessToken.SignedString(s.server.Config.Security.Secret)
	if err != nil {
		return nil, s.ErrorResponse("ERROR_GENERATE_ACCESS_JWT", err)
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":     user.ID.String(),
		"exp":    time.Now().Add(time.Duration((time.Hour * 24) * time.Duration(s.server.Config.Security.RefreshTokenDuration))).Unix(),
		"scopes": []string{"api:read", "api:write"},
	})
	r, err := refreshToken.SignedString(s.server.Config.Security.Secret)
	if err != nil {
		return nil, s.ErrorResponse("ERROR_GENERATE_REFRESH_JWT", err)
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
		return nil, s.ErrorResponse("ERROR_LOGIN_USER", err)
	}
	// Generate access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":     user.ID.String(),
		"exp":    time.Now().Add(time.Duration((time.Hour * 24) * time.Duration(s.server.Config.Security.AccessTokenDuration))).Unix(),
		"scopes": []string{"api:read", "api:write"},
	})
	t, err := accessToken.SignedString([]byte(s.server.Config.Security.Secret))
	if err != nil {
		return nil, s.ErrorResponse("ERROR_GENERATE_ACCESS_JWT", err)
	}

	// Generate refresh token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":     user.ID.String(),
		"exp":    time.Now().Add(time.Duration((time.Hour * 24) * time.Duration(s.server.Config.Security.RefreshTokenDuration))).Unix(),
		"scopes": []string{"api:read", "api:write"},
	})
	r, err := refreshToken.SignedString([]byte(s.server.Config.Security.Secret))
	if err != nil {
		return nil, s.ErrorResponse("ERROR_GENERATE_REFRESH_JWT", err)
	}

	response := jwttoken.Sign{
		AccessToken:  t,
		RefreshToken: r,
		Success:      true,
	}
	return &response, nil
}

// func NewOAuth2ClientBasicAuthMiddleware() goa.Middleware {
// 	return func(h goa.Handler) goa.Handler {
// 		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
// 			// Retrieve basic auth info, TBD these are urlencoded as per the spec...
// 			clientID, clientSecret, ok := req.BasicAuth()
// 			if !ok {
// 				// return clientID, clientSecret, re
// 			}

// 		}
// 	}
// }