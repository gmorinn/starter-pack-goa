package design

import (
	. "goa.design/goa/v3/dsl"
)

// Service describes a service
var _ = Service("boUsers", func() {
	Description("users of the api")

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

	Method("getAllusers", func() {
		Description("Get All users")
		Payload(func() {
			Attribute("offset", Int32, func() {
				Description("Offset for pagination")
				Example(0)
			})
			Attribute("limit", Int32, func() {
				Description("Limit of items listed for pagination")
				Example(9)
			})
			Attribute("field", String, func() {
				Description("Items order by {field}")
				Example("name")
				Default("name")
			})
			Attribute("direction", String, func() {
				Description("Items order by {field} ASC/DESC")
				Enum("asc", "desc")
				Example("asc")
				Default("asc")
			})
			TokenField(1, "jwtToken", String, func() {
				Description("JWT used for authentication after Signin/Signup")
			})
			AccessTokenField(2, "oauth", String, func() {
				Description("Use to generate Oauth with /authorization")
			})
			Required("limit", "offset")
		})
		HTTP(func() {
			GET("/users/{offset}/{limit}")
			Params(func() {
				Param("field", String, func() {
					Description("Items order by {field}")
					Example("name")
					Default("name")
				})
				Param("direction", String, func() {
					Description("Items order by {field} ASC/DESC")
					Enum("asc", "desc")
					Example("asc")
					Default("asc")
				})
			})
			Response(StatusOK)
		})
		Result(func() {
			Attribute("users", ArrayOf(resBoUser), "All users")
			Attribute("count", Int64, "total of users")
			Attribute("success", Boolean)
			Required("users", "success", "count")
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
			Attribute("firstname", String, func() {
				Example("Guillaume")
				MinLength(3)
			})
			Attribute("lastname", String, func() {
				Example("Morin")
				MinLength(3)
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
			Attribute("role", String, func() {
				Default("user")
				Enum("user", "pro", "admin")
				Example("user")
			})
			Attribute("password", String, func() {
				Description("Minimum 8 charactères / Chiffre Obligatoire")
				Pattern("\\d")
				MinLength(8)
				Example("JeSuisUnTest974")
			})
			Attribute("confirm_password", String, func() {
				Description("Minimum 8 charactères / Chiffre Obligatoire")
				Pattern("\\d")
				MinLength(8)
				Example("JeSuisUnTest974")
			})
			TokenField(1, "jwtToken", String, func() {
				Description("JWT used for authentication after Signin/Signup")
			})
			AccessTokenField(2, "oauth", String, func() {
				Description("Use to generate Oauth with /authorization")
			})
			Required("firstname", "lastname", "email", "password", "confirm_password")
		})
		HTTP(func() {
			POST("/user/add")
			Response(StatusCreated)
		})
		Result(func() {
			Attribute("user", resBoUser, "Result is an object")
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
			Attribute("user", resBoUser, "Result is an Object")
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
			Attribute("user", resBoUser, "Result is an object")
			Attribute("success", Boolean)
			Required("user", "success")
		})
	})

	Method("deleteManyUsers", func() {
		Description("Delete many users with IDs send in body")
		Security(OAuth2, JWTAuth)
		Payload(func() {
			Attribute("tab", ArrayOf(String))
			TokenField(1, "jwtToken", String, func() {
				Description("JWT used for authentication after Signin/Signup")
			})
			AccessTokenField(2, "oauth", String, func() {
				Description("Use to generate Oauth with /authorization")
			})
			Required("tab")
		})
		HTTP(func() {
			PATCH("/users/remove")
			Response(StatusOK)
		})
		Result(func() {
			Attribute("success", Boolean)
			Required("success")
		})
	})

	Method("newPassword", func() {
		Description("Change your current password from Back Office")

		Payload(func() {
			Attribute("id", String, func() {
				Format(FormatUUID)
				Description("Unique ID of the User")
				Example("5dfb0bf7-597a-4250-b7ad-63a43ff59c25")
			})
			Attribute("password", String, func() {
				Description("Minimum 8 charactères / Chiffre Obligatoire")
				Pattern("\\d")
				MinLength(8)
				Example("JeSuisUnTest974")
			})
			Attribute("confirm", String, func() {
				Description("Minimum 8 charactères / Chiffre Obligatoire")
				Pattern("\\d")
				MinLength(8)
				Example("JeSuisUnTest974")
			})
			TokenField(1, "jwtToken", String, func() {
				Description("JWT used for authentication after Signin/Signup")
			})
			AccessTokenField(2, "oauth", String, func() {
				Description("Use to generate Oauth with /authorization")
			})
			Required("id", "password", "confirm")
		})

		Result(func() {
			Attribute("success", Boolean)
			Required("success")
		})

		HTTP(func() {
			PATCH("/user/change/password/{id}")
			Response(StatusOK)
		})
	})
})

var resBoUser = Type("resBoUser", func() {
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
	Attribute("role", String, func() {
		Default("user")
		Enum("user", "pro", "admin")
		Example("user")
	})
	Required("id", "email")
})

var payloadUser = Type("payloadUser", func() {
	Attribute("firstname", String, func() {
		Example("Guillaume")
		MinLength(3)
	})
	Attribute("lastname", String, func() {
		Example("Morin")
		MinLength(3)
	})
	Attribute("email", String, func() {
		Format(FormatEmail)
		Example("guillaume.morin@epitech.eu")
	})
	Attribute("birthday", String, func() {
		Default("")
		Example("01/09/2002")
	})
	Attribute("role", String, func() {
		Default("user")
		Enum("user", "pro", "admin")
		Example("user")
	})
	Attribute("phone", String, func() {
		Default("")
		Example("+262 692 12 34 56")
	})
	Required("email", "firstname", "lastname")
})
