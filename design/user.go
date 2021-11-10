package design

import . "goa.design/goa/v3/dsl"

// Service describes a service
var _ = Service("users", func() {
	Description("users of the api")

	Error("timeout", func() { // Use default error type
		Timeout()
	})

	Security(OAuth2, JWTAuth)

	Error("unknown_error", unknownError, "Error not identified (500)")

	HTTP(func() {
		Path("/v1/web")
		Header("oauth:Authorization", String, "OAuth token", func() {
			Pattern("^Bearer [^ ]+$")
		})
		Header("jwtToken:jwtToken", String, "Jwt token", func() {
			Pattern("^Bearer [^ ]+$")
		})
		Response("unknown_error", StatusInternalServerError)
	})

	Method("getAllusers", func() {
		Description("Get All users")
		Payload(func() {
			TokenField(1, "jwtToken", String, func() {
				Description("JWT used for authentication after Signin/Signup")
			})
			AccessTokenField(2, "oauth", String, func() {
				Description("Use to generate Oauth with /authorization")
			})
		})
		HTTP(func() {
			GET("/users")
			Response(StatusOK)
		})
		Result(func() {
			Attribute("users", ArrayOf(resUser), "All users by category")
			Attribute("success", Boolean)
			Required("users", "success")
		})
	})

	Method("deleteUser", func() {
		Description("Delete one User by ID")
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
			DELETE("/user/remove/{id}")
			Response(StatusOK)
		})
		Result(func() {
			Attribute("success", Boolean)
			Required("success")
		})
	})

	Method("createUser", func() {
		Description("Create one User")
		Payload(func() {
			Attribute("user", payloadUser)
			TokenField(1, "jwtToken", String, func() {
				Description("JWT used for authentication after Signin/Signup")
			})
			AccessTokenField(2, "oauth", String, func() {
				Description("Use to generate Oauth with /authorization")
			})
			Required("user")
		})
		HTTP(func() {
			POST("/user/add")
			Response(StatusCreated)
		})
		Result(func() {
			Attribute("user", resUser, "Result is an object")
			Attribute("success", Boolean)
			Required("user", "success")
		})
	})

	Method("updateUser", func() {
		Description("Update one User")
		Payload(func() {
			Attribute("id", String, func() {
				Format(FormatUUID)
				Example("5dfb0bf7-597a-4250-b7ad-63a43ff59c25")
			})
			Attribute("User", payloadUser)
			TokenField(1, "jwtToken", String, func() {
				Description("JWT used for authentication after Signin/Signup")
			})
			AccessTokenField(2, "oauth", String, func() {
				Description("Use to generate Oauth with /authorization")
			})
			Required("User", "id")
		})
		HTTP(func() {
			PUT("/user/{id}")
			Response(StatusOK)
		})
		Result(func() {
			Attribute("user", resUser, "Result is an Object")
			Attribute("success", Boolean)
			Required("user", "success")
		})
	})

	Method("getUser", func() {
		Description("Get one User")
		Payload(func() {
			Attribute("id", String, func() {
				Format(FormatUUID)
				Description("Unique ID of the User")
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
			GET("/user/{id}")
			Response(StatusOK)
		})
		Result(func() {
			Attribute("user", resUser, "Result is an object")
			Attribute("success", Boolean)
			Required("user", "success")
		})
	})
})

var resUser = Type("resUser", func() {
	Attribute("id", String, func() {
		Format(FormatUUID)
		Example("5dfb0bf7-597a-4250-b7ad-63a43ff59c25")
	})
	Attribute("firstname", String, func() {
		Example("Guillaume")
	})
	Attribute("lastname", String, func() {
		Example("Morin")
	})
	Attribute("email", String, func() {
		Format(FormatEmail)
		Example("guillaume.morin@epitech.eu")
	})
	Attribute("birthday", String, func() {
		Default("")
		Example("01/09/2002")
	})
	Attribute("phone", String, func() {
		Default("")
		Example("+262 692 12 34 56")
	})
	Required("id", "email")
})

var payloadUser = Type("payloadUser", func() {
	Attribute("firstname", String, func() {
		Example("Guillaume")
	})
	Attribute("lastname", String, func() {
		Example("Morin")
	})
	Attribute("email", String, func() {
		Format(FormatEmail)
		Example("guillaume.morin@epitech.eu")
	})
	Attribute("birthday", String, func() {
		Default("")
		Example("01/09/2002")
	})
	Attribute("phone", String, func() {
		Default("")
		Example("+262 692 12 34 56")
	})
})
