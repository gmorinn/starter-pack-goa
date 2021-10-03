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
	Attribute("id", String)
	Attribute("name", String)
	Attribute("price", Float64)
})

// Service describes a service
var _ = Service("crud", func() {
	Description("The principe of CRUD API with GET, PUT, POST, DELETE")

	Method("getBook", func() {
		Description("Read Book")
		Payload(func() {
			Attribute("id", String, "Unique Id Book")
			Required("id")
		})
		Error("cannot_convert_string_to_uuid", CannotConvertStringToUuid, "cannot convert string to uuid")
		Error("id_doesnt_exist", IdDoesntExist, "Book with his id doesn't exist")
		HTTP(func() {
			GET("/book/{id}")
			Response("cannot_convert_string_to_uuid", StatusBadRequest, func() {
				Description("Id parameter is bad")
			})
			Response("id_doesnt_exist", StatusBadRequest, func() {
				Description("Book with his id doesn't exist")
			})
			Response(StatusOK)
		})
		Result(BookResponse)
	})

	Method("getAllBooks", func() {
		Description("Read All Books")
		HTTP(func() {
			GET("/books")
			Response(StatusOK)
		})
		Result(ArrayOf(BookResponse))
	})

	Method("deleteBook", func() {
		Description("Delete Book")
		Payload(String, "UUID of an existing book")
		HTTP(func() {
			DELETE("/book/remove/{id}")
			Response(StatusOK)
		})
	})

	Method("createBook", func() {
		Description("Create Book")
		Payload(func() {
			Attribute("name", String)
			Attribute("price", Float64)
		})
		HTTP(func() {
			POST("/book/add")
			Response(StatusCreated)
		})
		Result(BookResponse)
	})

})

var _ = Service("openapi", func() {
	Files("/openapi.json", "openapi3.json")
})

var CannotConvertStringToUuid = Type("CannotConvertStringToUuid", func() {
	Description("CannotConvertStringToUuid is the error returned when id paramater is bad")
	Field(1, "message", String, "Returning error")
	Field(1, "id", String, "Wrong Id")
	Required("message", "id")
})

var IdDoesntExist = Type("IdDoesntExist", func() {
	Description("IdDoesntExist is the error returned when 0 book have the id corresponding")
	Field(1, "message", String, "Returning error")
	Field(1, "id", String, "Wrong Id")
	Required("message", "id")
})
