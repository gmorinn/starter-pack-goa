package api

import (
	files "api_crud/gen/files"
	"api_crud/utils"
	"errors"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/disintegration/imageorient"
	"github.com/disintegration/imaging"
)

func convertByteToImg(format string, file **os.File, img *image.Image) error {
	switch true {
	case strings.Contains(format, "png"):
		if err := png.Encode(*file, *img); err != nil {
			return err
		}
		break
	case strings.Contains(format, "jpeg"):
		if err := jpeg.Encode(*file, *img, &jpeg.Options{Quality: 90}); err != nil {
			return err
		}
		break
	case strings.Contains(format, "jpg"):
		if err := jpeg.Encode(*file, *img, &jpeg.Options{Quality: 90}); err != nil {
			return err
		}
		break
	case strings.Contains(format, "gif"):
		if err := gif.Encode(*file, *img, &gif.Options{NumColors: 256}); err != nil {
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
		fmt.Printf("\n************* Error crop ************\n")
	}
}

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
	for {
		part, err := mr.NextPart()
		if err != nil {
			fmt.Printf("Error 1 => %v\n", err)
			return err
		}

		_, params, err := mime.ParseMediaType(part.Header.Get("Content-Disposition"))
		if err != nil {
			// can't process this entry, it probably isn't an image
			fmt.Printf("Error 2 => %v\n", err)
			return err
		}

		if params["h"] != "" {
			byteheight, err := ioutil.ReadAll(part)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			byteToInt, err := strconv.Atoi(string(byteheight))
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			fmt.Println(byteToInt)
		}

		disposition, _, err := mime.ParseMediaType(part.Header.Get("Content-Type"))
		// the disposition can be, for example 'image/jpeg' or 'video/mp4'
		// I want to support only image files!
		if err != nil || !strings.HasPrefix(disposition, "image/") {
			// can't process this entry, it probably isn't an image
			fmt.Printf("Error 3 => %v\n", err)
			return err
		}
		if params["filename"] != "" {
			bytefile, err := ioutil.ReadAll(part)
			if err != nil {
				// can't process this entry, for some reason
				fmt.Fprintln(os.Stderr, err)
				fmt.Printf("Error 4 => %v\n", err)
				return err
			}

			// // OPEN FILE
			// upldir, uploadDirFull := uploadDir()
			// ext := filepath.Ext(params["filename"])
			// un := fmt.Sprintf("%s%s", uuid.New(), ext)
			// fn := fmt.Sprintf("%s/%s", upldir, un)
			// dst, err := os.Create(fmt.Sprintf("%s/%s", uploadDirFull, un))
			// if err != nil {
			// 	fmt.Printf("Error 5 => %v\n", err)
			// 	return err
			// }
			// defer dst.Close()

			// img, _, err := image.Decode(bytes.NewReader(bytefile))
			// if err != nil {
			// 	fmt.Printf("Error 6 => %v\n", err)
			// 	return err
			// }

			// if err := convertByteToImg(disposition, &dst, &img); err != nil {
			// 	fmt.Printf("Error 7 => %v\n", err)
			// 	return err
			// }

			mimefile := string(part.Header.Get("Content-Type")[0])
			filename := params["filename"]
			// url := "/" + fn
			imageUpload := files.ImportFilePayload{
				Format:   disposition,
				Content:  bytefile,
				Filename: filename,
				Mime:     &mimefile,
				// URL:      &url,
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
