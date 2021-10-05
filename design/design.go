package design

import . "goa.design/goa/v3/dsl"

// API describes the global properties of the API server.
var _ = API("basic", func() {
	Title("Basic CRUD")
	Description("Exemple for api crud")
	Version("1.0")
	Server("crud", func() {
		Host("localhost", func() { URI("http://localhost:8088") })
	})
})

var BookResponse = Type("BookResponse", func() {
	Attribute("id", String, func() {
		Format(FormatUUID)
	})
	Attribute("name", String)
	Attribute("price", Float64)
	Required("id", "name", "price")
})

// Service describes a service
var _ = Service("crud", func() {
	Description("The principe of CRUD API with GET, PUT, POST, DELETE")

	Error("timeout", func() { // Use default error type
		Timeout()
	})

	Method("getBook", func() {
		Description("Read Book")
		Payload(func() {
			Attribute("id", String, func() {
				Format(FormatUUID)
			})
			Required("id")
		})

		Error("id_doesnt_exist", IdDoesntExist, "Book with his id doesn't exist")
		Error("unknown_error", unknownError, "Error not identified")
		HTTP(func() {
			GET("/book/{id}")
			Response("id_doesnt_exist", StatusInternalServerError, func() {
				Description("Book with his id doesn't exist")
			})
			Response("unknown_error", StatusInternalServerError, func() {
				Description("Error in general")
			})
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
		Error("id_doesnt_exist", IdDoesntExist, "Book with his id doesn't exist")
		Error("unknown_error", unknownError, "Error not identified")
		HTTP(func() {
			PUT("/book/{id}")
			Response("id_doesnt_exist", StatusInternalServerError, func() {
				Description("Book with his id doesn't exist")
			})
			Response("unknown_error", StatusInternalServerError, func() {
				Description("Error in general")
			})
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
		Error("unknown_error", unknownError, "Error not identified")
		HTTP(func() {
			GET("/books")
			Response("unknown_error", StatusInternalServerError, func() {
				Description("Error in general")
			})
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
		Error("id_doesnt_exist", IdDoesntExist, "Book with his id doesn't exist")
		Error("unknown_error", unknownError, "Error not identified")
		Payload(func() {
			Attribute("id", String, func() {
				Format(FormatUUID)
			})
			Required("id")
		})
		HTTP(func() {
			DELETE("/book/remove/{id}")
			Response("id_doesnt_exist", StatusInternalServerError, func() {
				Description("Book with his id doesn't exist")
			})
			Response("unknown_error", StatusInternalServerError, func() {
				Description("Error in general")
			})
			Response(StatusOK)
		})
		Result(func() {
			Attribute("success", Boolean)
			Required("success")
		})
	})

	Method("createBook", func() {
		Description("Create Book")
		Error("unknown_error", unknownError, "Error not identified")
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
			Response("unknown_error", StatusInternalServerError, func() {
				Description("Error in general")
			})
			Response(StatusCreated)
		})
		Result(func() {
			Attribute("book", BookResponse)
			Attribute("success", Boolean)
			Required("book", "success")
		})
	})

})

var _ = Service("openapi", func() {
	Files("/openapi.json", "openapi3.json")
})

var IdDoesntExist = Type("IdDoesntExist", func() {
	Description("IdDoesntExist is the error returned when 0 book have the id corresponding")
	Field(1, "message", String, "Returning error")
	Field(1, "id", String, "Wrong Id")
	Field(1, "success", Boolean)
	Required("message", "id", "success")
})

var unknownError = Type("unknownError", func() {
	Field(1, "message", String, "Returning error")
	Field(1, "success", Boolean)
	Required("message", "success")
})
