package api

import (
	files "api_crud/gen/files"
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"os"
	"strings"
)

// FilesImportFileDecoderFunc implements the multipart decoder for service
// "files" endpoint "importFile". The decoder must populate the argument p
// after encoding.
func FilesImportFileDecoderFunc(mr *multipart.Reader, p **files.ImportFilePayload) error {
	r := files.ImportFilePayload{}

	part, err := mr.NextPart()
	if err == io.EOF {
		fmt.Printf("Error 1 => %v\n", err)
		return err
	}
	if err != nil {
		fmt.Printf("Error 2 => %v\n", err)
		return err
	}
	_, params, err := mime.ParseMediaType(part.Header.Get("Content-Disposition"))
	if err != nil {
		// can't process this entry, it probably isn't an image
		fmt.Printf("Error 3 => %v\n", err)
		return err
	}

	disposition, _, err := mime.ParseMediaType(part.Header.Get("Content-Type"))
	// the disposition can be, for example 'image/jpeg' or 'video/mp4'
	// I want to support only image files!
	if err != nil || !strings.HasPrefix(disposition, "image/") {
		// can't process this entry, it probably isn't an image
		fmt.Printf("Error 4 => %v\n", err)
		return err
	}
	if params["filename"] != "" {
		bytes, err := ioutil.ReadAll(part)
		if err != nil {
			// can't process this entry, for some reason
			fmt.Fprintln(os.Stderr, err)
			fmt.Printf("Error 5 => %v\n", err)
			return err
		}
		filename := params["filename"]
		imageUpload := files.ImportFilePayload{
			Format:   disposition,
			Content:  bytes,
			FileName: filename,
		}
		r = imageUpload
	}
	*p = &r
	return nil
}

// FilesImportFileEncoderFunc implements the multipart encoder for service
// "files" endpoint "importFile".
func FilesImportFileEncoderFunc(mw *multipart.Writer, p *files.ImportFilePayload) error {
	// Add multipart request encoder logic here
	return nil
}
