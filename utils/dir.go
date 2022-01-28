package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

// Dir return full path
func Dir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Println(err)
	}
	return dir
}

func UploadDir() (string, string) {
	t := time.Now()
	upldir := fmt.Sprintf("public/uploads/%s/%s", t.Format("2006"), t.Format("01"))
	os.MkdirAll(Dir()+"/"+upldir, os.ModePerm)
	return upldir, Dir() + "/" + upldir
}
