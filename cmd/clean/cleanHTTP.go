package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func cleanHttp(method string) error {
	input, err := ioutil.ReadFile("./http.go")
	if err != nil {
		return err
	}
	mydir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	dir := strings.Split(mydir, "/")
	name := dir[len(dir)-1]
	lines := strings.Split(string(input), "\n")
	for i, line := range lines {
		// Add to the import
		if strings.Contains(line, `jwttokensvr "`) {
			lines[i+1] = fmt.Sprintf("\t"+`%vsvr "%v/gen/http/%v/server"`+"\n", method, name, method) + lines[i+1]
		}
		// if strings.Contains(line, `jwtTokenServer    *jwttokensvr.Server"`) {
		// 	lines[i+1] = fmt.Sprintf("\t%vServer\t*%vsvr.Server\t=%vsvr.New(api.%vEndpoints, mux, dec, enc, eh, nil)\n%v", method, method, method, method, lines[i+1])
		// }
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
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile("./http.go", []byte(output), 0644)
	if err != nil {
		return err
	}
	return nil
}
