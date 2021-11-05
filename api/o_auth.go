package api

import (
	oauth "api_crud/gen/o_auth"
	"context"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
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
	if p.GrantType != "client_credentials" {
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
