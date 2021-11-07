package main

import (
	"io/fs"
	"io/ioutil"
	"log"
	"os"
)

func checkNameFile(api []fs.FileInfo, fileName string) bool {
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
		if checkNameFile(apiFolder, f.Name()) {
			e := os.Remove(f.Name())
			if e != nil {
				log.Fatal(e)
			}
		}
	}

	for _, v := range currentFolder {
		if v.Name() == "http.go" || v.Name() == "main.go" || len(v.Name()) < 3 {
			continue
		} else if len(v.Name()) > 3 && v.Name()[len(v.Name())-3:len(v.Name())] == ".go" {
			if !checkNameFile(apiFolder, v.Name()) {
				err := os.Rename(v.Name(), "./api/"+v.Name())
				if err != nil {
					log.Fatal(err)
				}
			}
		}

	}
}
