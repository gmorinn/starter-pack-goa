// Code generated by goa v3.5.2, DO NOT EDIT.
//
// boUsers HTTP server encoders and decoders
//
// Command:
// $ goa gen api_crud/design

package server

import (
	bousers "api_crud/gen/bo_users"
	"context"
	"io"
	"net/http"
	"strings"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeGetAllusersResponse returns an encoder for responses returned by the
// boUsers getAllusers endpoint.
func EncodeGetAllusersResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*bousers.GetAllusersResult)
		enc := encoder(ctx, w)
		body := NewGetAllusersResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetAllusersRequest returns a decoder for requests sent to the boUsers
// getAllusers endpoint.
func DecodeGetAllusersRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			oauth    *string
			jwtToken *string
		)
		oauthRaw := r.Header.Get("Authorization")
		if oauthRaw != "" {
			oauth = &oauthRaw
		}
		jwtTokenRaw := r.Header.Get("jwtToken")
		if jwtTokenRaw != "" {
			jwtToken = &jwtTokenRaw
		}
		payload := NewGetAllusersPayload(oauth, jwtToken)
		if payload.Oauth != nil {
			if strings.Contains(*payload.Oauth, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Oauth, " ", 2)[1]
				payload.Oauth = &cred
			}
		}
		if payload.JWTToken != nil {
			if strings.Contains(*payload.JWTToken, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.JWTToken, " ", 2)[1]
				payload.JWTToken = &cred
			}
		}

		return payload, nil
	}
}

// EncodeGetAllusersError returns an encoder for errors returned by the
// getAllusers boUsers endpoint.
func EncodeGetAllusersError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unknown_error":
			res := v.(*bousers.UnknownError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewGetAllusersUnknownErrorResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeDeleteUserResponse returns an encoder for responses returned by the
// boUsers deleteUser endpoint.
func EncodeDeleteUserResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*bousers.DeleteUserResult)
		enc := encoder(ctx, w)
		body := NewDeleteUserResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeDeleteUserRequest returns a decoder for requests sent to the boUsers
// deleteUser endpoint.
func DecodeDeleteUserRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id       string
			oauth    *string
			jwtToken *string
			err      error

			params = mux.Vars(r)
		)
		id = params["id"]
		err = goa.MergeErrors(err, goa.ValidateFormat("id", id, goa.FormatUUID))

		oauthRaw := r.Header.Get("Authorization")
		if oauthRaw != "" {
			oauth = &oauthRaw
		}
		jwtTokenRaw := r.Header.Get("jwtToken")
		if jwtTokenRaw != "" {
			jwtToken = &jwtTokenRaw
		}
		if err != nil {
			return nil, err
		}
		payload := NewDeleteUserPayload(id, oauth, jwtToken)
		if payload.Oauth != nil {
			if strings.Contains(*payload.Oauth, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Oauth, " ", 2)[1]
				payload.Oauth = &cred
			}
		}
		if payload.JWTToken != nil {
			if strings.Contains(*payload.JWTToken, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.JWTToken, " ", 2)[1]
				payload.JWTToken = &cred
			}
		}

		return payload, nil
	}
}

// EncodeDeleteUserError returns an encoder for errors returned by the
// deleteUser boUsers endpoint.
func EncodeDeleteUserError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unknown_error":
			res := v.(*bousers.UnknownError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewDeleteUserUnknownErrorResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeCreateUserResponse returns an encoder for responses returned by the
// boUsers createUser endpoint.
func EncodeCreateUserResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*bousers.CreateUserResult)
		enc := encoder(ctx, w)
		body := NewCreateUserResponseBody(res)
		w.WriteHeader(http.StatusCreated)
		return enc.Encode(body)
	}
}

// DecodeCreateUserRequest returns a decoder for requests sent to the boUsers
// createUser endpoint.
func DecodeCreateUserRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body CreateUserRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateCreateUserRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			oauth    *string
			jwtToken *string
		)
		oauthRaw := r.Header.Get("Authorization")
		if oauthRaw != "" {
			oauth = &oauthRaw
		}
		jwtTokenRaw := r.Header.Get("jwtToken")
		if jwtTokenRaw != "" {
			jwtToken = &jwtTokenRaw
		}
		payload := NewCreateUserPayload(&body, oauth, jwtToken)
		if payload.Oauth != nil {
			if strings.Contains(*payload.Oauth, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Oauth, " ", 2)[1]
				payload.Oauth = &cred
			}
		}
		if payload.JWTToken != nil {
			if strings.Contains(*payload.JWTToken, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.JWTToken, " ", 2)[1]
				payload.JWTToken = &cred
			}
		}

		return payload, nil
	}
}

// EncodeCreateUserError returns an encoder for errors returned by the
// createUser boUsers endpoint.
func EncodeCreateUserError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unknown_error":
			res := v.(*bousers.UnknownError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewCreateUserUnknownErrorResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeUpdateUserResponse returns an encoder for responses returned by the
// boUsers updateUser endpoint.
func EncodeUpdateUserResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*bousers.UpdateUserResult)
		enc := encoder(ctx, w)
		body := NewUpdateUserResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeUpdateUserRequest returns a decoder for requests sent to the boUsers
// updateUser endpoint.
func DecodeUpdateUserRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body UpdateUserRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateUpdateUserRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			id       string
			oauth    *string
			jwtToken *string

			params = mux.Vars(r)
		)
		id = params["id"]
		err = goa.MergeErrors(err, goa.ValidateFormat("id", id, goa.FormatUUID))

		oauthRaw := r.Header.Get("Authorization")
		if oauthRaw != "" {
			oauth = &oauthRaw
		}
		jwtTokenRaw := r.Header.Get("jwtToken")
		if jwtTokenRaw != "" {
			jwtToken = &jwtTokenRaw
		}
		if err != nil {
			return nil, err
		}
		payload := NewUpdateUserPayload(&body, id, oauth, jwtToken)
		if payload.Oauth != nil {
			if strings.Contains(*payload.Oauth, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Oauth, " ", 2)[1]
				payload.Oauth = &cred
			}
		}
		if payload.JWTToken != nil {
			if strings.Contains(*payload.JWTToken, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.JWTToken, " ", 2)[1]
				payload.JWTToken = &cred
			}
		}

		return payload, nil
	}
}

// EncodeUpdateUserError returns an encoder for errors returned by the
// updateUser boUsers endpoint.
func EncodeUpdateUserError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unknown_error":
			res := v.(*bousers.UnknownError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewUpdateUserUnknownErrorResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeGetUserResponse returns an encoder for responses returned by the
// boUsers getUser endpoint.
func EncodeGetUserResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*bousers.GetUserResult)
		enc := encoder(ctx, w)
		body := NewGetUserResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetUserRequest returns a decoder for requests sent to the boUsers
// getUser endpoint.
func DecodeGetUserRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id       string
			oauth    *string
			jwtToken *string
			err      error

			params = mux.Vars(r)
		)
		id = params["id"]
		err = goa.MergeErrors(err, goa.ValidateFormat("id", id, goa.FormatUUID))

		oauthRaw := r.Header.Get("Authorization")
		if oauthRaw != "" {
			oauth = &oauthRaw
		}
		jwtTokenRaw := r.Header.Get("jwtToken")
		if jwtTokenRaw != "" {
			jwtToken = &jwtTokenRaw
		}
		if err != nil {
			return nil, err
		}
		payload := NewGetUserPayload(id, oauth, jwtToken)
		if payload.Oauth != nil {
			if strings.Contains(*payload.Oauth, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Oauth, " ", 2)[1]
				payload.Oauth = &cred
			}
		}
		if payload.JWTToken != nil {
			if strings.Contains(*payload.JWTToken, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.JWTToken, " ", 2)[1]
				payload.JWTToken = &cred
			}
		}

		return payload, nil
	}
}

// EncodeGetUserError returns an encoder for errors returned by the getUser
// boUsers endpoint.
func EncodeGetUserError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unknown_error":
			res := v.(*bousers.UnknownError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewGetUserUnknownErrorResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeDeleteManyUsersResponse returns an encoder for responses returned by
// the boUsers deleteManyUsers endpoint.
func EncodeDeleteManyUsersResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*bousers.DeleteManyUsersResult)
		enc := encoder(ctx, w)
		body := NewDeleteManyUsersResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeDeleteManyUsersRequest returns a decoder for requests sent to the
// boUsers deleteManyUsers endpoint.
func DecodeDeleteManyUsersRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body DeleteManyUsersRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateDeleteManyUsersRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			oauth    *string
			jwtToken *string
		)
		oauthRaw := r.Header.Get("Authorization")
		if oauthRaw != "" {
			oauth = &oauthRaw
		}
		jwtTokenRaw := r.Header.Get("jwtToken")
		if jwtTokenRaw != "" {
			jwtToken = &jwtTokenRaw
		}
		payload := NewDeleteManyUsersPayload(&body, oauth, jwtToken)
		if payload.Oauth != nil {
			if strings.Contains(*payload.Oauth, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Oauth, " ", 2)[1]
				payload.Oauth = &cred
			}
		}
		if payload.JWTToken != nil {
			if strings.Contains(*payload.JWTToken, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.JWTToken, " ", 2)[1]
				payload.JWTToken = &cred
			}
		}

		return payload, nil
	}
}

// EncodeDeleteManyUsersError returns an encoder for errors returned by the
// deleteManyUsers boUsers endpoint.
func EncodeDeleteManyUsersError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unknown_error":
			res := v.(*bousers.UnknownError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewDeleteManyUsersUnknownErrorResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeNewPasswordResponse returns an encoder for responses returned by the
// boUsers newPassword endpoint.
func EncodeNewPasswordResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*bousers.NewPasswordResult)
		enc := encoder(ctx, w)
		body := NewNewPasswordResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeNewPasswordRequest returns a decoder for requests sent to the boUsers
// newPassword endpoint.
func DecodeNewPasswordRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body NewPasswordRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateNewPasswordRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			id       string
			oauth    *string
			jwtToken *string

			params = mux.Vars(r)
		)
		id = params["id"]
		err = goa.MergeErrors(err, goa.ValidateFormat("id", id, goa.FormatUUID))

		oauthRaw := r.Header.Get("Authorization")
		if oauthRaw != "" {
			oauth = &oauthRaw
		}
		jwtTokenRaw := r.Header.Get("jwtToken")
		if jwtTokenRaw != "" {
			jwtToken = &jwtTokenRaw
		}
		if err != nil {
			return nil, err
		}
		payload := NewNewPasswordPayload(&body, id, oauth, jwtToken)
		if payload.Oauth != nil {
			if strings.Contains(*payload.Oauth, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Oauth, " ", 2)[1]
				payload.Oauth = &cred
			}
		}
		if payload.JWTToken != nil {
			if strings.Contains(*payload.JWTToken, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.JWTToken, " ", 2)[1]
				payload.JWTToken = &cred
			}
		}

		return payload, nil
	}
}

// EncodeNewPasswordError returns an encoder for errors returned by the
// newPassword boUsers endpoint.
func EncodeNewPasswordError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unknown_error":
			res := v.(*bousers.UnknownError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewNewPasswordUnknownErrorResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// marshalBousersResBoUserToResBoUserResponseBody builds a value of type
// *ResBoUserResponseBody from a value of type *bousers.ResBoUser.
func marshalBousersResBoUserToResBoUserResponseBody(v *bousers.ResBoUser) *ResBoUserResponseBody {
	res := &ResBoUserResponseBody{
		ID:        v.ID,
		Firstname: v.Firstname,
		Lastname:  v.Lastname,
		Email:     v.Email,
		Birthday:  v.Birthday,
		Phone:     v.Phone,
		Role:      v.Role,
	}
	{
		var zero string
		if res.Birthday == zero {
			res.Birthday = ""
		}
	}
	{
		var zero string
		if res.Phone == zero {
			res.Phone = ""
		}
	}
	{
		var zero string
		if res.Role == zero {
			res.Role = "user"
		}
	}

	return res
}

// unmarshalPayloadUserRequestBodyToBousersPayloadUser builds a value of type
// *bousers.PayloadUser from a value of type *PayloadUserRequestBody.
func unmarshalPayloadUserRequestBodyToBousersPayloadUser(v *PayloadUserRequestBody) *bousers.PayloadUser {
	res := &bousers.PayloadUser{
		Firstname: *v.Firstname,
		Lastname:  *v.Lastname,
		Email:     *v.Email,
	}
	if v.Birthday != nil {
		res.Birthday = *v.Birthday
	}
	if v.Role != nil {
		res.Role = *v.Role
	}
	if v.Phone != nil {
		res.Phone = *v.Phone
	}
	if v.Birthday == nil {
		res.Birthday = ""
	}
	if v.Role == nil {
		res.Role = "user"
	}
	if v.Phone == nil {
		res.Phone = ""
	}

	return res
}
