package design

import . "goa.design/goa/v3/dsl"

var OAuth2 = OAuth2Security("OAuth2", func() {
	Description("Use OAuth2 to authenticate")
	ClientCredentialsFlow("/authorization", "/refresh")
	Scope("api:read", "Provides read access")
})

var _ = Service("oAuth", func() {
	Description("Oauth to authentificate")

	Error("unknown_error", unknownError, "Error not identified (500)")
	Error("invalid_scopes", String, "Token scopes are invalid")
	Error("unauthorized", String, "Identifiers are invalid")
	Error("oauth_error", String, "Error when a request is send before asking for oAuth")

	HTTP(func() {
		Response("unknown_error", StatusInternalServerError)
		Response("invalid_scopes", StatusForbidden)
		Response("oauth_error", StatusForbidden)
		Response("unauthorized", StatusUnauthorized)
	})

	Method("oAuth", func() {
		Description("oAuth")
		Payload(OauthPayload)
		Result(oAuthResponse)
		HTTP(func() {
			POST("/authorization")
			Response(StatusCreated)
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
	Required("access_token", "token_type", "expires_in", "success")
})
