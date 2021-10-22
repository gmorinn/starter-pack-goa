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

	Error("email_already_exist", emailAlreadyExist, "When email already exist")
	Error("unknown_error", unknownError, "Error not identified (500)")
	Error("invalid_scopes", String, "Token scopes are invalid")
	Error("unauthorized", String, "Identifiers are invalid")

	HTTP(func() {
		Response("email_already_exist", StatusBadRequest)
		Response("unknown_error", StatusInternalServerError)
		Response("invalid_scopes", StatusForbidden)
		Response("unauthorized", StatusUnauthorized)
	})

	Method("signup", func() {
		Description("signup")

		Payload(func() {
			Description("Use client ID and client secret to oAuth")
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
				MinLength(8)

			})
			Attribute("email", String, func() {
				Format(FormatEmail)
				Example("guillaume@epitech.eu")
			})
			Required("firstname", "lastname", "password", "email")
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
				MinLength(8)
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
			Required("refresh_token")
		})

		Result(Sign)

		HTTP(func() {
			POST("/resfresh")
			Response(StatusOK)
		})
	})

})

var emailAlreadyExist = Type("emailAlreadyExist", func() {
	Field(1, "message", String)
	Field(2, "success", Boolean, func() {
		Default(false)
	})
	Required("message", "success")
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
