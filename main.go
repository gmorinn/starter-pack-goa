package main

import (
	"api_crud/api"
	auth "api_crud/gen/auth"
	boproducts "api_crud/gen/bo_products"
	bousers "api_crud/gen/bo_users"
	jwttoken "api_crud/gen/jwt_token"
	oauth "api_crud/gen/o_auth"
	products "api_crud/gen/products"
	users "api_crud/gen/users"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type ApiEndpoints struct {
	authEndpoints        *auth.Endpoints
	bo_productsEndpoints *boproducts.Endpoints
	usersEndpoints       *users.Endpoints
	bo_usersEndpoints    *bousers.Endpoints
	jwtTokenEndpoints    *jwttoken.Endpoints
	oAuthEndpoints       *oauth.Endpoints
	productsEndpoints    *products.Endpoints
}

func main() {
	// Setup logger. Replace logger with your own log package of choice.
	var (
		logger *log.Logger
		server *api.Server
	)
	{
		logger = log.New(os.Stderr, "[ecommerce] ", log.Ltime)
		server = api.NewServer()
	}

	// Initialize the services.
	var (
		authSvc       auth.Service       = api.NewAuth(logger, server)
		boproductsSvc boproducts.Service = api.NewBoProducts(logger, server)
		usersSvc      users.Service      = api.NewUsers(logger, server)
		bo_usersSvc   bousers.Service    = api.NewBoUsers(logger, server)
		jwtTokenSvc   jwttoken.Service   = api.NewJWTToken(logger, server)
		oAuthSvc      oauth.Service      = api.NewOAuth(logger, server)
		productsSvc   products.Service   = api.NewProducts(logger, server)
	)

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	var (
		apiEndpoints ApiEndpoints = ApiEndpoints{
			authEndpoints:        auth.NewEndpoints(authSvc),
			bo_productsEndpoints: boproducts.NewEndpoints(boproductsSvc),
			bo_usersEndpoints:    bousers.NewEndpoints(bo_usersSvc),
			usersEndpoints:       users.NewEndpoints(usersSvc),
			jwtTokenEndpoints:    jwttoken.NewEndpoints(jwtTokenSvc),
			oAuthEndpoints:       oauth.NewEndpoints(oAuthSvc),
			productsEndpoints:    products.NewEndpoints(productsSvc),
		}
	)
	// Define command line flags, add any other flag required to configure the
	// service.
	var (
		hostF     = flag.String("host", "localhost", "Server host (valid values: localhost)")
		domainF   = flag.String("domain", "", "Host domain name (overrides host domain specified in service design)")
		httpPortF = flag.String("http-port", "", "HTTP port (overrides host HTTP port specified in service design)")
		secureF   = flag.Bool("secure", server.Config.SSL, "Use secure scheme (https or grpcs)")
		dbgF      = flag.Bool("debug", false, "Log request and response bodies")
	)
	flag.Parse()

	// Create channel used by both the signal handler and server goroutines
	// to notify the main goroutine when to stop the server.
	errc := make(chan error)

	// Setup interrupt handler. This optional step configures the process so
	// that SIGINT and SIGTERM signals cause the services to stop gracefully.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	if server.Config.SSL {
		fmt.Printf("https://%s", server.Config.Host)
	} else {
		fmt.Printf("http://%s", server.Config.Host)
	}
	// Start the servers and send errors (if any) to the error channel.
	switch *hostF {
	case server.Config.Domain:
		{
			addr := server.Config.Host
			u, err := url.Parse(addr)
			if err != nil {
				fmt.Fprintf(os.Stderr, "invalid URL %#v: %s\n", addr, err)
				os.Exit(1)
			}
			if *secureF {
				u.Scheme = "https"
			}
			if *domainF != "" {
				u.Host = *domainF
			}
			if *httpPortF != "" {
				h, _, err := net.SplitHostPort(u.Host)
				if err != nil {
					fmt.Fprintf(os.Stderr, "invalid URL %#v: %s\n", u.Host, err)
					os.Exit(1)
				}
				u.Host = net.JoinHostPort(h, *httpPortF)
			}
			handleHTTPServer(ctx, u, &apiEndpoints, &wg, errc, logger, *dbgF)
		}

	default:
		fmt.Fprintf(os.Stderr, "invalid host argument: %q (valid hosts: %v)\n", *hostF, server.Config.Host)
	}

	// Wait for signal.
	logger.Printf("exiting (%v)", <-errc)

	// Send cancellation signal to the goroutines.
	cancel()

	wg.Wait()
	logger.Println("exited")
}
