package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
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
	//OPEN CURRENT FOLDER
	currentFolder, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
		os.Exit(84)
	}
	// OPEN API FOLDER
	apiFolder, err := ioutil.ReadDir("./api")
	if err != nil {
		log.Fatal(err)
		os.Exit(84)
	}
	// CHECK IF NEW FILE ALREADY EXIST IN API FOLDER AND IF TRUE REMOVE IN THE CURRENT FOLDER
	for _, f := range currentFolder {
		if checkNameFile(apiFolder, f.Name()) {
			fmt.Printf("API file => %v\n", f.Name())
			e := os.Remove(f.Name())
			if e != nil {
				log.Fatal(e)
				os.Exit(84)
			}
		}
	}
	// CHECK IF NEW FILES AND REPLACE LIKE ALL OTHERS FILES IN API FOLDER
	for _, v := range currentFolder {
		if v.Name() == "http.go" || v.Name() == "main.go" || len(v.Name()) < 3 {
			continue
		} else if len(v.Name()) > 3 && v.Name()[len(v.Name())-3:len(v.Name())] == ".go" {
			if !checkNameFile(apiFolder, v.Name()) {
				fmt.Printf("Move to api folder => %v\n", v.Name())
				err := os.Rename(v.Name(), "./api/"+v.Name())
				if err != nil {
					log.Fatal("error remove file => ", err)
					os.Exit(84)
				}
				input, err := ioutil.ReadFile("./api/" + v.Name())
				if err != nil {
					log.Fatal(err)
					os.Exit(84)
				}

				/////////: WRITE IN THE NEW FILE IN API's FOLDER ////////////
				fmt.Println("Clean files in api's folder")
				if !strings.Contains(v.Name(), "multipart") {
					lines := cleanFolderApi(input, v.Name())
					output := strings.Join(lines, "\n")
					err = ioutil.WriteFile("./api/"+v.Name(), []byte(output), 0644)
					if err != nil {
						log.Fatal("write in new file => ", err)
						os.Exit(84)
					}
				}
				// ////////////////////////////////////////////////////////////

				// /////// ONCE FILE IS CHANGED, ADD THE NEW METHOD IN main.go //
				fmt.Println("Clean main.go")
				method := strings.ReplaceAll(v.Name(), ".go", "")
				if !strings.Contains(v.Name(), "multipart") {
					if err := cleanMain(method); err != nil {
						log.Fatal("change in main.go => ", err)
						os.Exit(84)
					}
				}
				// ////////////////////////////////////////////////////////////

				////// WE ALSO NEED TO ADD THE METHOD IN THE HTTP.GO ////
				fmt.Println("Clean http.go")
				if !strings.Contains(v.Name(), "multipart") {
					if err := cleanHttp(method); err != nil {
						log.Fatal("change in http.go => ", err)
						os.Exit(84)
					}
				}
				// ////////////////////////////////////////////////////:
			}
		}
	}
	// REMOVE CMD'S FOLDER
	godotenv.Load(".env")
	e := os.RemoveAll(fmt.Sprintf("cmd/%s", strings.ToLower(os.Getenv("PROJECT"))))
	if e != nil {
		log.Fatal(e)
	}
	os.Exit(0)
}
