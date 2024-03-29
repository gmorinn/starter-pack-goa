// Code generated by goa v3.10.2, DO NOT EDIT.
//
// auth HTTP client encoders and decoders
//
// Command:
// $ goa gen starter-pack-goa/design

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	auth "starter-pack-goa/gen/auth"
	"strings"

	goahttp "goa.design/goa/v3/http"
)

// BuildEmailExistRequest instantiates a HTTP request object with method and
// path set to call the "auth" service "email-exist" endpoint
func (c *Client) BuildEmailExistRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: EmailExistAuthPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("auth", "email-exist", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeEmailExistRequest returns an encoder for requests sent to the auth
// email-exist server.
func EncodeEmailExistRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*auth.EmailExistPayload)
		if !ok {
			return goahttp.ErrInvalidType("auth", "email-exist", "*auth.EmailExistPayload", v)
		}
		if p.Oauth != nil {
			head := *p.Oauth
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		body := NewEmailExistRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("auth", "email-exist", err)
		}
		return nil
	}
}

// DecodeEmailExistResponse returns a decoder for responses returned by the
// auth email-exist endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeEmailExistResponse may return the following errors:
//   - "unknown_error" (type *auth.UnknownError): http.StatusInternalServerError
//   - error: internal error
func DecodeEmailExistResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body EmailExistResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("auth", "email-exist", err)
			}
			err = ValidateEmailExistResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("auth", "email-exist", err)
			}
			res := NewEmailExistResultOK(&body)
			return res, nil
		case http.StatusInternalServerError:
			var (
				body EmailExistUnknownErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("auth", "email-exist", err)
			}
			err = ValidateEmailExistUnknownErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("auth", "email-exist", err)
			}
			return nil, NewEmailExistUnknownError(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("auth", "email-exist", resp.StatusCode, string(body))
		}
	}
}

// BuildSendConfirmationRequest instantiates a HTTP request object with method
// and path set to call the "auth" service "send-confirmation" endpoint
func (c *Client) BuildSendConfirmationRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: SendConfirmationAuthPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("auth", "send-confirmation", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeSendConfirmationRequest returns an encoder for requests sent to the
// auth send-confirmation server.
func EncodeSendConfirmationRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*auth.SendConfirmationPayload)
		if !ok {
			return goahttp.ErrInvalidType("auth", "send-confirmation", "*auth.SendConfirmationPayload", v)
		}
		if p.Oauth != nil {
			head := *p.Oauth
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		body := NewSendConfirmationRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("auth", "send-confirmation", err)
		}
		return nil
	}
}

// DecodeSendConfirmationResponse returns a decoder for responses returned by
// the auth send-confirmation endpoint. restoreBody controls whether the
// response body should be restored after having been read.
// DecodeSendConfirmationResponse may return the following errors:
//   - "unknown_error" (type *auth.UnknownError): http.StatusInternalServerError
//   - error: internal error
func DecodeSendConfirmationResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body SendConfirmationResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("auth", "send-confirmation", err)
			}
			err = ValidateSendConfirmationResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("auth", "send-confirmation", err)
			}
			res := NewSendConfirmationResultOK(&body)
			return res, nil
		case http.StatusInternalServerError:
			var (
				body SendConfirmationUnknownErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("auth", "send-confirmation", err)
			}
			err = ValidateSendConfirmationUnknownErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("auth", "send-confirmation", err)
			}
			return nil, NewSendConfirmationUnknownError(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("auth", "send-confirmation", resp.StatusCode, string(body))
		}
	}
}

// BuildResetPasswordRequest instantiates a HTTP request object with method and
// path set to call the "auth" service "reset-password" endpoint
func (c *Client) BuildResetPasswordRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ResetPasswordAuthPath()}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("auth", "reset-password", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeResetPasswordRequest returns an encoder for requests sent to the auth
// reset-password server.
func EncodeResetPasswordRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*auth.ResetPasswordPayload)
		if !ok {
			return goahttp.ErrInvalidType("auth", "reset-password", "*auth.ResetPasswordPayload", v)
		}
		if p.Oauth != nil {
			head := *p.Oauth
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		body := NewResetPasswordRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("auth", "reset-password", err)
		}
		return nil
	}
}

// DecodeResetPasswordResponse returns a decoder for responses returned by the
// auth reset-password endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeResetPasswordResponse may return the following errors:
//   - "unknown_error" (type *auth.UnknownError): http.StatusInternalServerError
//   - error: internal error
func DecodeResetPasswordResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ResetPasswordResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("auth", "reset-password", err)
			}
			err = ValidateResetPasswordResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("auth", "reset-password", err)
			}
			res := NewResetPasswordResultOK(&body)
			return res, nil
		case http.StatusInternalServerError:
			var (
				body ResetPasswordUnknownErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("auth", "reset-password", err)
			}
			err = ValidateResetPasswordUnknownErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("auth", "reset-password", err)
			}
			return nil, NewResetPasswordUnknownError(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("auth", "reset-password", resp.StatusCode, string(body))
		}
	}
}
