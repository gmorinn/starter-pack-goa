package design

import . "goa.design/goa/v3/dsl"

// JWTAuth defines a security scheme that uses JWT tokens.
var JWTAuth = JWTSecurity("jwt", func() {
	Description(`Secures endpoint by requiring a valid JWT token retrieved via the signin endpoint. Supports scopes "api:read" and "api:write".`)
	Scope("api:read", "Read-only access")
	Scope("api:write", "Read and write access")
})

var _ = Service("jwtToken", func() {
	Description("Use Token to authenticate. Signin and Signup")

	Security(OAuth2)

	Error("email_already_exist", emailAlreadyExist, "When email already exist")
	Error("unknown_error", unknownError, "Error not identified (500)")
	Error("invalid_scopes", String, "Token scopes are invalid")
	Error("unauthorized", String, "Identifiers are invalid")

	HTTP(func() {
		Header("oauth:Authorization", String, "OAuth token", func() {
			Pattern("^Bearer [^ ]+$")
		})
		Response("email_already_exist", StatusBadRequest)
		Response("unknown_error", StatusInternalServerError)
		Response("invalid_scopes", StatusForbidden)
		Response("unauthorized", StatusUnauthorized)
	})

	Method("signup", func() {
		Description("signup to generate jwt token")

		Payload(func() {
			Attribute("firstname", String, func() {
				MinLength(3)
				MaxLength(15)
				Example("Guillaume")
			})
			Attribute("lastname", String, func() {
				MinLength(3)
				Example("Morin")
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
			Attribute("email", String, func() {
				Format(FormatEmail)
				Example("guillaume@epitech.eu")
			})
			Attribute("birthday", String, func() {
				Default("")
			})
			Attribute("phone", String, func() {
				Default("")
				Example("+262 692 12 34 56")
			})
			AccessTokenField(1, "oauth", String, func() {
				Description("Use to generate Oauth with /authorization")
			})
			Required("firstname", "lastname", "password", "email", "confirm_password")
		})

		Result(Sign)

		HTTP(func() {
			POST("/signup")
			Response(StatusOK)
		})
	})

	Method("signin", func() {
		Description("signin")

		Payload(func() {
			Attribute("email", String, func() {
				Format(FormatEmail)
				Example("guillaume@epitech.eu")
			})
			Attribute("password", String, func() {
				Description("Minimum 8 charactères / Chiffre Obligatoire")
				Pattern("\\d")
				MinLength(8)
				Example("JeSuisUnTest974")
			})
			AccessTokenField(1, "oauth", String, func() {
				Description("Use to generate Oauth with /authorization")
			})
			Required("password", "email")
		})

		Result(Sign)

		HTTP(func() {
			POST("/signin")
			Response(StatusOK)
		})
	})

	Method("refresh", func() {
		Description("Refresh Token")

		Payload(func() {
			Attribute("refresh_token", String, func() {
				Example("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ")
			})
			AccessTokenField(1, "oauth", String, func() {
				Description("Use to generate Oauth with /authorization")
			})
			Required("refresh_token")
		})

		Result(Sign)

		HTTP(func() {
			POST("/resfresh")
			Response(StatusOK)
		})
	})

	Method("email-exist", func() {
		Description("Check if email exist in database")

		Payload(func() {
			Attribute("email", String, func() {
				Example("guillaume@gmail.com")
				Format(FormatEmail)
			})
			AccessTokenField(1, "oauth", String, func() {
				Description("Use to generate Oauth with /authorization")
			})
			Required("email")
		})

		Result(func() {
			Attribute("success", Boolean)
			Attribute("exist", Boolean)
			Required("exist", "success")
		})

		HTTP(func() {
			POST("/email-exist")
			Response(StatusOK)
		})
	})

	Method("auth-providers", func() {
		Description("Register or login by Google, Facebook")

		Payload(func() {
			Attribute("firstname", String, func() {
				MinLength(3)
				MaxLength(15)
				Example("Guillaume")
			})
			Attribute("lastname", String, func() {
				MinLength(3)
				Example("Morin")
			})
			Attribute("email", String, func() {
				Format(FormatEmail)
				Example("guillaume@epitech.eu")
			})
			Attribute("firebase_id_token", String, func() {
				MinLength(400)
			})
			Attribute("firebase_uid", String, func() {
				MinLength(15)
				Example("zgmURRUlcJfgDMRyjJ20xs7Rxxw2")
			})
			Attribute("firebase_provider", String, func() {
				Example("facebook.com")
			})
			AccessTokenField(1, "oauth", String, func() {
				Description("Use to generate Oauth with /authorization")
			})
			Required("email", "firstname", "lastname", "firebase_id_token", "firebase_uid", "firebase_provider")
		})

		Result(Sign)

		HTTP(func() {
			POST("/sign-providers")
			Response(StatusCreated)
			Response(StatusBadRequest)
		})
	})

})

var emailAlreadyExist = Type("emailAlreadyExist", func() {
	Field(1, "err", String)
	Field(2, "success", Boolean, func() {
		Default(false)
	})
	Required("err", "success")
})

var Sign = Type("Sign", func() {
	Field(1, "access_token", String, func() {
		Example("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ")
	})
	Field(2, "refresh_token", String, func() {
		Example("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ")
	})
	Field(3, "success", Boolean)
	Required("access_token", "refresh_token", "success")
})
