package design

import . "goa.design/goa/v3/dsl"

// Service describes a service
var _ = Service("book", func() {
	Description("The principe of CRUD API with GET, PUT, POST, DELETE")

	Error("timeout", func() { // Use default error type
		Timeout()
	})

	Error("id_doesnt_exist", idDoesntExist, "When ID doesn't exist")
	Error("unknown_error", unknownError, "Error not identified (500)")

	HTTP(func() {
		Response("id_doesnt_exist", StatusInternalServerError)
		Response("unknown_error", StatusInternalServerError)
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
