package main

import (
	"fmt"
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

func cleanFolderApi(input []byte, name string) []string {
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
			lines[i] = strings.ReplaceAll(line, "// OAuth2Auth", errorRes(name))
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
	return lines
}
