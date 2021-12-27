package api

import (
	files "api_crud/gen/files"
	db "api_crud/internal"
	"api_crud/utils"
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
	result := &files.ImportFileResult{}
	// fmt.Println("content => ", p.Content)
	fmt.Println("name => ", p.Filename)
	fmt.Println("format => ", p.Format)
	err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		arg := db.CreateFileParams{
			Name: utils.NullS(p.Filename),
			Url:  utils.NullS(*p.URL),
			Mime: utils.NullS(*p.Mime),
		}
		newFile, err := q.CreateFile(ctx, arg)
		if err != nil {
			return fmt.Errorf("ERROR_CREATE_FILE %v", err)
		}
		result = &files.ImportFileResult{
			File: &files.ResFile{
				ID:   newFile.ID.String(),
				Name: newFile.Name.String,
				Mime: &newFile.Mime.String,
				Size: &newFile.Size.Int64,
				URL:  newFile.Url.String,
			},
			Success: true,
		}
		return nil
	})
	if err != nil {
		return nil, s.errorResponse("TX_CREATE_FILE", err)
	}
	return result, nil
}
