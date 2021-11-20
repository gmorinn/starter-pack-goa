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
	//OPEN CURRENT FOLDER
	currentFolder, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}
	// OPEN API FOLDER
	apiFolder, err := ioutil.ReadDir("./api")
	if err != nil {
		log.Fatal(err)
	}
	// CHECK IF NEW FILE ALREADY EXIST IN API FOLDER AND IF TRUE REMOVE IN THE CURRENT FOLDER
	for _, f := range currentFolder {
		if checkNameFile(apiFolder, f.Name()) {
			e := os.Remove(f.Name())
			if e != nil {
				log.Fatal(e)
			}
		}
	}
	// CHECK IF NEW FILES AND REPLACE LIKE ALL OTHERS FILES IN API FOLDER
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
				// ONCE FILE IS CHANGED, ADD THE NEW METHOD IN main.go
				input, err = ioutil.ReadFile("./main.go")
				if err != nil {
					log.Fatal(err)
				}
				lines = strings.Split(string(input), "\n")
				method := strings.ReplaceAll(v.Name(), ".go", "")
				for i, line := range lines {
					//ADD NEW METHOD IN STRUCT
					if strings.Contains(line, "type ApiEndpoints struct") {
						lines[i+1] = fmt.Sprintf("\t%vEndpoints *%v.Endpoints\n", method, method) + lines[i+1]
					}
					// INITIALIZE IT
					if strings.Contains(line, "// Initialize the services.") {
						lines[i+2] = fmt.Sprintf("\t\t%vSvc %v.Service = api.New%v(logger, server)\n", method, method, strings.Title(method)) + lines[i+1]
					}
					// Wrap the services
					if strings.Contains(line, "apiEndpoints ApiEndpoints = ApiEndpoints{") {
						lines[i+1] = fmt.Sprintf("\t\t\t%vEndpoints: %v.NewEndpoints(%vSvc),\n", method, method, method) + lines[i+1]
					}
				}
				output = strings.Join(lines, "\n")
				err = ioutil.WriteFile("./main.go", []byte(output), 0644)
				if err != nil {
					log.Fatal(err)
				}
				// WE ALSO NEED TO ADD THE METHOD IN THE HTTP.GO
				input, err = ioutil.ReadFile("./http.go")
				if err != nil {
					log.Fatal(err)
				}
				lines = strings.Split(string(input), "\n")
				for i, line := range lines {
					// Add to the import
					if strings.Contains(line, `jwttokensvr "`) {
						lines[i+1] = fmt.Sprintf("\t"+`%vsvr "/gen/http/%v/server"`+"\n", method, method) + lines[i+1]
					}
					// Build the service HTTP request
					if strings.Contains(line, "openapiServer  *openapisvr.Server") {
						lines[i+1] = fmt.Sprintf("\t\t\t%vServer *%vsvr.Server = %vsvr.New(api.%vEndpoints, mux, dec, enc, eh, nil)\n", method, method, method, method) + lines[i+1]
					}
					if strings.Contains(line, "servers := goahttp.Servers{") {
						lines[i+1] = fmt.Sprintf("\t\t\t\t%vServer,\n", method) + lines[i+1]
					}
					// Configure the mux.
					if strings.Contains(line, "openapisvr.Mount(mux, openapiServer)") {
						lines[i+1] = fmt.Sprintf("\t%vsvr.Mount(mux, %vServer)\n", method, method) + lines[i+1]
					}
				}
				output = strings.Join(lines, "\n")
				err = ioutil.WriteFile("./http.go", []byte(output), 0644)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}
