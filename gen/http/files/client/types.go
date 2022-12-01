// Code generated by goa v3.10.2, DO NOT EDIT.
//
// files HTTP client types
//
// Command:
// $ goa gen starter-pack-goa/design

package client

import (
	files "starter-pack-goa/gen/files"
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// ImportFileRequestBody is the type of the "files" service "importFile"
// endpoint HTTP request body.
type ImportFileRequestBody struct {
	// Files to import
	Files []*PayloadFileRequestBody `form:"files" json:"files" xml:"files"`
}

// DeleteFileRequestBody is the type of the "files" service "deleteFile"
// endpoint HTTP request body.
type DeleteFileRequestBody struct {
	URL string `form:"url" json:"url" xml:"url"`
}

// ImportFileResponseBody is the type of the "files" service "importFile"
// endpoint HTTP response body.
type ImportFileResponseBody struct {
	File    []*ResFileResponseBody `form:"file,omitempty" json:"file,omitempty" xml:"file,omitempty"`
	Success *bool                  `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
}

// DeleteFileResponseBody is the type of the "files" service "deleteFile"
// endpoint HTTP response body.
type DeleteFileResponseBody struct {
	Success *bool `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
}

// ImportFileUnknownErrorResponseBody is the type of the "files" service
// "importFile" endpoint HTTP response body for the "unknown_error" error.
type ImportFileUnknownErrorResponseBody struct {
	Err       *string `form:"err,omitempty" json:"err,omitempty" xml:"err,omitempty"`
	ErrorCode *string `form:"error_code,omitempty" json:"error_code,omitempty" xml:"error_code,omitempty"`
	Success   *bool   `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
}

// DeleteFileUnknownErrorResponseBody is the type of the "files" service
// "deleteFile" endpoint HTTP response body for the "unknown_error" error.
type DeleteFileUnknownErrorResponseBody struct {
	Err       *string `form:"err,omitempty" json:"err,omitempty" xml:"err,omitempty"`
	ErrorCode *string `form:"error_code,omitempty" json:"error_code,omitempty" xml:"error_code,omitempty"`
	Success   *bool   `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
}

// PayloadFileRequestBody is used to define fields on request body types.
type PayloadFileRequestBody struct {
	// uploaded file name
	Filename string `form:"filename" json:"filename" xml:"filename"`
	// url file
	URL string `form:"url" json:"url" xml:"url"`
	// width of image if you crop
	W *int64 `form:"w,omitempty" json:"w,omitempty" xml:"w,omitempty"`
	// height of image if you crop
	H *int64 `form:"h,omitempty" json:"h,omitempty" xml:"h,omitempty"`
	// content of image
	Content []byte `form:"content" json:"content" xml:"content"`
	// size of image
	Size int64 `form:"size" json:"size" xml:"size"`
	// uploaded file format
	Format string `form:"format" json:"format" xml:"format"`
}

// ResFileResponseBody is used to define fields on response body types.
type ResFileResponseBody struct {
	ID   *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	URL  *string `form:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
	Mime *string `form:"mime,omitempty" json:"mime,omitempty" xml:"mime,omitempty"`
	Size *int64  `form:"size,omitempty" json:"size,omitempty" xml:"size,omitempty"`
}

// NewImportFileRequestBody builds the HTTP request body from the payload of
// the "importFile" endpoint of the "files" service.
func NewImportFileRequestBody(p *files.ImportFilePayload) *ImportFileRequestBody {
	body := &ImportFileRequestBody{}
	if p.Files != nil {
		body.Files = make([]*PayloadFileRequestBody, len(p.Files))
		for i, val := range p.Files {
			body.Files[i] = marshalFilesPayloadFileToPayloadFileRequestBody(val)
		}
	}
	return body
}

// NewDeleteFileRequestBody builds the HTTP request body from the payload of
// the "deleteFile" endpoint of the "files" service.
func NewDeleteFileRequestBody(p *files.DeleteFilePayload) *DeleteFileRequestBody {
	body := &DeleteFileRequestBody{
		URL: p.URL,
	}
	return body
}

// NewImportFileResultCreated builds a "files" service "importFile" endpoint
// result from a HTTP "Created" response.
func NewImportFileResultCreated(body *ImportFileResponseBody) *files.ImportFileResult {
	v := &files.ImportFileResult{
		Success: *body.Success,
	}
	v.File = make([]*files.ResFile, len(body.File))
	for i, val := range body.File {
		v.File[i] = unmarshalResFileResponseBodyToFilesResFile(val)
	}

	return v
}

// NewImportFileUnknownError builds a files service importFile endpoint
// unknown_error error.
func NewImportFileUnknownError(body *ImportFileUnknownErrorResponseBody) *files.UnknownError {
	v := &files.UnknownError{
		Err:       *body.Err,
		ErrorCode: *body.ErrorCode,
		Success:   *body.Success,
	}

	return v
}

// NewDeleteFileResultOK builds a "files" service "deleteFile" endpoint result
// from a HTTP "OK" response.
func NewDeleteFileResultOK(body *DeleteFileResponseBody) *files.DeleteFileResult {
	v := &files.DeleteFileResult{
		Success: *body.Success,
	}

	return v
}

// NewDeleteFileUnknownError builds a files service deleteFile endpoint
// unknown_error error.
func NewDeleteFileUnknownError(body *DeleteFileUnknownErrorResponseBody) *files.UnknownError {
	v := &files.UnknownError{
		Err:       *body.Err,
		ErrorCode: *body.ErrorCode,
		Success:   *body.Success,
	}

	return v
}

// ValidateImportFileResponseBody runs the validations defined on
// ImportFileResponseBody
func ValidateImportFileResponseBody(body *ImportFileResponseBody) (err error) {
	if body.Success == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("success", "body"))
	}
	if body.File == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("file", "body"))
	}
	for _, e := range body.File {
		if e != nil {
			if err2 := ValidateResFileResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateDeleteFileResponseBody runs the validations defined on
// DeleteFileResponseBody
func ValidateDeleteFileResponseBody(body *DeleteFileResponseBody) (err error) {
	if body.Success == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("success", "body"))
	}
	return
}

// ValidateImportFileUnknownErrorResponseBody runs the validations defined on
// importFile_unknown_error_response_body
func ValidateImportFileUnknownErrorResponseBody(body *ImportFileUnknownErrorResponseBody) (err error) {
	if body.Err == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("err", "body"))
	}
	if body.Success == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("success", "body"))
	}
	if body.ErrorCode == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("error_code", "body"))
	}
	return
}

// ValidateDeleteFileUnknownErrorResponseBody runs the validations defined on
// deleteFile_unknown_error_response_body
func ValidateDeleteFileUnknownErrorResponseBody(body *DeleteFileUnknownErrorResponseBody) (err error) {
	if body.Err == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("err", "body"))
	}
	if body.Success == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("success", "body"))
	}
	if body.ErrorCode == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("error_code", "body"))
	}
	return
}

// ValidatePayloadFileRequestBody runs the validations defined on
// payloadFileRequestBody
func ValidatePayloadFileRequestBody(body *PayloadFileRequestBody) (err error) {
	if body.Content == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("content", "body"))
	}
	if utf8.RuneCountInString(body.Filename) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("body.filename", body.Filename, utf8.RuneCountInString(body.Filename), 2, true))
	}
	if !(body.Format == "image/jpeg" || body.Format == "image/png" || body.Format == "image/jpg") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.format", body.Format, []interface{}{"image/jpeg", "image/png", "image/jpg"}))
	}
	return
}

// ValidateResFileResponseBody runs the validations defined on
// resFileResponseBody
func ValidateResFileResponseBody(body *ResFileResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.URL == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("url", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	return
}
