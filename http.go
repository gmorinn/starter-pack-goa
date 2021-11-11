package main

import (
	jwttokensvr "api_crud/gen/http/jwt_token/server"
	oauthsvr "api_crud/gen/http/o_auth/server"
	openapisvr "api_crud/gen/http/openapi/server"
	productssvr "api_crud/gen/http/products/server"
	"context"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"

	goahttp "goa.design/goa/v3/http"
	httpmdlwr "goa.design/goa/v3/http/middleware"
	"goa.design/goa/v3/middleware"
)

// handleHTTPServer starts configures and starts a HTTP server on the given
// URL. It shuts down the server if any error is received in the error channel.
func handleHTTPServer(ctx context.Context, u *url.URL, api *ApiEndpoints, wg *sync.WaitGroup, errc chan error, logger *log.Logger, debug bool) {

	// Setup goa log adapter.
	var (
		adapter middleware.Logger
	)
	{
		adapter = middleware.NewLogger(logger)
	}

	// Provide the transport specific request decoder and response encoder.
	// The goa http package has built-in support for JSON, XML and gob.
	// Other encodings can be used by providing the corresponding functions,
	// see goa.design/implement/encoding.
	var (
		dec = goahttp.RequestDecoder
		enc = goahttp.ResponseEncoder
	)

	// Build the service HTTP request multiplexer and configure it to serve
	// HTTP requests to the service endpoints.
	var mux goahttp.Muxer
	{
		mux = goahttp.NewMuxer()
	}

	// Wrap the endpoints with the transport specific layers. The generated
	// server packages contains code generated from the design which maps
	// the service input and output data structures to HTTP requests and
	// responses.
	var (
		openapiServer  *openapisvr.Server
		jwtTokenServer *jwttokensvr.Server
		oAuthServer    *oauthsvr.Server
		productsServer *productssvr.Server
	)
	{
		eh := errorHandler(logger)
		openapiServer = openapisvr.New(nil, mux, dec, enc, nil, nil, http.Dir("../../gen/http"))
		jwtTokenServer = jwttokensvr.New(api.jwtTokenEndpoints, mux, dec, enc, eh, nil)
		oAuthServer = oauthsvr.New(api.oAuthEndpoints, mux, dec, enc, eh, nil)
		productsServer = productssvr.New(api.productsEndpoints, mux, dec, enc, eh, nil)
		if debug {
			servers := goahttp.Servers{
				openapiServer,
				jwtTokenServer,
				oAuthServer,
				productsServer,
			}
			servers.Use(httpmdlwr.Debug(mux, os.Stdout))
		}
	}
	// Configure the mux.
	openapisvr.Mount(mux, openapiServer)
	jwttokensvr.Mount(mux, jwtTokenServer)
	oauthsvr.Mount(mux, oAuthServer)
	productssvr.Mount(mux, productsServer)

	// Wrap the multiplexer with additional middlewares. Middlewares mounted
	// here apply to all the service endpoints.
	var handler http.Handler = mux
	{
		handler = httpmdlwr.Log(adapter)(handler)
		handler = httpmdlwr.RequestID()(handler)
	}

	// Start HTTP server using default configuration, change the code to
	// configure the server as required by your service.
	srv := &http.Server{Addr: u.Host, Handler: handler}
	logger.Printf(`
	
	 ██████╗  ██████╗  █████╗     ██╗  ██╗     ██████╗ ███╗   ███╗
	██╔════╝ ██╔═══██╗██╔══██╗    ╚██╗██╔╝    ██╔════╝ ████╗ ████║
	██║  ███╗██║   ██║███████║     ╚███╔╝     ██║  ███╗██╔████╔██║
	██║   ██║██║   ██║██╔══██║     ██╔██╗     ██║   ██║██║╚██╔╝██║
	╚██████╔╝╚██████╔╝██║  ██║    ██╔╝ ██╗    ╚██████╔╝██║ ╚═╝ ██║
	 ╚═════╝  ╚═════╝ ╚═╝  ╚═╝    ╚═╝  ╚═╝     ╚═════╝ ╚═╝     ╚═╝
																  
																  `)

	(*wg).Add(1)
	go func() {
		defer (*wg).Done()

		// Start HTTP server in a separate goroutine.
		go func() {
			logger.Printf("HTTP server listening on %q", u.Host)
			errc <- srv.ListenAndServe()
		}()

		<-ctx.Done()
		logger.Printf("shutting down HTTP server at %q", u.Host)

		// Shutdown gracefully with a 30s timeout.
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		_ = srv.Shutdown(ctx)
	}()
}

// errorHandler returns a function that writes and logs the given error.
// The function also writes and logs the error unique ID so that it's possible
// to correlate.
func errorHandler(logger *log.Logger) func(context.Context, http.ResponseWriter, error) {
	return func(ctx context.Context, w http.ResponseWriter, err error) {
		id := ctx.Value(middleware.RequestIDKey).(string)
		_, _ = w.Write([]byte("[" + id + "] encoding: " + err.Error()))
		logger.Printf("[%s] ERROR: %s", id, err.Error())
	}
}
