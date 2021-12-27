package utils

import (
	"log"
	"os"
	"path/filepath"
)

// Dir return full path
func Dir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Println(err)
	}
	return dir
}
