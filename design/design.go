package design

import (
	"api_crud/config"

	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
	_ "goa.design/plugins/v3/docs"
)

// API describes the global properties of the API server.
var _ = API("basic", func() {

	Title("Basic CRUD")
	Description("Stater Pack")
	Version("1.0")

	// Get .env
	cnf := config.New()

	cors.Origin("*", func() {
		cors.Headers("Authorization", "Content-Type", "jwtToken")
		cors.Methods("POST", "GET", "PUT", "OPTIONS", "DELETE", "PATCH")
		cors.Expose("Content-Type", "Origin")
		cors.MaxAge(100)
		cors.Credentials()
	})

	Server("crud", func() {
		Host(cnf.Domain, func() {
			URI(cnf.Host)
		})
	})

	Contact(func() {
		Name("Guillaume Morin")
		Email("guillaume.morin@epitech.eu")
		URL("https://guillaume-morin.fr")
	})

	License(func() {
		Name("Apache License Version 2.0")
		URL("https://guillaume-morin.fr")
	})

})

// Download Postman
var _ = Service("openapi", func() {
	Files("/openapi.json", "openapi3.json")
})
