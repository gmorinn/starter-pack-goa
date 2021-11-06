// Code generated by goa v3.5.2, DO NOT EDIT.
//
// products HTTP server
//
// Command:
// $ goa gen api_crud/design

package server

import (
	products "api_crud/gen/products"
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	"goa.design/plugins/v3/cors"
)

// Server lists the products service endpoint HTTP handlers.
type Server struct {
	Mounts                   []*MountPoint
	GetAllProductsByCategory http.Handler
	DeleteProduct            http.Handler
	CreateProduct            http.Handler
	UpdateProduct            http.Handler
	GetProduct               http.Handler
	CORS                     http.Handler
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

// New instantiates HTTP handlers for all the products service endpoints using
// the provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *products.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"GetAllProductsByCategory", "GET", "/web/products"},
			{"DeleteProduct", "DELETE", "/web/product/remove/{id}"},
			{"CreateProduct", "POST", "/web/product/add"},
			{"UpdateProduct", "PUT", "/web/product/{id}"},
			{"GetProduct", "GET", "/web/product/{id}"},
			{"CORS", "OPTIONS", "/web/products"},
			{"CORS", "OPTIONS", "/web/product/remove/{id}"},
			{"CORS", "OPTIONS", "/web/product/add"},
			{"CORS", "OPTIONS", "/web/product/{id}"},
		},
		GetAllProductsByCategory: NewGetAllProductsByCategoryHandler(e.GetAllProductsByCategory, mux, decoder, encoder, errhandler, formatter),
		DeleteProduct:            NewDeleteProductHandler(e.DeleteProduct, mux, decoder, encoder, errhandler, formatter),
		CreateProduct:            NewCreateProductHandler(e.CreateProduct, mux, decoder, encoder, errhandler, formatter),
		UpdateProduct:            NewUpdateProductHandler(e.UpdateProduct, mux, decoder, encoder, errhandler, formatter),
		GetProduct:               NewGetProductHandler(e.GetProduct, mux, decoder, encoder, errhandler, formatter),
		CORS:                     NewCORSHandler(),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "products" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.GetAllProductsByCategory = m(s.GetAllProductsByCategory)
	s.DeleteProduct = m(s.DeleteProduct)
	s.CreateProduct = m(s.CreateProduct)
	s.UpdateProduct = m(s.UpdateProduct)
	s.GetProduct = m(s.GetProduct)
	s.CORS = m(s.CORS)
}

// Mount configures the mux to serve the products endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountGetAllProductsByCategoryHandler(mux, h.GetAllProductsByCategory)
	MountDeleteProductHandler(mux, h.DeleteProduct)
	MountCreateProductHandler(mux, h.CreateProduct)
	MountUpdateProductHandler(mux, h.UpdateProduct)
	MountGetProductHandler(mux, h.GetProduct)
	MountCORSHandler(mux, h.CORS)
}

// MountGetAllProductsByCategoryHandler configures the mux to serve the
// "products" service "getAllProductsByCategory" endpoint.
func MountGetAllProductsByCategoryHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleProductsOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/web/products", f)
}

// NewGetAllProductsByCategoryHandler creates a HTTP handler which loads the
// HTTP request and calls the "products" service "getAllProductsByCategory"
// endpoint.
func NewGetAllProductsByCategoryHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeGetAllProductsByCategoryRequest(mux, decoder)
		encodeResponse = EncodeGetAllProductsByCategoryResponse(encoder)
		encodeError    = EncodeGetAllProductsByCategoryError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "getAllProductsByCategory")
		ctx = context.WithValue(ctx, goa.ServiceKey, "products")
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

// MountDeleteProductHandler configures the mux to serve the "products" service
// "deleteProduct" endpoint.
func MountDeleteProductHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleProductsOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("DELETE", "/web/product/remove/{id}", f)
}

// NewDeleteProductHandler creates a HTTP handler which loads the HTTP request
// and calls the "products" service "deleteProduct" endpoint.
func NewDeleteProductHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeDeleteProductRequest(mux, decoder)
		encodeResponse = EncodeDeleteProductResponse(encoder)
		encodeError    = EncodeDeleteProductError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "deleteProduct")
		ctx = context.WithValue(ctx, goa.ServiceKey, "products")
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

// MountCreateProductHandler configures the mux to serve the "products" service
// "createProduct" endpoint.
func MountCreateProductHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleProductsOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/web/product/add", f)
}

// NewCreateProductHandler creates a HTTP handler which loads the HTTP request
// and calls the "products" service "createProduct" endpoint.
func NewCreateProductHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeCreateProductRequest(mux, decoder)
		encodeResponse = EncodeCreateProductResponse(encoder)
		encodeError    = EncodeCreateProductError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "createProduct")
		ctx = context.WithValue(ctx, goa.ServiceKey, "products")
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

// MountUpdateProductHandler configures the mux to serve the "products" service
// "updateProduct" endpoint.
func MountUpdateProductHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleProductsOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("PUT", "/web/product/{id}", f)
}

// NewUpdateProductHandler creates a HTTP handler which loads the HTTP request
// and calls the "products" service "updateProduct" endpoint.
func NewUpdateProductHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeUpdateProductRequest(mux, decoder)
		encodeResponse = EncodeUpdateProductResponse(encoder)
		encodeError    = EncodeUpdateProductError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "updateProduct")
		ctx = context.WithValue(ctx, goa.ServiceKey, "products")
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

// MountGetProductHandler configures the mux to serve the "products" service
// "getProduct" endpoint.
func MountGetProductHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleProductsOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/web/product/{id}", f)
}

// NewGetProductHandler creates a HTTP handler which loads the HTTP request and
// calls the "products" service "getProduct" endpoint.
func NewGetProductHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeGetProductRequest(mux, decoder)
		encodeResponse = EncodeGetProductResponse(encoder)
		encodeError    = EncodeGetProductError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "getProduct")
		ctx = context.WithValue(ctx, goa.ServiceKey, "products")
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
// service products.
func MountCORSHandler(mux goahttp.Muxer, h http.Handler) {
	h = HandleProductsOrigin(h)
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("OPTIONS", "/web/products", f)
	mux.Handle("OPTIONS", "/web/product/remove/{id}", f)
	mux.Handle("OPTIONS", "/web/product/add", f)
	mux.Handle("OPTIONS", "/web/product/{id}", f)
}

// NewCORSHandler creates a HTTP handler which returns a simple 200 response.
func NewCORSHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
}

// HandleProductsOrigin applies the CORS response headers corresponding to the
// origin for the service products.
func HandleProductsOrigin(h http.Handler) http.Handler {
	origHndlr := h.(http.HandlerFunc)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			origHndlr(w, r)
			return
		}
		if cors.MatchOrigin(origin, "http://localhost:3000") {
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