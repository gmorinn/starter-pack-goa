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
			Attribute("filename", String, "uploaded file name", func() {
				Example("foo.jpg")
			})
			Attribute("url", String, "url file")
			Attribute("mime", String, "url file")
			Attribute("content", Bytes, "content of image")
			Attribute("size", Int64, "size of image")
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
				"filename",
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
			Attribute("file", resFile)
			Attribute("success", Boolean)
			Required("success", "file")
		})
	})

	// Method("deleteFile", func() {
	// 	Description("Delete one file by ID")
	// 	Payload(func() {
	// 		Attribute("id", String, func() {
	// 			Format(FormatUUID)
	// 			Example("5dfb0bf7-597a-4250-b7ad-63a43ff59c25")
	// 		})
	// 		TokenField(1, "jwtToken", String, func() {
	// 			Description("JWT used for authentication after Signin/Signup")
	// 		})
	// 		AccessTokenField(2, "oauth", String, func() {
	// 			Description("Use to generate Oauth with /authorization")
	// 		})
	// 		Required("id")
	// 	})
	// 	HTTP(func() {
	// 		DELETE("/remove/{id}")
	// 		Response(StatusOK)
	// 	})
	// 	Result(func() {
	// 		Attribute("success", Boolean)
	// 		Required("success")
	// 	})
	// })
})

var resFile = Type("resFile", func() {
	Attribute("id", String, func() {
		Format(FormatUUID)
		Example("5dfb0bf7-597a-4250-b7ad-63a43ff59c25")
	})
	Attribute("name", String, func() {
		Example("foo.png")
	})
	Attribute("url", String)
	Attribute("mime", String)
	Attribute("size", Int64)
	Required("id", "name", "url")
})
