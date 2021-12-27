package api

import (
	fileapi "api_crud/gen/fileapi"
	"log"
)

// fileapi service example implementation.
// The example methods log the requests and return zero values.
type fileapisrvc struct {
	logger *log.Logger
}

// Give access to see files for client
func NewFileapi(logger *log.Logger) fileapi.Service {
	return &fileapisrvc{logger}
}
