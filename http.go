package main

import (
	filesclient "api_crud/api"
	"api_crud/config"
	authsvr "api_crud/gen/http/auth/server"
	boProductssrvc "api_crud/gen/http/bo_products/server"
	bouserssrv "api_crud/gen/http/bo_users/server"
	fileapisvr "api_crud/gen/http/fileapi/server"
	filessvr "api_crud/gen/http/files/server"
	jwttokensvr "api_crud/gen/http/jwt_token/server"
	oauthsvr "api_crud/gen/http/o_auth/server"
	openapisvr "api_crud/gen/http/openapi/server"
	productssvr "api_crud/gen/http/products/server"
	userssvr "api_crud/gen/http/users/server"
	"context"
	"fmt"
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

	// Get .env
	cnf := config.New()

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
		eh                                       = errorHandler(logger)
		openapiServer     *openapisvr.Server     = openapisvr.New(nil, mux, dec, enc, nil, nil, http.Dir("gen/http"))
		filesServer       *filessvr.Server       = filessvr.New(api.filesEndpoints, mux, dec, enc, eh, nil, filesclient.FilesImportFileDecoderFunc)
		bo_productsServer *boProductssrvc.Server = boProductssrvc.New(api.bo_productsEndpoints, mux, dec, enc, eh, nil)
		bo_usersServer    *bouserssrv.Server     = bouserssrv.New(api.bo_usersEndpoints, mux, dec, enc, eh, nil)
		usersServer       *userssvr.Server       = userssvr.New(api.usersEndpoints, mux, dec, enc, eh, nil)
		jwtTokenServer    *jwttokensvr.Server    = jwttokensvr.New(api.jwtTokenEndpoints, mux, dec, enc, eh, nil)
		fileapiServer     *fileapisvr.Server     = fileapisvr.New(nil, mux, dec, enc, nil, nil, http.Dir("cmd"))
		oAuthServer       *oauthsvr.Server       = oauthsvr.New(api.oAuthEndpoints, mux, dec, enc, eh, nil)
		authServer        *authsvr.Server        = authsvr.New(api.authEndpoints, mux, dec, enc, eh, nil)
		productsServer    *productssvr.Server    = productssvr.New(api.productsEndpoints, mux, dec, enc, eh, nil)
	)
	{
		if debug {
			servers := goahttp.Servers{
				fileapiServer,
				filesServer,
				authServer,
				bo_productsServer,
				bo_usersServer,
				usersServer,
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
	fileapisvr.Mount(mux, fileapiServer)
	filessvr.Mount(mux, filesServer)
	authsvr.Mount(mux, authServer)
	boProductssrvc.Mount(mux, bo_productsServer)
	bouserssrv.Mount(mux, bo_usersServer)
	userssvr.Mount(mux, usersServer)
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
	srv := &http.Server{Addr: fmt.Sprintf(":%v", cnf.Port), Handler: handler}
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
			if cnf.SSL {
				if err := srv.ListenAndServeTLS(cnf.Cert, cnf.Key); err != nil {
					log.Fatal("ListenAndServe: ", err)
				}
			} else {
				if err := srv.ListenAndServe(); err != nil {
					log.Fatal("ListenAndServe: ", err)
				}
			}
			logger.Printf("HTTP server listening on %q", u.Host)
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
