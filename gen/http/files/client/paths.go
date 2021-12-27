// Code generated by goa v3.5.2, DO NOT EDIT.
//
// HTTP request path constructors for the files service.
//
// Command:
// $ goa gen api_crud/design

package client

import (
	"fmt"
)

// ImportFileFilesPath returns the URL path to the files service importFile HTTP endpoint.
func ImportFileFilesPath() string {
	return "/v1/bo/files/add"
}

// DeleteFileFilesPath returns the URL path to the files service deleteFile HTTP endpoint.
func DeleteFileFilesPath(url_ string) string {
	return fmt.Sprintf("/v1/bo/files/remove/%v", url_)
}
