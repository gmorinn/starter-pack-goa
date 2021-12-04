package design

import . "goa.design/goa/v3/dsl"

// Service describes a service
var _ = Service("products", func() {
	Description("Products of the E-Commerce")

	Error("timeout", func() { // Use default error type
		Timeout()
	})

	Security(OAuth2)

	Error("unknown_error", unknownError, "Error not identified (500)")

	HTTP(func() {
		Path("/v1/web")
		Header("oauth:Authorization", String, "OAuth token", func() {
			Pattern("^Bearer [^ ]+$")
		})
		Response("unknown_error", StatusInternalServerError)
	})

	Method("getAllProductsByCategory", func() {
		Description("Get All products by category")
		Payload(func() {
			Attribute("category", String, func() {
				Enum("men", "women", "hat", "jacket", "sneaker", "nothing")
				Example("men")
				Default("nothing")
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
			Attribute("products", ArrayOf(resProduct), "Result is an array of object")
			Attribute("success", Boolean)
			Required("products", "success")
		})
	})

	Method("getAllProducts", func() {
		Description("Get All products")
		Payload(func() {
			AccessTokenField(2, "oauth", String, func() {
				Description("Use to generate Oauth with /authorization")
			})
		})
		HTTP(func() {
			GET("/products")
			Response(StatusOK)
		})
		Result(func() {
			Attribute("products", ArrayOf(resProduct), "Result is an array of object")
			Attribute("success", Boolean)
			Required("products", "success")
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
			Attribute("product", resProduct, "Result is an object")
			Attribute("success", Boolean)
			Required("product", "success")
		})
	})
})

var resProduct = Type("resProduct", func() {
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
