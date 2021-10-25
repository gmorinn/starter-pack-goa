package api

import (
	oauth "api_crud/gen/o_auth"
	"context"
	"log"

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

// oAuth
func (s *oAuthsrvc) OAuth(ctx context.Context, p *oauth.OauthPayload) (res *oauth.OAuthResponse, err error) {
	res = &oauth.OAuthResponse{}
	s.logger.Print("oAuth.oAuth")
	return
}

func (s *booksrvc) OAuth2Auth(ctx context.Context, token string, scheme *security.OAuth2Scheme) (context.Context, error) {
	return nil, nil
}
