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

	Security(OAuth2)

	Error("unknown_error", unknownError, "Error not identified (500)")

	HTTP(func() {
		Path("/v1/file")
		Header("oauth:Authorization", String, "OAuth token", func() {
			Pattern("^Bearer [^ ]+$")
		})
		Response("unknown_error", StatusInternalServerError)
	})

	Method("importFile", func() {
		Description("Import file")
		Payload(func() {
			Attribute("files", ArrayOf(payloadFile), "Files to import")
			AccessTokenField(1, "oauth", String, func() {
				Description("Use to generate Oauth with /authorization")
			})
			Required(
				"files",
			)
		})
		HTTP(func() {
			POST("/add")
			MultipartRequest()
			Response(StatusCreated)
		})
		Result(func() {
			Attribute("file", ArrayOf(resFile))
			Attribute("success", Boolean)
			Required("success", "file")
		})
	})

	Method("deleteFile", func() {
		Description("Delete one file by ID")
		Payload(func() {
			Attribute("url", String, func() {
				Example("/public/uploads/2021/12/2ca51d10-b660-4b2c-b27f-f7a119642885.png")
				MinLength(23)
			})
			AccessTokenField(1, "oauth", String, func() {
				Description("Use to generate Oauth with /authorization")
			})
			Required("url")
		})
		HTTP(func() {
			PATCH("/remove")
			Response(StatusOK)
		})
		Result(func() {
			Attribute("success", Boolean)
			Required("success")
		})
	})
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

var payloadFile = Type("payloadFile", func() {
	Attribute("filename", String, "uploaded file name", func() {
		Example("foo.jpg")
		MinLength(2)
	})
	Attribute("url", String, "url file")
	Attribute("w", Int64, "width of image if you crop")
	Attribute("h", Int64, "height of image if you crop")
	Attribute("content", Bytes, "content of image")
	Attribute("size", Int64, "size of image")
	Attribute("format", String, "uploaded file format", func() {
		Example("image/jpeg")
		Enum("image/jpeg", "image/png", "image/jpg")
	})
	Required(
		"filename",
		"content",
		"format",
		"size",
		"url",
	)
})
