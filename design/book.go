package design

import . "goa.design/goa/v3/dsl"

// Service describes a service
var _ = Service("book", func() {
	Description("The principe of CRUD API with GET, PUT, POST, DELETE with Table Book")

	Error("timeout", func() { // Use default error type
		Timeout()
	})

	Error("unknown_error", unknownError, "Error not identified (500)")
	Error("unauthorized", String, "Credentials are invalid")

	HTTP(func() {
		Path("/web")
		Response("unknown_error", StatusInternalServerError)
		Response("unauthorized", StatusUnauthorized)
	})

	Method("getBook", func() {
		Description("Get one item")
		Security(OAuth2, JWTAuth)
		Payload(func() {
			Attribute("id", String, func() {
				Format(FormatUUID)
				Description("Unique ID of the book")
				Example("5dfb0bf7-597a-4250-b7ad-63a43ff59c25")
			})
			TokenField(1, "jwtToken", String, func() {
				Description("JWT used for authentication after Signin/Signup")
			})
			AccessTokenField(2, "oauth", String, func() {
				Description("Use to generate Oauth with /authorization")
			})
			Required("id", "oauth", "jwtToken")
		})

		HTTP(func() {
			GET("/book/{id}")
			Header("oauth:Authorization", String, "OAuth token", func() {
				Pattern("^Bearer [^ ]+$")
			})
			Header("jwtToken:jwtToken", String, "Jwt token", func() {
				Pattern("^Bearer [^ ]+$")
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
		Description("Update one item")
		Security(OAuth2, JWTAuth)
		Payload(func() {
			Attribute("id", String, func() {
				Format(FormatUUID)
				Example("5dfb0bf7-597a-4250-b7ad-63a43ff59c25")
			})
			Attribute("name", String, func() {
				Example("Guillaume")
				MinLength(3)
				MaxLength(10)
			})
			Attribute("price", Float64, func() {
				Example(69.0)
				Minimum(0.1)
			})
			TokenField(1, "jwtToken", String, func() {
				Description("JWT used for authentication after Signin/Signup")
			})
			AccessTokenField(2, "oauth", String, func() {
				Description("Use to generate Oauth with /authorization")
			})
			Required("id", "name", "price", "oauth", "jwtToken")
		})
		HTTP(func() {
			PUT("/book/{id}")
			Header("oauth:Authorization", String, "OAuth token", func() {
				Pattern("^Bearer [^ ]+$")
			})
			Header("jwtToken:jwtToken", String, "Jwt token", func() {
				Pattern("^Bearer [^ ]+$")
			})
			Response(StatusOK)
		})
		Result(func() {
			Attribute("id", String, func() {
				Example("5dfb0bf7-597a-4250-b7ad-63a43ff59c25")
			})
			Attribute("name", String, func() {
				Example("Père riche père pauvre")
			})
			Attribute("price", Float64, func() {
				Example(14.5)
			})
			Attribute("success", Boolean, func() {
				Example(true)
			})
			Required("id", "name", "price", "success")
		})
	})

	Method("getAllBooks", func() {
		Description("Get All items")
		Payload(func() {
			Required("oauth", "jwtToken")
		})
		HTTP(func() {
			GET("/books")
			Response(StatusOK)
		})
		Result(func() {
			Attribute("books", ArrayOf(BookResponse), "Result is an array of object")
			Attribute("success", Boolean)
			Required("books", "success")
		})
	})

	Method("deleteBook", func() {
		Description("Delete one item by ID")
		Security(OAuth2, JWTAuth)
		Payload(func() {
			Attribute("id", String, func() {
				Format(FormatUUID)
				Example("5dfb0bf7-597a-4250-b7ad-63a43ff59c25")
			})
			TokenField(1, "jwtToken", String, func() {
				Description("JWT used for authentication after Signin/Signup")
			})
			AccessTokenField(2, "oauth", String, func() {
				Description("Use to generate Oauth with /authorization")
			})
			Required("id", "oauth", "jwtToken")
		})
		HTTP(func() {
			DELETE("/book/remove/{id}")
			Header("oauth:Authorization", String, "OAuth token", func() {
				Pattern("^Bearer [^ ]+$")
			})
			Header("jwtToken:jwtToken", String, "Jwt token", func() {
				Pattern("^Bearer [^ ]+$")
			})
			Response(StatusOK)
		})
		Result(func() {
			Attribute("success", Boolean)
			Required("success")
		})
	})

	Method("createBook", func() {
		Description("Create one item")
		Security(OAuth2, JWTAuth)
		Payload(func() {
			Attribute("name", String, func() {
				MinLength(3)
				MaxLength(10)
				Description("Name of the book")
				Example("Guillaume")
			})
			Attribute("price", Float64, func() {
				Description("Price of the book")
				Minimum(0.1)
			})
			TokenField(1, "jwtToken", String, func() {
				Description("JWT used for authentication after Signin/Signup")
			})
			AccessTokenField(2, "oauth", String, func() {
				Description("Use to generate Oauth with /authorization")
			})
			Required("name", "price", "oauth", "jwtToken")
		})
		HTTP(func() {
			POST("/book/add")
			Header("oauth:Authorization", String, "OAuth token", func() {
				Pattern("^Bearer [^ ]+$")
			})
			Header("jwtToken:jwtToken", String, "Jwt token", func() {
				Pattern("^Bearer [^ ]+$")
			})
			Response(StatusCreated)
		})
		Result(func() {
			Attribute("book", BookResponse, "Result is an object")
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
	Attribute("price", Float64, func() {
		Example(69.0)
	})
	Required("id", "name", "price")
})

var unknownError = Type("unknownError", func() {
	Field(1, "err", String, func() {
		Example("sql no rows affected")
	})
	Field(2, "error_code", String, func() {
		Example("TX_UPDATE_ITEM")
	})
	Field(3, "success", Boolean, func() {
		Default(false)
	})
	Required("err", "success", "error_code")
})
