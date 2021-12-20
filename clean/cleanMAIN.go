package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func importMain(name, line string) string {
	return fmt.Sprintf(`	%v "api_crud/gen/%v"
	%v`, name, name, line)
}

func cleanMain(method string) error {
	input, err := ioutil.ReadFile("./main.go")
	if err != nil {
		return err
	}
	lines := strings.Split(string(input), "\n")
	for i, line := range lines {
		//ADD NEW METHOD IN STRUCT
		if strings.Contains(line, `auth "api_crud/gen/auth"`) {
			lines[i+1] = importMain(method, lines[i+1])
		}
		if strings.Contains(line, "type ApiEndpoints struct") {
			lines[i+1] = fmt.Sprintf("\t%vEndpoints *%v.Endpoints\n", method, method) + lines[i+1]
		}
		// INITIALIZE IT
		if strings.Contains(line, "jwttoken.Service") {
			lines[i+1] = fmt.Sprintf("\t\t%vSvc %v.Service = api.New%v(logger, server)\n", method, method, strings.Title(method)) + lines[i+1]
		}
		// Wrap the services
		if strings.Contains(line, "apiEndpoints ApiEndpoints = ApiEndpoints{") {
			lines[i+1] = fmt.Sprintf("\t\t\t%vEndpoints: %v.NewEndpoints(%vSvc),\n", method, method, method) + lines[i+1]
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile("./main.go", []byte(output), 0644)
	if err != nil {
		return err
	}
	return nil
}
