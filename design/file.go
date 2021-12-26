package design

import (
	. "goa.design/goa/v3/dsl"
)

// Service describes a service
var _ = Service("files", func() {
	Description("files of the api")

	Error("timeout", func() { // Use default error type
		Timeout()
	})

	Security(OAuth2, JWTAuth)

	Error("unknown_error", unknownError, "Error not identified (500)")

	HTTP(func() {
		Path("/v1/bo/files")
		Header("oauth:Authorization", String, "OAuth token", func() {
			Pattern("^Bearer [^ ]+$")
		})
		Header("jwtToken:jwtToken", String, "Jwt token", func() {
			Pattern("^Bearer [^ ]+$")
		})
		Response("unknown_error", StatusInternalServerError)
	})

	Method("importFile", func() {
		Description("Import file")
		Payload(func() {
			Attribute("file_name", String, "uploaded file name", func() {
				Example("foo.jpg")
			})
			Attribute("content", Bytes, "content of image")
			Attribute("format", String, "uploaded file format", func() {
				Example("image/jpeg")
			})
			TokenField(1, "jwtToken", String, func() {
				Description("JWT used for authentication after Signin/Signup")
			})
			AccessTokenField(2, "oauth", String, func() {
				Description("Use to generate Oauth with /authorization")
			})
			Required(
				"file_name",
				"content",
				"format",
			)
		})
		HTTP(func() {
			POST("/add")
			MultipartRequest()
			Response(StatusCreated)
		})
		Result(func() {
			Attribute("success", Boolean)
			Required("success")
		})
	})

})
