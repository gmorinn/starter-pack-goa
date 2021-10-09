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

// Service describes a service
var _ = Service("crud", func() {
	Description("The principe of CRUD API with GET, PUT, POST, DELETE")

	Error("timeout", func() { // Use default error type
		Timeout()
	})

	Error("unauthorized", String, "Identifiers are invalid")
	Error("id_doesnt_exist", idDoesntExist)
	Error("unknown_error", unknownError, "Error not identified")
	Error("email_already_exist", emailAlreadyExist)

	HTTP(func() {
		Response("unauthorized", StatusUnauthorized)
		Response("id_doesnt_exist", StatusInternalServerError)
		Response("email_already_exist", StatusBadRequest)
		Response("unknown_error", StatusInternalServerError)
	})

	Method("getBook", func() {
		Description("Read Book")
		Payload(func() {
			Attribute("id", String, func() {
				Format(FormatUUID)
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
		Description("Update One Book")
		Payload(func() {
			Attribute("id", String, func() {
				Format(FormatUUID)
			})
			Attribute("name", String)
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
		Description("Read All Books")
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
		Description("Delete Book")
		Payload(func() {
			Attribute("id", String, func() {
				Format(FormatUUID)
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
		Description("Create Book")
		Payload(func() {
			Attribute("name", String, func() {
				MinLength(3)
				MaxLength(10)
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

	Method("signup", func() {
		Description("signup")

		Payload(func() {
			Description("Use client ID and client secret to oAuth")
			Attribute("firstname", String, func() {
				MinLength(3)
			})
			Attribute("lastname", String, func() {
				MinLength(3)
			})
			Attribute("password", String, func() {
				MinLength(7)

			})
			Attribute("email", String, func() {
				Format(FormatEmail)
			})
			Required("firstname", "lastname", "password", "email")
		})

		Result(Register)

		HTTP(func() {
			POST("/signup")
			Response(StatusOK)
		})
	})

})

var _ = Service("openapi", func() {
	Files("/openapi.json", "openapi3.json")
})

var BookResponse = Type("BookResponse", func() {
	Attribute("id", String, func() {
		Format(FormatUUID)
	})
	Attribute("name", String)
	Attribute("price", Float64)
	Required("id", "name", "price")
})

var idDoesntExist = Type("IdDoesntExist", func() {
	Description("IdDoesntExist is the error returned when 0 book have the id corresponding")
	Field(1, "message", String, "Returning error")
	Field(2, "id", String)
	Field(3, "success", Boolean, func() {
		Default(false)
	})
	Required("message", "success", "id")
})

var emailAlreadyExist = Type("emailAlreadyExist", func() {
	Field(1, "message", String)
	Field(2, "success", Boolean, func() {
		Default(false)
	})
	Required("message", "success")
})

var unknownError = Type("unknownError", func() {
	Field(1, "message", String, "Returning error")
	Field(2, "success", Boolean, func() {
		Default(false)
	})
	Required("message", "success")
})

var Register = Type("Register", func() {
	Field(1, "access_token", String, func() {
		Example("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ")
	})
	Field(2, "refresh_token", String, func() {
		Example("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ")
	})
	Field(3, "success", Boolean)
	Required("access_token", "refresh_token", "success")
})
