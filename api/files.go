package api

import (
	files "api_crud/gen/files"
	"context"
	"fmt"
	"log"

	"goa.design/goa/v3/security"
)

// files service example implementation.
// The example methods log the requests and return zero values.
type filessrvc struct {
	logger *log.Logger
	server *Server
}

// NewFiles returns the files service implementation.
func NewFiles(logger *log.Logger, server *Server) files.Service {
	return &filessrvc{logger, server}
}

func (s *filessrvc) errorResponse(msg string, err error) *files.UnknownError {
	return &files.UnknownError{
		Err:       err.Error(),
		ErrorCode: msg,
	}
}

// OAuth2Auth implements the authorization logic for service "files" for the
// "OAuth2" security scheme.
func (s *filessrvc) OAuth2Auth(ctx context.Context, token string, scheme *security.OAuth2Scheme) (context.Context, error) {
	return s.server.CheckAuth(ctx, token, scheme)
}

// JWTAuth implements the authorization logic for service "files" for the "jwt"
// security scheme.
func (s *filessrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	return s.server.CheckJWT(ctx, token, scheme)
}

// Import file
func (s *filessrvc) ImportFile(ctx context.Context, p *files.ImportFilePayload) (res *files.ImportFileResult, err error) {
	result := &files.ImportFileResult{Success: true}
	fmt.Println("content => ", p.Content)
	fmt.Println("name => ", p.FileName)
	fmt.Println("format => ", p.Format)
	return result, nil
}
