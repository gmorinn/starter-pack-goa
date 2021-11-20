package design

import . "goa.design/goa/v3/dsl"

// Service describes a service
var _ = Service("boProducts", func() {
	Description("Products BO of the E-Commerce")

	Error("timeout", func() { // Use default error type
		Timeout()
	})

	Security(OAuth2, JWTAuth)

	Error("unknown_error", unknownError, "Error not identified (500)")

	HTTP(func() {
		Path("/v1/bo")
		Header("oauth:Authorization", String, "OAuth token", func() {
			Pattern("^Bearer [^ ]+$")
		})
		Header("jwtToken:jwtToken", String, "Jwt token", func() {
			Pattern("^Bearer [^ ]+$")
		})
		Response("unknown_error", StatusInternalServerError)
	})

	Method("getAllProducts", func() {
		Description("Get All products")
		Payload(func() {
			TokenField(1, "jwtToken", String, func() {
				Description("JWT used for authentication after Signin/Signup")
			})
			AccessTokenField(2, "oauth", String, func() {
				Description("Use to generate Oauth with /authorization")
			})
		})
		HTTP(func() {
			GET("/products")
			Response(StatusOK)
		})
		Result(func() {
			Attribute("products", ArrayOf(resBoProduct), "All products by category")
			Attribute("success", Boolean)
			Required("products", "success")
		})
	})

	Method("getAllProductsByCategory", func() {
		Description("Get All products by category")
		Payload(func() {
			Attribute("category", String, func() {
				Enum("men", "women", "hat", "jacket", "sneaker", "nothing")
				Example("men")
				Default("nothing")
			})
			TokenField(1, "jwtToken", String, func() {
				Description("JWT used for authentication after Signin/Signup")
			})
			AccessTokenField(2, "oauth", String, func() {
				Description("Use to generate Oauth with /authorization")
			})
			Required("category")
		})
		HTTP(func() {
			GET("/products/category/{category}")
			Response(StatusOK)
		})
		Result(func() {
			Attribute("products", ArrayOf(resBoProduct), "Result is an array of object")
			Attribute("success", Boolean)
			Required("products", "success")
		})
	})

	Method("deleteProduct", func() {
		Description("Delete one product by ID")
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
			Required("id")
		})
		HTTP(func() {
			DELETE("/product/remove/{id}")
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

	Method("createProduct", func() {
		Description("Create one product")
		Security(OAuth2, JWTAuth)
		Payload(func() {
			Attribute("product", payloadProduct)
			TokenField(1, "jwtToken", String, func() {
				Description("JWT used for authentication after Signin/Signup")
			})
			AccessTokenField(2, "oauth", String, func() {
				Description("Use to generate Oauth with /authorization")
			})
			Required("product")
		})
		HTTP(func() {
			POST("/product/add")
			Header("jwtToken:jwtToken", String, "Jwt token", func() {
				Pattern("^Bearer [^ ]+$")
			})
			Response(StatusCreated)
		})
		Result(func() {
			Attribute("product", resBoProduct, "Result is an object")
			Attribute("success", Boolean)
			Required("product", "success")
		})
	})

	Method("updateProduct", func() {
		Description("Update one product")
		Security(OAuth2, JWTAuth)
		Payload(func() {
			Attribute("id", String, func() {
				Format(FormatUUID)
				Example("5dfb0bf7-597a-4250-b7ad-63a43ff59c25")
			})
			Attribute("product", payloadProduct)
			TokenField(1, "jwtToken", String, func() {
				Description("JWT used for authentication after Signin/Signup")
			})
			AccessTokenField(2, "oauth", String, func() {
				Description("Use to generate Oauth with /authorization")
			})
			Required("product", "id")
		})
		HTTP(func() {
			PUT("/product/{id}")
			Header("jwtToken:jwtToken", String, "Jwt token", func() {
				Pattern("^Bearer [^ ]+$")
			})
			Response(StatusOK)
		})
		Result(func() {
			Attribute("product", resBoProduct, "Result is an Object")
			Attribute("success", Boolean)
			Required("product", "success")
		})
	})

	Method("getProduct", func() {
		Description("Get one product")
		Payload(func() {
			Attribute("id", String, func() {
				Format(FormatUUID)
				Description("Unique ID of the product")
				Example("5dfb0bf7-597a-4250-b7ad-63a43ff59c25")
			})
			TokenField(1, "jwtToken", String, func() {
				Description("JWT used for authentication after Signin/Signup")
			})
			AccessTokenField(2, "oauth", String, func() {
				Description("Use to generate Oauth with /authorization")
			})
			Required("id")
		})

		HTTP(func() {
			GET("/product/{id}")
			Response(StatusOK)
		})
		Result(func() {
			Attribute("product", resBoProduct, "Result is an object")
			Attribute("success", Boolean)
			Required("product", "success")
		})
	})
})

var resBoProduct = Type("resBoProduct", func() {
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
	Attribute("cover", String, func() {
		Example("https://i.ibb.co/ypkgK0X/blue-beanie.png")
	})
	Attribute("category", String, func() {
		Example("men")
	})
	Required("id", "name", "price", "cover", "category")
})

var payloadProduct = Type("payloadProduct", func() {
	Attribute("name", String, func() {
		Example("Guillaume")
		MinLength(3)
	})
	Attribute("price", Float64, func() {
		Example(69.0)
		Minimum(0)
	})
	Attribute("cover", String, func() {
		Example("https://i.ibb.co/ypkgK0X/blue-beanie.png")
	})
	Attribute("category", String, func() {
		Enum("men", "women", "hat", "jacket", "sneaker", "nothing")
		Example("men")
		Default("nothing")
	})
	Required("name", "price", "cover", "category")
})
