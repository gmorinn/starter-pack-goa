package api

import (
	files "api_crud/gen/files"
	"api_crud/utils"
	"bytes"
	"fmt"
	"image"
	"image/png"
	"io"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
)

func uploadDir() (string, string) {
	t := time.Now()
	upldir := fmt.Sprintf("public/uploads/%s/%s", t.Format("2006"), t.Format("01"))
	os.MkdirAll(utils.Dir()+"/"+upldir, os.ModePerm)
	return upldir, utils.Dir() + "/" + upldir
}

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
		bytefile, err := ioutil.ReadAll(part)
		if err != nil {
			// can't process this entry, for some reason
			fmt.Fprintln(os.Stderr, err)
			fmt.Printf("Error 5 => %v\n", err)
			return err
		}

		// OPEN FILE
		upldir, uploadDirFull := uploadDir()
		ext := filepath.Ext(params["filename"])
		un := fmt.Sprintf("%s%s", uuid.New(), ext)
		fn := fmt.Sprintf("%s/%s", upldir, un)
		dst, err := os.Create(fmt.Sprintf("%s/%s", uploadDirFull, un))
		if err != nil {
			fmt.Printf("Error 6 => %v\n", err)
			return err
		}
		defer dst.Close()

		img, _, err := image.Decode(bytes.NewReader(bytefile))
		if err != nil {
			fmt.Printf("Error 7 => %v\n", err)
			return err
		}

		err = png.Encode(dst, img)
		if err != nil {
			fmt.Printf("Error 8 => %v\n", err)
			return err
		}

		mimefile := string(part.Header.Get("Content-Type")[0])
		filename := params["filename"]
		url := "/" + fn
		imageUpload := files.ImportFilePayload{
			Format:   disposition,
			Content:  bytefile,
			Filename: filename,
			Mime:     &mimefile,
			URL:      &url,
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
