package design

import . "goa.design/goa/v3/dsl"

var _ = Service("oAuth", func() {
	Description("Oauth to authentificate")

	Error("unknown_error", unknownError, "Error not identified (500)")
	Error("invalid_scopes", String, "Token scopes are invalid")
	Error("unauthorized", String, "Identifiers are invalid")

	HTTP(func() {
		Response("unknown_error", StatusInternalServerError)
		Response("invalid_scopes", StatusForbidden)
		Response("unauthorized", StatusUnauthorized)
	})

	Method("oAuth", func() {
		Description("oAuth")
		Payload(OauthPayload)
		Result(oAuthResponse)
		HTTP(func() {
			POST("/authorization")
			Headers(func() {
				Param("client_id", String, "The client identifier ID", func() {
					Example("00000")
				})
				Param("client_secret", String, "The client identifier secret", func() {
					Example("99999")
				})
				Param("grant_type", String, "The type of grant", func() {
					Default("client_credentials")
				})
				Required("client_id", "client_secret", "grant_type")
			})
			Response(StatusFound)
			Response(StatusBadRequest)
		})
	})

})

var OauthPayload = Type("OauthPayload", func() {
	Attribute("client_id", String)
	Attribute("client_secret", String)
	Attribute("grant_type", String)
	Required("client_id", "client_secret", "grant_type")
})

var oAuthResponse = Type("oAuthResponse", func() {
	Field(1, "access_token", String)
	Field(2, "token_type", String)
	Field(3, "expires_in", Int64)
	Field(4, "success", Boolean)
})
