package design

import (
	. "goa.design/goa/v3/dsl"
)

// API describes the global properties of the API server.
var _ = API("basic", func() {
	Title("Basic CRUD")
	Description("Exemple for api crud")
	Version("1.0")
	Server("crud", func() {
		Host("localhost", func() { URI("http://localhost:8088") })
	})
})

// JWTAuth defines a security scheme that uses JWT tokens.
var JWTAuth = JWTSecurity("jwt", func() {
	Description(`Secures endpoint by requiring a valid JWT token retrieved via the signin endpoint. Supports scopes "api:read" and "api:write".`)
	Scope("api:read", "Read-only access")
	Scope("api:write", "Read and write access")
})

// // OAuth2Auth defines a security scheme that uses OAuth2 tokens.
var OAuth2Auth = OAuth2Security("oauth2", func() {
	AuthorizationCodeFlow("/authorization", "/token", "/refresh")
	Description(`Secures endpoint by requiring a valid OAuth2 token retrieved via the signin endpoint. Supports scopes "api:read" and "api:write".`)
	Scope("api:read", "Read-only access")
})

// Service describes a service
var _ = Service("crud", func() {
	Description("The principe of CRUD API with GET, PUT, POST, DELETE")

	Error("timeout", func() { // Use default error type
		Timeout()
	})

	Error("unauthorized", String, "Identifiers are invalid")
	Error("id_doesnt_exist", idDoesntExist, "When ID doesn't exist")
	Error("unknown_error", unknownError, "Error not identified (500)")
	Error("invalid_scopes", String, "Token scopes are invalid")

	HTTP(func() {
		Response("unauthorized", StatusUnauthorized)
		Response("id_doesnt_exist", StatusInternalServerError)
		Response("unknown_error", StatusInternalServerError)
		Response("invalid_scopes", StatusForbidden)
	})

	Method("getBook", func() {
		Description("Get one item")

		Security(JWTAuth)

		Payload(func() {
			Attribute("id", String, func() {
				Format(FormatUUID)
				Example("5dfb0bf7-597a-4250-b7ad-63a43ff59c25")
			})
			TokenField(1, "jwtToken", String, func() {
				Description("JWT used for authentication")
			})
			Required("id")
		})

		HTTP(func() {
			GET("/book/{id}")
			Response(StatusOK)
		})
		Result(func() {
			Attribute("id", String)
			Attribute("name", String)
			Attribute("price", Float64)
			Attribute("success", Boolean)
			Required("id", "name", "price", "success")
		})

	})

	Method("updateBook", func() {
		Description("Update one item")
		Payload(func() {
			Attribute("id", String, func() {
				Format(FormatUUID)
				Example("5dfb0bf7-597a-4250-b7ad-63a43ff59c25")
			})
			Attribute("name", String, func() {
				Example("Guillaume")
			})
			Attribute("price", Float64)
			Required("id", "name", "price")
		})
		HTTP(func() {
			PUT("/book/{id}")
			Response(StatusOK)
		})
		Result(func() {
			Attribute("id", String)
			Attribute("name", String)
			Attribute("price", Float64)
			Attribute("success", Boolean)
			Required("id", "name", "price", "success")
		})
	})

	Method("getAllBooks", func() {
		Description("Read All items")
		HTTP(func() {
			GET("/books")
			Response(StatusOK)
		})
		Result(func() {
			Attribute("books", ArrayOf(BookResponse))
			Attribute("success", Boolean)
			Required("books", "success")
		})
	})

	Method("deleteBook", func() {
		Description("Delete one item by ID")
		Payload(func() {
			Attribute("id", String, func() {
				Format(FormatUUID)
				Example("5dfb0bf7-597a-4250-b7ad-63a43ff59c25")
			})
			Required("id")
		})
		HTTP(func() {
			DELETE("/book/remove/{id}")
			Response(StatusOK)
		})
		Result(func() {
			Attribute("success", Boolean)
			Required("success")
		})
	})

	Method("createBook", func() {
		Description("Create one item")
		Payload(func() {
			Attribute("name", String, func() {
				MinLength(3)
				MaxLength(10)
				Example("Guillaume")
			})
			Attribute("price", Float64)
			Required("name", "price")
		})
		HTTP(func() {
			POST("/book/add")
			Response(StatusCreated)
		})
		Result(func() {
			Attribute("book", BookResponse)
			Attribute("success", Boolean)
			Required("book", "success")
		})
	})

	Method("oAuth", func() {
		Description("oAuth")

		Payload(func() {
			Attribute("grant_type", String)
			Attribute("client_id", String, func() {
				Example("00000")
			})
			Attribute("client_secret", String, func() {
				Example("99999")
			})
			Required("grant_type", "client_id", "client_secret")
		})

		Result(oAuthResponse)

		HTTP(func() {
			POST("/authorization")
			Response(StatusOK)
		})
	})

})

// Download Postman
var _ = Service("openapi", func() {
	Files("/openapi.json", "openapi3.json")
})

var BookResponse = Type("BookResponse", func() {
	Attribute("id", String, func() {
		Format(FormatUUID)
		Example("5dfb0bf7-597a-4250-b7ad-63a43ff59c25")
	})
	Attribute("name", String, func() {
		Example("Guillaume")
	})
	Attribute("price", Float64)
	Required("id", "name", "price")
})

var idDoesntExist = Type("IdDoesntExist", func() {
	Description("IdDoesntExist is the error returned when 0 book have the id corresponding")
	Field(1, "err", String, "Returning error")
	Field(2, "id", String)
	Field(3, "success", Boolean, func() {
		Default(false)
	})
	Required("err", "success", "id")
})

var unknownError = Type("unknownError", func() {
	Field(1, "err", String)
	Field(2, "error_code", String)
	Field(3, "success", Boolean, func() {
		Default(false)
	})
	Required("err", "success", "error_code")
})

var oAuthResponse = Type("oAuthResponse", func() {
	Field(1, "access_token", String)
	Field(2, "token_type", String)
	Field(3, "expires_in", Int64)
	Field(4, "success", Boolean)
})
