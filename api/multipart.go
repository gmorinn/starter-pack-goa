package api

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"os"
	"path/filepath"
	files "starter-pack-goa/gen/files"
	"starter-pack-goa/utils"
	"strconv"
	"strings"

	"github.com/disintegration/imageorient"
	"github.com/disintegration/imaging"
	"github.com/google/uuid"
)

func convertByteToImg(format string, file **os.File, img *image.Image) error {
	switch true {
	case strings.Contains(format, "png"):
		if err := png.Encode(*file, *img); err != nil {
			return err
		}
		break
	case strings.Contains(format, "jpeg"):
		if err := jpeg.Encode(*file, *img, &jpeg.Options{Quality: 40}); err != nil {
			return err
		}
		break
	case strings.Contains(format, "jpg"):
		if err := jpeg.Encode(*file, *img, &jpeg.Options{Quality: 40}); err != nil {
			return err
		}
		break
	default:
		return errors.New("Erreur format")
	}
	return nil
}

func crop(file string, w, h int) {
	f, err := os.Open(utils.Dir() + "/" + file)
	if err != nil {
		fmt.Printf("\nos.Open failed: %v\n", err)
	}

	img, _, err := imageorient.Decode(f)
	if err != nil {
		fmt.Printf("\nimageorient.Decode failed: %v\n", err)
	}

	centercropimg := imaging.Fill(img, w, h, imaging.Center, imaging.Lanczos)

	err = imaging.Save(centercropimg, utils.Dir()+"/"+file)

	if err != nil {
		fmt.Printf("Error crop\n")
	}
}

func parsePart(part *multipart.Part) (int, error) {
	byteParse, err := ioutil.ReadAll(part)
	if err != nil {
		return 0, fmt.Errorf("%v", err)
	}
	byteToInt, err := strconv.Atoi(string(byteParse))
	if err != nil {
		return 0, fmt.Errorf("%v", err)
	}
	return byteToInt, nil
}

// FilesImportFileDecoderFunc implements the multipart decoder for service
// "files" endpoint "importFile". The decoder must populate the argument p
// after encoding.
func FilesImportFileDecoderFunc(mr *multipart.Reader, p **files.ImportFilePayload) error {
	r := files.ImportFilePayload{}
	var height int = 0
	var width int = 0
	var fn string = ""
	for {
		if height > 0 && width > 0 && fn != "" {
			crop(fn, width, height)
			break
		}
		part, err := mr.NextPart()
		if err != nil {
			fmt.Printf("Error part => %v\n", err)
			break
		}

		_, params, err := mime.ParseMediaType(part.Header.Get("Content-Disposition"))
		if err != nil {
			// can't process this entry, it probably isn't an image
			fmt.Printf("Error get mime => %v\n", err)
			return err
		}
		if params["name"] == "h" {
			height, err = parsePart(part)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			continue
		} else if params["name"] == "w" {
			width, err = parsePart(part)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			continue
		}

		if params["name"] != "content" {
			continue
		}

		disposition, _, err := mime.ParseMediaType(part.Header.Get("Content-Type"))
		// the disposition can be, for example 'image/jpeg' or 'video/mp4'
		// I want to support only image files!
		if err != nil || !strings.HasPrefix(disposition, "image/") {
			// can't process this entry, it probably isn't an image
			fmt.Printf("Error format => %v\n", errors.New("Wrong format"))
			return errors.New("Wrong format")
		}
		if params["filename"] != "" {
			bytefile, err := ioutil.ReadAll(part)
			if err != nil {
				// can't process this entry, for some reason
				fmt.Fprintln(os.Stderr, err)
				fmt.Printf("Error read byte => %v\n", err)
				return err
			}

			// OPEN FILE
			upldir, uploadDirFull := utils.UploadDir()
			ext := filepath.Ext(params["filename"])
			un := fmt.Sprintf("%s%s", uuid.New(), ext)
			fn = fmt.Sprintf("%s/%s", upldir, un)
			dst, err := os.Create(fmt.Sprintf("%s/%s", uploadDirFull, un))
			if err != nil {
				fmt.Printf("Error create file => %v\n", err)
				return err
			}
			defer dst.Close()

			img, _, err := image.Decode(bytes.NewReader(bytefile))
			if err != nil {
				fmt.Printf("Error decode image => %v\n", err)
				return err
			}

			if err := convertByteToImg(disposition, &dst, &img); err != nil {
				fmt.Printf("Error convert format => %v\n", err)
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
	}

	return nil
}

// FilesImportFileEncoderFunc implements the multipart encoder for service
// "files" endpoint "importFile".
func FilesImportFileEncoderFunc(mw *multipart.Writer, p *files.ImportFilePayload) error {
	// Add multipart request encoder logic here
	return nil
}
