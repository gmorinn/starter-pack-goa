package api

import (
	files "api_crud/gen/files"
	"bytes"
	"fmt"
	"image"
	"io"
	"io/ioutil"
	"mime/multipart"
)

// FilesImportFileDecoderFunc implements the multipart decoder for service
// "files" endpoint "importFile". The decoder must populate the argument p
// after encoding.
func FilesImportFileDecoderFunc(mr *multipart.Reader, p **files.ImportFilePayload) error {
	var r files.ImportFilePayload
	for {
		part, err := mr.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("error 1 => ", err)
			return err
		}
		if part.FileName() != "" {
			slurp, err := ioutil.ReadAll(part)
			if err != nil {
				fmt.Println("error 2 => ", err)
				return err
			}

			r.FileName = part.FileName()
			r.Content = slurp

			// guess mime type
			slurpReader := bytes.NewReader(slurp)
			_, ft, err := image.Decode(slurpReader)
			if err != nil {
				fmt.Println("error 3 => ", err)
				return err
			}
			r.Format = ft
			break
		}
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
