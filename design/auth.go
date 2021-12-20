package design

import . "goa.design/goa/v3/dsl"

var _ = Service("auth", func() {
	Description("Forget password / reset password / send Email Code")

	Security(OAuth2)

	Error("unknown_error", unknownError, "Error not identified (500)")

	HTTP(func() {
		Path("/v1")
		Header("oauth:Authorization", String, "OAuth token", func() {
			Pattern("^Bearer [^ ]+$")
		})
		Response("unknown_error", StatusInternalServerError)
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

	Method("send-confirmation", func() {
		Description("Check if email exist in database and send code by email to reset password")

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
			POST("/lost")
			Response(StatusOK)
		})
	})

	Method("reset-password", func() {
		Description("Reset forget password of the user with the correct confirm code")

		Payload(func() {
			Attribute("email", String, func() {
				Example("guillaume@gmail.com")
				Format(FormatEmail)
			})
			Attribute("code", String, func() {
				Example("ZGI5EV")
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
			AccessTokenField(1, "oauth", String, func() {
				Description("Use to generate Oauth with /authorization")
			})
			Required("email", "code", "password", "confirm_password")
		})

		Result(func() {
			Attribute("success", Boolean)
			Required("success")
		})

		HTTP(func() {
			PUT("/reset-password")
			Response(StatusOK)
		})
	})
})
