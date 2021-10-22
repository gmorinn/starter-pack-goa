package design

import (
	"api_crud/config"

	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
	_ "goa.design/plugins/v3/docs"
)

// API describes the global properties of the API server.
var _ = API("basic", func() {

	// Get .env
	cnf := config.New()

	cors.Origin("/.*localhost.*/", func() {
		cors.Methods("POST", "GET", "PUT", "OPTIONS", "DELETE", "PATCH")
		cors.Credentials()
		cors.Headers("Authorization", "Content-Type", "jwtToken")
	})

	Title("Basic CRUD")
	Description("Stater Pack")
	Version("1.0")
	Server("crud", func() {
		Host(cnf.Domain, func() {
			URI(cnf.Host)
		})
	})
})

// Download Postman
var _ = Service("openapi", func() {
	Files("/openapi.json", "openapi3.json")
})
