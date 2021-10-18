// Code generated by goa v3.5.2, DO NOT EDIT.
//
// oAuth HTTP client encoders and decoders
//
// Command:
// $ goa gen api_crud/design

package client

import (
	oauth "api_crud/gen/o_auth"
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	goahttp "goa.design/goa/v3/http"
)

// BuildOAuthRequest instantiates a HTTP request object with method and path
// set to call the "oAuth" service "oAuth" endpoint
func (c *Client) BuildOAuthRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: OAuthOAuthPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("oAuth", "oAuth", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeOAuthRequest returns an encoder for requests sent to the oAuth oAuth
// server.
func EncodeOAuthRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*oauth.OauthPayload)
		if !ok {
			return goahttp.ErrInvalidType("oAuth", "oAuth", "*oauth.OauthPayload", v)
		}
		{
			head := p.ClientID
			req.Header.Set("client_id", head)
		}
		{
			head := p.ClientSecret
			req.Header.Set("client_secret", head)
		}
		{
			head := p.GrantType
			req.Header.Set("grant_type", head)
		}
		return nil
	}
}

// DecodeOAuthResponse returns a decoder for responses returned by the oAuth
// oAuth endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeOAuthResponse may return the following errors:
//	- "unknown_error" (type *oauth.UnknownError): http.StatusInternalServerError
//	- "invalid_scopes" (type oauth.InvalidScopes): http.StatusForbidden
//	- "unauthorized" (type oauth.Unauthorized): http.StatusUnauthorized
//	- error: internal error
func DecodeOAuthResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusFound:
			var (
				body OAuthFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("oAuth", "oAuth", err)
			}
			res := NewOAuthResponseFound(&body)
			return res, nil
		case http.StatusInternalServerError:
			var (
				body OAuthUnknownErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("oAuth", "oAuth", err)
			}
			err = ValidateOAuthUnknownErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("oAuth", "oAuth", err)
			}
			return nil, NewOAuthUnknownError(&body)
		case http.StatusForbidden:
			var (
				body OAuthInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("oAuth", "oAuth", err)
			}
			return nil, NewOAuthInvalidScopes(body)
		case http.StatusUnauthorized:
			var (
				body OAuthUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("oAuth", "oAuth", err)
			}
			return nil, NewOAuthUnauthorized(body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("oAuth", "oAuth", resp.StatusCode, string(body))
		}
	}
}
