package api

import (
	"context"
	"fmt"
	"log"
	"os"
	files "starter-pack-goa/gen/files"
	db "starter-pack-goa/internal"
	"starter-pack-goa/utils"

	"goa.design/goa/v3/security"
)

// ******************************************************************** ///

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
	err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		for _, v := range p.Files {
			arg := db.CreateFileParams{
				Name: utils.NullS(v.Filename),
				Url:  utils.NullS(v.URL),
				Mime: utils.NullS(v.Format),
				Size: utils.NullI64(int64(v.Size)),
			}
			newFile, err := q.CreateFile(ctx, arg)
			if err != nil {
				return fmt.Errorf("ERROR_CREATE_FILE %v", err)
			}
			result.File = append(result.File, &files.ResFile{
				ID:   newFile.ID.String(),
				Name: newFile.Name.String,
				Mime: &newFile.Mime.String,
				Size: &newFile.Size.Int64,
				URL:  newFile.Url.String,
			})
		}
		return nil
	})
	if err != nil {
		return nil, s.errorResponse("TX_CREATE_FILE", err)
	}
	result.Success = true
	return result, nil
}

// Delete file
func (s *filessrvc) DeleteFile(ctx context.Context, p *files.DeleteFilePayload) (res *files.DeleteFileResult, err error) {
	err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		file, err := q.GetFileByURL(ctx, utils.NullS(p.URL))
		if err != nil {
			return fmt.Errorf("ERROR_GET_FILE_BY_URL %v", err)
		}
		if err := q.DeleteFile(ctx, utils.NullS(file.Url.String)); err != nil {
			return fmt.Errorf("ERROR_DELETE_FILE_BY_ID %v", err)
		}

		if s.server.Config.Mode == "dev" {
			if err = os.Remove("bin" + p.URL); err != nil {
				return fmt.Errorf("ERROR_REMOVE_FILE_IN_FOLDER %v", err)
			}
		} else {
			if err = os.Remove("/go/bin" + p.URL); err != nil {
				return fmt.Errorf("ERROR_REMOVE_FILE_IN_FOLDER %v", err)
			}
		}
		return nil
	})
	if err != nil {
		return nil, s.errorResponse("TX_DELETE_FILE", err)
	}
	return &files.DeleteFileResult{Success: true}, nil
}
