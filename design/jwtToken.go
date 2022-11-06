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

	Error("unknown_error", unknownError, "Error not identified (500)")

	HTTP(func() {
		Header("oauth:Authorization", String, "OAuth token", func() {
			Pattern("^Bearer [^ ]+$")
		})
		Response("unknown_error", StatusInternalServerError)
	})

	Method("signup", func() {
		Description("signup to generate jwt token")

		Payload(func() {
			Attribute("email", String, func() {
				Format(FormatEmail)
				Example("guillaume@epitech.eu")
			})
			Attribute("password", String, func() {
				Description("Minimum 9 charactères / Chiffre Obligatoire")
				Pattern("\\d")
				MinLength(9)
				Example("JeSuisUnTest974")
			})
			Attribute("confirm_password", String, func() {
				Description("Minimum 9 charactères / Chiffre Obligatoire")
				Pattern("\\d")
				MinLength(9)
				Example("JeSuisUnTest974")
			})
			AccessTokenField(1, "oauth", String, func() {
				Description("Use to generate Oauth with /authorization")
			})
			Required("email", "password", "confirm_password")
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
				Description("Minimum 9 charactères / Chiffre Obligatoire")
				Pattern("\\d")
				MinLength(9)
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
				MinLength(32)
			})
			AccessTokenField(1, "oauth", String, func() {
				Description("Use to generate Oauth with /authorization")
			})
			Required("refresh_token")
		})

		Result(Sign)

		HTTP(func() {
			POST("/refresh")
			Response(StatusOK)
		})
	})
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
