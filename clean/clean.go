package main

import (
	"io/fs"
	"io/ioutil"
	"log"
	"os"
)

func CheckNameFile(api []fs.FileInfo, fileName string) bool {
	for _, v := range api {
		if v.Name() == fileName {
			return true
		}
	}
	return false
}

func main() {
	currentFolder, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	apiFolder, err := ioutil.ReadDir("./api")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range currentFolder {
		if CheckNameFile(apiFolder, f.Name()) {
			e := os.Remove(f.Name())
			if e != nil {
				log.Fatal(e)
			}
		}
	}
}
