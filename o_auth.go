package basic

import (
	"api_crud/api"
	oauth "api_crud/gen/o_auth"
	"context"
	"log"
)

// oAuth service example implementation.
// The example methods log the requests and return zero values.
type oAuthsrvc struct {
	logger *log.Logger
	server *api.Server
}

// NewOAuth returns the oAuth service implementation.
func NewOAuth(logger *log.Logger, server *api.Server) oauth.Service {
	return &oAuthsrvc{logger, server}
}

// oAuth
func (s *oAuthsrvc) OAuth(ctx context.Context, p *oauth.OauthPayload) (res *oauth.OAuthResponse, err error) {
	res = &oauth.OAuthResponse{}
	s.logger.Print("oAuth.oAuth")
	return
}
