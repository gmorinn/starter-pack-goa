// Code generated by goa v3.5.2, DO NOT EDIT.
//
// boUsers HTTP server
//
// Command:
// $ goa gen api_crud/design

package server

import (
	bousers "api_crud/gen/bo_users"
	"context"
	"net/http"
	"regexp"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	"goa.design/plugins/v3/cors"
)

// Server lists the boUsers service endpoint HTTP handlers.
type Server struct {
	Mounts          []*MountPoint
	GetAllusers     http.Handler
	DeleteUser      http.Handler
	CreateUser      http.Handler
	UpdateUser      http.Handler
	GetUser         http.Handler
	DeleteManyUsers http.Handler
	NewPassword     http.Handler
	CORS            http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the boUsers service endpoints using
// the provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *bousers.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"GetAllusers", "GET", "/v1/bo/users/{offset}/{limit}"},
			{"DeleteUser", "DELETE", "/v1/bo/user/remove/{id}"},
			{"CreateUser", "POST", "/v1/bo/user/add"},
			{"UpdateUser", "PUT", "/v1/bo/user/{id}"},
			{"GetUser", "GET", "/v1/bo/user/{id}"},
			{"DeleteManyUsers", "PATCH", "/v1/bo/users/remove"},
			{"NewPassword", "PATCH", "/v1/bo/user/change/password/{id}"},
			{"CORS", "OPTIONS", "/v1/bo/users/{offset}/{limit}"},
			{"CORS", "OPTIONS", "/v1/bo/user/remove/{id}"},
			{"CORS", "OPTIONS", "/v1/bo/user/add"},
			{"CORS", "OPTIONS", "/v1/bo/user/{id}"},
			{"CORS", "OPTIONS", "/v1/bo/users/remove"},
			{"CORS", "OPTIONS", "/v1/bo/user/change/password/{id}"},
		},
		GetAllusers:     NewGetAllusersHandler(e.GetAllusers, mux, decoder, encoder, errhandler, formatter),
		DeleteUser:      NewDeleteUserHandler(e.DeleteUser, mux, decoder, encoder, errhandler, formatter),
		CreateUser:      NewCreateUserHandler(e.CreateUser, mux, decoder, encoder, errhandler, formatter),
		UpdateUser:      NewUpdateUserHandler(e.UpdateUser, mux, decoder, encoder, errhandler, formatter),
		GetUser:         NewGetUserHandler(e.GetUser, mux, decoder, encoder, errhandler, formatter),
		DeleteManyUsers: NewDeleteManyUsersHandler(e.DeleteManyUsers, mux, decoder, encoder, errhandler, formatter),
		NewPassword:     NewNewPasswordHandler(e.NewPassword, mux, decoder, encoder, errhandler, formatter),
		CORS:            NewCORSHandler(),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "boUsers" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.GetAllusers = m(s.GetAllusers)
	s.DeleteUser = m(s.DeleteUser)
	s.CreateUser = m(s.CreateUser)
	s.UpdateUser = m(s.UpdateUser)
	s.GetUser = m(s.GetUser)
	s.DeleteManyUsers = m(s.DeleteManyUsers)
	s.NewPassword = m(s.NewPassword)
	s.CORS = m(s.CORS)
}

// Mount configures the mux to serve the boUsers endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountGetAllusersHandler(mux, h.GetAllusers)
	MountDeleteUserHandler(mux, h.DeleteUser)
	MountCreateUserHandler(mux, h.CreateUser)
	MountUpdateUserHandler(mux, h.UpdateUser)
	MountGetUserHandler(mux, h.GetUser)
	MountDeleteManyUsersHandler(mux, h.DeleteManyUsers)
	MountNewPasswordHandler(mux, h.NewPassword)
	MountCORSHandler(mux, h.CORS)
}

// MountGetAllusersHandler configures the mux to serve the "boUsers" service
// "getAllusers" endpoint.
func MountGetAllusersHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleBoUsersOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/v1/bo/users/{offset}/{limit}", f)
}

// NewGetAllusersHandler creates a HTTP handler which loads the HTTP request
// and calls the "boUsers" service "getAllusers" endpoint.
func NewGetAllusersHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeGetAllusersRequest(mux, decoder)
		encodeResponse = EncodeGetAllusersResponse(encoder)
		encodeError    = EncodeGetAllusersError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "getAllusers")
		ctx = context.WithValue(ctx, goa.ServiceKey, "boUsers")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountDeleteUserHandler configures the mux to serve the "boUsers" service
// "deleteUser" endpoint.
func MountDeleteUserHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleBoUsersOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("DELETE", "/v1/bo/user/remove/{id}", f)
}

// NewDeleteUserHandler creates a HTTP handler which loads the HTTP request and
// calls the "boUsers" service "deleteUser" endpoint.
func NewDeleteUserHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeDeleteUserRequest(mux, decoder)
		encodeResponse = EncodeDeleteUserResponse(encoder)
		encodeError    = EncodeDeleteUserError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "deleteUser")
		ctx = context.WithValue(ctx, goa.ServiceKey, "boUsers")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountCreateUserHandler configures the mux to serve the "boUsers" service
// "createUser" endpoint.
func MountCreateUserHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleBoUsersOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/v1/bo/user/add", f)
}

// NewCreateUserHandler creates a HTTP handler which loads the HTTP request and
// calls the "boUsers" service "createUser" endpoint.
func NewCreateUserHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeCreateUserRequest(mux, decoder)
		encodeResponse = EncodeCreateUserResponse(encoder)
		encodeError    = EncodeCreateUserError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "createUser")
		ctx = context.WithValue(ctx, goa.ServiceKey, "boUsers")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountUpdateUserHandler configures the mux to serve the "boUsers" service
// "updateUser" endpoint.
func MountUpdateUserHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleBoUsersOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("PUT", "/v1/bo/user/{id}", f)
}

// NewUpdateUserHandler creates a HTTP handler which loads the HTTP request and
// calls the "boUsers" service "updateUser" endpoint.
func NewUpdateUserHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeUpdateUserRequest(mux, decoder)
		encodeResponse = EncodeUpdateUserResponse(encoder)
		encodeError    = EncodeUpdateUserError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "updateUser")
		ctx = context.WithValue(ctx, goa.ServiceKey, "boUsers")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountGetUserHandler configures the mux to serve the "boUsers" service
// "getUser" endpoint.
func MountGetUserHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleBoUsersOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/v1/bo/user/{id}", f)
}

// NewGetUserHandler creates a HTTP handler which loads the HTTP request and
// calls the "boUsers" service "getUser" endpoint.
func NewGetUserHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeGetUserRequest(mux, decoder)
		encodeResponse = EncodeGetUserResponse(encoder)
		encodeError    = EncodeGetUserError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "getUser")
		ctx = context.WithValue(ctx, goa.ServiceKey, "boUsers")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountDeleteManyUsersHandler configures the mux to serve the "boUsers"
// service "deleteManyUsers" endpoint.
func MountDeleteManyUsersHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleBoUsersOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("PATCH", "/v1/bo/users/remove", f)
}

// NewDeleteManyUsersHandler creates a HTTP handler which loads the HTTP
// request and calls the "boUsers" service "deleteManyUsers" endpoint.
func NewDeleteManyUsersHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeDeleteManyUsersRequest(mux, decoder)
		encodeResponse = EncodeDeleteManyUsersResponse(encoder)
		encodeError    = EncodeDeleteManyUsersError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "deleteManyUsers")
		ctx = context.WithValue(ctx, goa.ServiceKey, "boUsers")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountNewPasswordHandler configures the mux to serve the "boUsers" service
// "newPassword" endpoint.
func MountNewPasswordHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleBoUsersOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("PATCH", "/v1/bo/user/change/password/{id}", f)
}

// NewNewPasswordHandler creates a HTTP handler which loads the HTTP request
// and calls the "boUsers" service "newPassword" endpoint.
func NewNewPasswordHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeNewPasswordRequest(mux, decoder)
		encodeResponse = EncodeNewPasswordResponse(encoder)
		encodeError    = EncodeNewPasswordError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "newPassword")
		ctx = context.WithValue(ctx, goa.ServiceKey, "boUsers")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountCORSHandler configures the mux to serve the CORS endpoints for the
// service boUsers.
func MountCORSHandler(mux goahttp.Muxer, h http.Handler) {
	h = HandleBoUsersOrigin(h)
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("OPTIONS", "/v1/bo/users/{offset}/{limit}", f)
	mux.Handle("OPTIONS", "/v1/bo/user/remove/{id}", f)
	mux.Handle("OPTIONS", "/v1/bo/user/add", f)
	mux.Handle("OPTIONS", "/v1/bo/user/{id}", f)
	mux.Handle("OPTIONS", "/v1/bo/users/remove", f)
	mux.Handle("OPTIONS", "/v1/bo/user/change/password/{id}", f)
}

// NewCORSHandler creates a HTTP handler which returns a simple 200 response.
func NewCORSHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
}

// HandleBoUsersOrigin applies the CORS response headers corresponding to the
// origin for the service boUsers.
func HandleBoUsersOrigin(h http.Handler) http.Handler {
	spec0 := regexp.MustCompile(".*localhost.*")
	origHndlr := h.(http.HandlerFunc)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			origHndlr(w, r)
			return
		}
		if cors.MatchOriginRegexp(origin, spec0) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Expose-Headers", "Content-Type, Origin")
			w.Header().Set("Access-Control-Max-Age", "600")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS, DELETE, PATCH")
				w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, jwtToken, Origin")
			}
			origHndlr(w, r)
			return
		}
		origHndlr(w, r)
		return
	})
}
