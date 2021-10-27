package api

import (
	jwttoken "api_crud/gen/jwt_token"
	oauth "api_crud/gen/o_auth"
	"context"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"goa.design/goa/v3/security"
)

// oAuth service example implementation.
// The example methods log the requests and return zero values.
type oAuthsrvc struct {
	logger *log.Logger
	server *Server
}

// NewOAuth returns the oAuth service implementation.
func NewOAuth(logger *log.Logger, server *Server) oauth.Service {
	return &oAuthsrvc{logger, server}
}

var (
	ErrUnsupportedGrantType error = oauth.Unauthorized("Unsupported grant")
	ErrInvalidRequest       error = oauth.Unauthorized("Invalid request")
)

func (s *oAuthsrvc) errorResponse(msg string, err error) *oauth.UnknownError {
	return &oauth.UnknownError{
		Err:       err.Error(),
		ErrorCode: msg,
	}
}

// oAuth
func (s *oAuthsrvc) OAuth(ctx context.Context, p *oauth.OauthPayload) (res *oauth.OAuthResponse, err error) {
	if p.GrantType != "grant_type" {
		return nil, ErrUnsupportedGrantType
	}

	if p.ClientID == "" || p.ClientSecret == "" {
		return nil, ErrInvalidRequest
	}

	if p.ClientID != s.server.Config.Security.OAuthID || p.ClientSecret != s.server.Config.Security.OAuthSecret {
		return nil, ErrInvalidRequest
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"expires_in": time.Now().Add(time.Duration(time.Second * 2)).Unix(),
		"scopes":     []string{"api:read", "api:write"},
		"token_type": "Bearer",
	})

	t, err := accessToken.SignedString([]byte(s.server.Config.Security.Secret))
	if err != nil {
		return nil, s.errorResponse("ERROR_GENERATE_ACCESS_JWT", err)
	}

	res = &oauth.OAuthResponse{
		AccessToken: t,
		ExpiresIn:   time.Now().Add(time.Duration(time.Second * 2)).Unix(),
		TokenType:   "Bearer",
		Success:     true,
	}
	return res, nil
}

func (s *booksrvc) OAuth2Auth(ctx context.Context, token string, scheme *security.OAuth2Scheme) (context.Context, error) {
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
	if claims["expires_in"] == nil {
		return ctx, ErrInvalidTokenScopes
	}
	if claims["token_type"] != "Bearer" {
		return ctx, ErrInvalidToken
	}
	scopes, ok := claims["scopes"].([]interface{})
	if !ok {
		return ctx, ErrInvalidTokenScopes
	}
	scopesInToken := make([]string, len(scopes))
	for _, scp := range scopes {
		scopesInToken = append(scopesInToken, scp.(string))
	}
	if err := scheme.Validate(scopesInToken); err != nil {
		return ctx, jwttoken.InvalidScopes(err.Error())
	}

	// 3. add authInfo to context
	ctx = contextWithAuthInfo(ctx, authInfo{
		claims: claims,
	})
	return ctx, nil
}
