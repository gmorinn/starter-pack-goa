package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func errorRes(name string) string {
	var newName = strings.ReplaceAll(name, ".go", "")
	return fmt.Sprintf(`func (s *%ssrvc) errorResponse(msg string, err error) *%s.UnknownError {
		return &%s.UnknownError{
			Err:       err.Error(),
			ErrorCode: msg,
		}
	}
	
// OAuth2Auth`, newName, newName, newName)
}

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
				input, err := ioutil.ReadFile("./api/" + v.Name())
				if err != nil {
					log.Fatal(err)
				}
				lines := strings.Split(string(input), "\n")
				for i, line := range lines {
					if strings.Contains(line, "{logger}") {
						lines[i] = strings.ReplaceAll(line, "{logger}", "{logger, server}")
					}
					if strings.Contains(line, "(logger *log.Logger") {
						lines[i] = strings.ReplaceAll(line, "(logger *log.Logger", "(logger *log.Logger, server *Server")
					}
					if strings.Contains(line, "\tlogger *log.Logger") {
						lines[i] = strings.ReplaceAll(line, "\tlogger *log.Logger", "\tlogger *log.Logger\n\tserver *Server")
					}
					if strings.Contains(line, "package") {
						lines[i] = "package api"
					}
					if strings.Contains(line, "// OAuth2Auth") {
						lines[i] = strings.ReplaceAll(line, "// OAuth2Auth", errorRes(v.Name()))
					}
					if strings.Contains(line, "JWTAuth(ctx") {
						lines[i+1] = "\treturn s.server.CheckJWT(ctx, token, scheme)"
					}
					if strings.Contains(line, "OAuth2Auth(ctx") {
						lines[i+1] = "\treturn s.server.CheckAuth(ctx, token, scheme)"
					}
					if strings.Contains(line, `return ctx, fmt.Errorf("not implemented")`) {
						lines[i] = ""
					}
				}
				output := strings.Join(lines, "\n")
				err = ioutil.WriteFile("./api/"+v.Name(), []byte(output), 0644)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}
