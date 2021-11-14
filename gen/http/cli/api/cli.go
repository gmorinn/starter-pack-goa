// Code generated by goa v3.5.2, DO NOT EDIT.
//
// api HTTP client CLI support package
//
// Command:
// $ goa gen api_crud/design

package cli

import (
	bousersc "api_crud/gen/http/bo_users/client"
	jwttokenc "api_crud/gen/http/jwt_token/client"
	oauthc "api_crud/gen/http/o_auth/client"
	productsc "api_crud/gen/http/products/client"
	usersc "api_crud/gen/http/users/client"
	"flag"
	"fmt"
	"net/http"
	"os"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `bo-users (get-allusers|delete-user|create-user|update-user|get-user)
jwt-token (signup|signin|refresh|auth-providers)
o-auth o-auth
products (get-all-products|get-all-products-by-category|delete-product|create-product|update-product|get-product)
users get-user
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` bo-users get-allusers --oauth "Illum iste aliquam non consequuntur cum." --jwt-token "Amet voluptas."` + "\n" +
		os.Args[0] + ` jwt-token signup --body '{
      "birthday": "Magni animi aliquid non ut corrupti consequatur.",
      "confirm_password": "JeSuisUnTest974",
      "email": "guillaume@epitech.eu",
      "firstname": "Guillaume",
      "lastname": "Morin",
      "password": "JeSuisUnTest974",
      "phone": "+262 692 12 34 56"
   }' --oauth "Enim aut."` + "\n" +
		os.Args[0] + ` o-auth o-auth --body '{
      "client_id": "Molestiae est earum est.",
      "client_secret": "At eum.",
      "grant_type": "Molestias rem molestias earum consequuntur."
   }'` + "\n" +
		os.Args[0] + ` products get-all-products --oauth "Eum quod dolore."` + "\n" +
		os.Args[0] + ` users get-user --id "5dfb0bf7-597a-4250-b7ad-63a43ff59c25" --oauth "Omnis veritatis et tempora reiciendis commodi inventore." --jwt-token "Omnis harum temporibus."` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		boUsersFlags = flag.NewFlagSet("bo-users", flag.ContinueOnError)

		boUsersGetAllusersFlags        = flag.NewFlagSet("get-allusers", flag.ExitOnError)
		boUsersGetAllusersOauthFlag    = boUsersGetAllusersFlags.String("oauth", "", "")
		boUsersGetAllusersJWTTokenFlag = boUsersGetAllusersFlags.String("jwt-token", "", "")

		boUsersDeleteUserFlags        = flag.NewFlagSet("delete-user", flag.ExitOnError)
		boUsersDeleteUserIDFlag       = boUsersDeleteUserFlags.String("id", "REQUIRED", "")
		boUsersDeleteUserOauthFlag    = boUsersDeleteUserFlags.String("oauth", "", "")
		boUsersDeleteUserJWTTokenFlag = boUsersDeleteUserFlags.String("jwt-token", "", "")

		boUsersCreateUserFlags        = flag.NewFlagSet("create-user", flag.ExitOnError)
		boUsersCreateUserBodyFlag     = boUsersCreateUserFlags.String("body", "REQUIRED", "")
		boUsersCreateUserOauthFlag    = boUsersCreateUserFlags.String("oauth", "", "")
		boUsersCreateUserJWTTokenFlag = boUsersCreateUserFlags.String("jwt-token", "", "")

		boUsersUpdateUserFlags        = flag.NewFlagSet("update-user", flag.ExitOnError)
		boUsersUpdateUserBodyFlag     = boUsersUpdateUserFlags.String("body", "REQUIRED", "")
		boUsersUpdateUserIDFlag       = boUsersUpdateUserFlags.String("id", "REQUIRED", "")
		boUsersUpdateUserOauthFlag    = boUsersUpdateUserFlags.String("oauth", "", "")
		boUsersUpdateUserJWTTokenFlag = boUsersUpdateUserFlags.String("jwt-token", "", "")

		boUsersGetUserFlags        = flag.NewFlagSet("get-user", flag.ExitOnError)
		boUsersGetUserIDFlag       = boUsersGetUserFlags.String("id", "REQUIRED", "Unique ID of the User")
		boUsersGetUserOauthFlag    = boUsersGetUserFlags.String("oauth", "", "")
		boUsersGetUserJWTTokenFlag = boUsersGetUserFlags.String("jwt-token", "", "")

		jwtTokenFlags = flag.NewFlagSet("jwt-token", flag.ContinueOnError)

		jwtTokenSignupFlags     = flag.NewFlagSet("signup", flag.ExitOnError)
		jwtTokenSignupBodyFlag  = jwtTokenSignupFlags.String("body", "REQUIRED", "")
		jwtTokenSignupOauthFlag = jwtTokenSignupFlags.String("oauth", "", "")

		jwtTokenSigninFlags     = flag.NewFlagSet("signin", flag.ExitOnError)
		jwtTokenSigninBodyFlag  = jwtTokenSigninFlags.String("body", "REQUIRED", "")
		jwtTokenSigninOauthFlag = jwtTokenSigninFlags.String("oauth", "", "")

		jwtTokenRefreshFlags     = flag.NewFlagSet("refresh", flag.ExitOnError)
		jwtTokenRefreshBodyFlag  = jwtTokenRefreshFlags.String("body", "REQUIRED", "")
		jwtTokenRefreshOauthFlag = jwtTokenRefreshFlags.String("oauth", "", "")

		jwtTokenAuthProvidersFlags     = flag.NewFlagSet("auth-providers", flag.ExitOnError)
		jwtTokenAuthProvidersBodyFlag  = jwtTokenAuthProvidersFlags.String("body", "REQUIRED", "")
		jwtTokenAuthProvidersOauthFlag = jwtTokenAuthProvidersFlags.String("oauth", "", "")

		oAuthFlags = flag.NewFlagSet("o-auth", flag.ContinueOnError)

		oAuthOAuthFlags    = flag.NewFlagSet("o-auth", flag.ExitOnError)
		oAuthOAuthBodyFlag = oAuthOAuthFlags.String("body", "REQUIRED", "")

		productsFlags = flag.NewFlagSet("products", flag.ContinueOnError)

		productsGetAllProductsFlags     = flag.NewFlagSet("get-all-products", flag.ExitOnError)
		productsGetAllProductsOauthFlag = productsGetAllProductsFlags.String("oauth", "", "")

		productsGetAllProductsByCategoryFlags        = flag.NewFlagSet("get-all-products-by-category", flag.ExitOnError)
		productsGetAllProductsByCategoryCategoryFlag = productsGetAllProductsByCategoryFlags.String("category", "REQUIRED", "")
		productsGetAllProductsByCategoryOauthFlag    = productsGetAllProductsByCategoryFlags.String("oauth", "", "")

		productsDeleteProductFlags        = flag.NewFlagSet("delete-product", flag.ExitOnError)
		productsDeleteProductIDFlag       = productsDeleteProductFlags.String("id", "REQUIRED", "")
		productsDeleteProductOauthFlag    = productsDeleteProductFlags.String("oauth", "", "")
		productsDeleteProductJWTTokenFlag = productsDeleteProductFlags.String("jwt-token", "", "")

		productsCreateProductFlags        = flag.NewFlagSet("create-product", flag.ExitOnError)
		productsCreateProductBodyFlag     = productsCreateProductFlags.String("body", "REQUIRED", "")
		productsCreateProductOauthFlag    = productsCreateProductFlags.String("oauth", "", "")
		productsCreateProductJWTTokenFlag = productsCreateProductFlags.String("jwt-token", "", "")

		productsUpdateProductFlags        = flag.NewFlagSet("update-product", flag.ExitOnError)
		productsUpdateProductBodyFlag     = productsUpdateProductFlags.String("body", "REQUIRED", "")
		productsUpdateProductIDFlag       = productsUpdateProductFlags.String("id", "REQUIRED", "")
		productsUpdateProductOauthFlag    = productsUpdateProductFlags.String("oauth", "", "")
		productsUpdateProductJWTTokenFlag = productsUpdateProductFlags.String("jwt-token", "", "")

		productsGetProductFlags     = flag.NewFlagSet("get-product", flag.ExitOnError)
		productsGetProductIDFlag    = productsGetProductFlags.String("id", "REQUIRED", "Unique ID of the product")
		productsGetProductOauthFlag = productsGetProductFlags.String("oauth", "", "")

		usersFlags = flag.NewFlagSet("users", flag.ContinueOnError)

		usersGetUserFlags        = flag.NewFlagSet("get-user", flag.ExitOnError)
		usersGetUserIDFlag       = usersGetUserFlags.String("id", "REQUIRED", "Unique ID of the User")
		usersGetUserOauthFlag    = usersGetUserFlags.String("oauth", "", "")
		usersGetUserJWTTokenFlag = usersGetUserFlags.String("jwt-token", "", "")
	)
	boUsersFlags.Usage = boUsersUsage
	boUsersGetAllusersFlags.Usage = boUsersGetAllusersUsage
	boUsersDeleteUserFlags.Usage = boUsersDeleteUserUsage
	boUsersCreateUserFlags.Usage = boUsersCreateUserUsage
	boUsersUpdateUserFlags.Usage = boUsersUpdateUserUsage
	boUsersGetUserFlags.Usage = boUsersGetUserUsage

	jwtTokenFlags.Usage = jwtTokenUsage
	jwtTokenSignupFlags.Usage = jwtTokenSignupUsage
	jwtTokenSigninFlags.Usage = jwtTokenSigninUsage
	jwtTokenRefreshFlags.Usage = jwtTokenRefreshUsage
	jwtTokenAuthProvidersFlags.Usage = jwtTokenAuthProvidersUsage

	oAuthFlags.Usage = oAuthUsage
	oAuthOAuthFlags.Usage = oAuthOAuthUsage

	productsFlags.Usage = productsUsage
	productsGetAllProductsFlags.Usage = productsGetAllProductsUsage
	productsGetAllProductsByCategoryFlags.Usage = productsGetAllProductsByCategoryUsage
	productsDeleteProductFlags.Usage = productsDeleteProductUsage
	productsCreateProductFlags.Usage = productsCreateProductUsage
	productsUpdateProductFlags.Usage = productsUpdateProductUsage
	productsGetProductFlags.Usage = productsGetProductUsage

	usersFlags.Usage = usersUsage
	usersGetUserFlags.Usage = usersGetUserUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "bo-users":
			svcf = boUsersFlags
		case "jwt-token":
			svcf = jwtTokenFlags
		case "o-auth":
			svcf = oAuthFlags
		case "products":
			svcf = productsFlags
		case "users":
			svcf = usersFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "bo-users":
			switch epn {
			case "get-allusers":
				epf = boUsersGetAllusersFlags

			case "delete-user":
				epf = boUsersDeleteUserFlags

			case "create-user":
				epf = boUsersCreateUserFlags

			case "update-user":
				epf = boUsersUpdateUserFlags

			case "get-user":
				epf = boUsersGetUserFlags

			}

		case "jwt-token":
			switch epn {
			case "signup":
				epf = jwtTokenSignupFlags

			case "signin":
				epf = jwtTokenSigninFlags

			case "refresh":
				epf = jwtTokenRefreshFlags

			case "auth-providers":
				epf = jwtTokenAuthProvidersFlags

			}

		case "o-auth":
			switch epn {
			case "o-auth":
				epf = oAuthOAuthFlags

			}

		case "products":
			switch epn {
			case "get-all-products":
				epf = productsGetAllProductsFlags

			case "get-all-products-by-category":
				epf = productsGetAllProductsByCategoryFlags

			case "delete-product":
				epf = productsDeleteProductFlags

			case "create-product":
				epf = productsCreateProductFlags

			case "update-product":
				epf = productsUpdateProductFlags

			case "get-product":
				epf = productsGetProductFlags

			}

		case "users":
			switch epn {
			case "get-user":
				epf = usersGetUserFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "bo-users":
			c := bousersc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "get-allusers":
				endpoint = c.GetAllusers()
				data, err = bousersc.BuildGetAllusersPayload(*boUsersGetAllusersOauthFlag, *boUsersGetAllusersJWTTokenFlag)
			case "delete-user":
				endpoint = c.DeleteUser()
				data, err = bousersc.BuildDeleteUserPayload(*boUsersDeleteUserIDFlag, *boUsersDeleteUserOauthFlag, *boUsersDeleteUserJWTTokenFlag)
			case "create-user":
				endpoint = c.CreateUser()
				data, err = bousersc.BuildCreateUserPayload(*boUsersCreateUserBodyFlag, *boUsersCreateUserOauthFlag, *boUsersCreateUserJWTTokenFlag)
			case "update-user":
				endpoint = c.UpdateUser()
				data, err = bousersc.BuildUpdateUserPayload(*boUsersUpdateUserBodyFlag, *boUsersUpdateUserIDFlag, *boUsersUpdateUserOauthFlag, *boUsersUpdateUserJWTTokenFlag)
			case "get-user":
				endpoint = c.GetUser()
				data, err = bousersc.BuildGetUserPayload(*boUsersGetUserIDFlag, *boUsersGetUserOauthFlag, *boUsersGetUserJWTTokenFlag)
			}
		case "jwt-token":
			c := jwttokenc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "signup":
				endpoint = c.Signup()
				data, err = jwttokenc.BuildSignupPayload(*jwtTokenSignupBodyFlag, *jwtTokenSignupOauthFlag)
			case "signin":
				endpoint = c.Signin()
				data, err = jwttokenc.BuildSigninPayload(*jwtTokenSigninBodyFlag, *jwtTokenSigninOauthFlag)
			case "refresh":
				endpoint = c.Refresh()
				data, err = jwttokenc.BuildRefreshPayload(*jwtTokenRefreshBodyFlag, *jwtTokenRefreshOauthFlag)
			case "auth-providers":
				endpoint = c.AuthProviders()
				data, err = jwttokenc.BuildAuthProvidersPayload(*jwtTokenAuthProvidersBodyFlag, *jwtTokenAuthProvidersOauthFlag)
			}
		case "o-auth":
			c := oauthc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "o-auth":
				endpoint = c.OAuth()
				data, err = oauthc.BuildOAuthPayload(*oAuthOAuthBodyFlag)
			}
		case "products":
			c := productsc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "get-all-products":
				endpoint = c.GetAllProducts()
				data, err = productsc.BuildGetAllProductsPayload(*productsGetAllProductsOauthFlag)
			case "get-all-products-by-category":
				endpoint = c.GetAllProductsByCategory()
				data, err = productsc.BuildGetAllProductsByCategoryPayload(*productsGetAllProductsByCategoryCategoryFlag, *productsGetAllProductsByCategoryOauthFlag)
			case "delete-product":
				endpoint = c.DeleteProduct()
				data, err = productsc.BuildDeleteProductPayload(*productsDeleteProductIDFlag, *productsDeleteProductOauthFlag, *productsDeleteProductJWTTokenFlag)
			case "create-product":
				endpoint = c.CreateProduct()
				data, err = productsc.BuildCreateProductPayload(*productsCreateProductBodyFlag, *productsCreateProductOauthFlag, *productsCreateProductJWTTokenFlag)
			case "update-product":
				endpoint = c.UpdateProduct()
				data, err = productsc.BuildUpdateProductPayload(*productsUpdateProductBodyFlag, *productsUpdateProductIDFlag, *productsUpdateProductOauthFlag, *productsUpdateProductJWTTokenFlag)
			case "get-product":
				endpoint = c.GetProduct()
				data, err = productsc.BuildGetProductPayload(*productsGetProductIDFlag, *productsGetProductOauthFlag)
			}
		case "users":
			c := usersc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "get-user":
				endpoint = c.GetUser()
				data, err = usersc.BuildGetUserPayload(*usersGetUserIDFlag, *usersGetUserOauthFlag, *usersGetUserJWTTokenFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// bo-usersUsage displays the usage of the bo-users command and its subcommands.
func boUsersUsage() {
	fmt.Fprintf(os.Stderr, `users of the api
Usage:
    %[1]s [globalflags] bo-users COMMAND [flags]

COMMAND:
    get-allusers: Get All users
    delete-user: Delete one User by ID
    create-user: Create one User
    update-user: Update one User
    get-user: Get one User

Additional help:
    %[1]s bo-users COMMAND --help
`, os.Args[0])
}
func boUsersGetAllusersUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] bo-users get-allusers -oauth STRING -jwt-token STRING

Get All users
    -oauth STRING: 
    -jwt-token STRING: 

Example:
    %[1]s bo-users get-allusers --oauth "Illum iste aliquam non consequuntur cum." --jwt-token "Amet voluptas."
`, os.Args[0])
}

func boUsersDeleteUserUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] bo-users delete-user -id STRING -oauth STRING -jwt-token STRING

Delete one User by ID
    -id STRING: 
    -oauth STRING: 
    -jwt-token STRING: 

Example:
    %[1]s bo-users delete-user --id "5dfb0bf7-597a-4250-b7ad-63a43ff59c25" --oauth "Libero non nisi aut eum." --jwt-token "Voluptatum necessitatibus tempora consectetur incidunt."
`, os.Args[0])
}

func boUsersCreateUserUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] bo-users create-user -body JSON -oauth STRING -jwt-token STRING

Create one User
    -body JSON: 
    -oauth STRING: 
    -jwt-token STRING: 

Example:
    %[1]s bo-users create-user --body '{
      "user": {
         "birthday": "01/09/2002",
         "email": "guillaume.morin@epitech.eu",
         "firstname": "Guillaume",
         "lastname": "Morin",
         "phone": "+262 692 12 34 56"
      }
   }' --oauth "Adipisci ab aut saepe molestias voluptatibus." --jwt-token "Quaerat numquam consequatur placeat possimus id."
`, os.Args[0])
}

func boUsersUpdateUserUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] bo-users update-user -body JSON -id STRING -oauth STRING -jwt-token STRING

Update one User
    -body JSON: 
    -id STRING: 
    -oauth STRING: 
    -jwt-token STRING: 

Example:
    %[1]s bo-users update-user --body '{
      "User": {
         "birthday": "01/09/2002",
         "email": "guillaume.morin@epitech.eu",
         "firstname": "Guillaume",
         "lastname": "Morin",
         "phone": "+262 692 12 34 56"
      }
   }' --id "5dfb0bf7-597a-4250-b7ad-63a43ff59c25" --oauth "Velit voluptas quas qui." --jwt-token "Veniam autem laudantium et earum."
`, os.Args[0])
}

func boUsersGetUserUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] bo-users get-user -id STRING -oauth STRING -jwt-token STRING

Get one User
    -id STRING: Unique ID of the User
    -oauth STRING: 
    -jwt-token STRING: 

Example:
    %[1]s bo-users get-user --id "5dfb0bf7-597a-4250-b7ad-63a43ff59c25" --oauth "Ab est voluptatem incidunt pariatur et." --jwt-token "Ipsa ad."
`, os.Args[0])
}

// jwt-tokenUsage displays the usage of the jwt-token command and its
// subcommands.
func jwtTokenUsage() {
	fmt.Fprintf(os.Stderr, `Use Token to authenticate. Signin and Signup
Usage:
    %[1]s [globalflags] jwt-token COMMAND [flags]

COMMAND:
    signup: signup to generate jwt token
    signin: signin
    refresh: Refresh Token
    auth-providers: Register or login by Google, Facebook

Additional help:
    %[1]s jwt-token COMMAND --help
`, os.Args[0])
}
func jwtTokenSignupUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] jwt-token signup -body JSON -oauth STRING

signup to generate jwt token
    -body JSON: 
    -oauth STRING: 

Example:
    %[1]s jwt-token signup --body '{
      "birthday": "Magni animi aliquid non ut corrupti consequatur.",
      "confirm_password": "JeSuisUnTest974",
      "email": "guillaume@epitech.eu",
      "firstname": "Guillaume",
      "lastname": "Morin",
      "password": "JeSuisUnTest974",
      "phone": "+262 692 12 34 56"
   }' --oauth "Enim aut."
`, os.Args[0])
}

func jwtTokenSigninUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] jwt-token signin -body JSON -oauth STRING

signin
    -body JSON: 
    -oauth STRING: 

Example:
    %[1]s jwt-token signin --body '{
      "email": "guillaume@epitech.eu",
      "password": "JeSuisUnTest974"
   }' --oauth "Debitis odit."
`, os.Args[0])
}

func jwtTokenRefreshUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] jwt-token refresh -body JSON -oauth STRING

Refresh Token
    -body JSON: 
    -oauth STRING: 

Example:
    %[1]s jwt-token refresh --body '{
      "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ"
   }' --oauth "Culpa enim."
`, os.Args[0])
}

func jwtTokenAuthProvidersUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] jwt-token auth-providers -body JSON -oauth STRING

Register or login by Google, Facebook
    -body JSON: 
    -oauth STRING: 

Example:
    %[1]s jwt-token auth-providers --body '{
      "email": "guillaume@epitech.eu",
      "firebase_id_token": "gan",
      "firebase_provider": "facebook.com",
      "firebase_uid": "zgmURRUlcJfgDMRyjJ20xs7Rxxw2",
      "firstname": "Guillaume",
      "lastname": "Morin"
   }' --oauth "Nemo temporibus."
`, os.Args[0])
}

// o-authUsage displays the usage of the o-auth command and its subcommands.
func oAuthUsage() {
	fmt.Fprintf(os.Stderr, `Oauth to authentificate
Usage:
    %[1]s [globalflags] o-auth COMMAND [flags]

COMMAND:
    o-auth: oAuth

Additional help:
    %[1]s o-auth COMMAND --help
`, os.Args[0])
}
func oAuthOAuthUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] o-auth o-auth -body JSON

oAuth
    -body JSON: 

Example:
    %[1]s o-auth o-auth --body '{
      "client_id": "Molestiae est earum est.",
      "client_secret": "At eum.",
      "grant_type": "Molestias rem molestias earum consequuntur."
   }'
`, os.Args[0])
}

// productsUsage displays the usage of the products command and its subcommands.
func productsUsage() {
	fmt.Fprintf(os.Stderr, `Products of the E-Commerce
Usage:
    %[1]s [globalflags] products COMMAND [flags]

COMMAND:
    get-all-products: Get All products
    get-all-products-by-category: Get All products by category
    delete-product: Delete one product by ID
    create-product: Create one product
    update-product: Update one product
    get-product: Get one product

Additional help:
    %[1]s products COMMAND --help
`, os.Args[0])
}
func productsGetAllProductsUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] products get-all-products -oauth STRING

Get All products
    -oauth STRING: 

Example:
    %[1]s products get-all-products --oauth "Eum quod dolore."
`, os.Args[0])
}

func productsGetAllProductsByCategoryUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] products get-all-products-by-category -category STRING -oauth STRING

Get All products by category
    -category STRING: 
    -oauth STRING: 

Example:
    %[1]s products get-all-products-by-category --category "men" --oauth "Voluptas numquam sint quibusdam ea hic."
`, os.Args[0])
}

func productsDeleteProductUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] products delete-product -id STRING -oauth STRING -jwt-token STRING

Delete one product by ID
    -id STRING: 
    -oauth STRING: 
    -jwt-token STRING: 

Example:
    %[1]s products delete-product --id "5dfb0bf7-597a-4250-b7ad-63a43ff59c25" --oauth "In ea laboriosam quia et est." --jwt-token "Distinctio impedit optio maiores."
`, os.Args[0])
}

func productsCreateProductUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] products create-product -body JSON -oauth STRING -jwt-token STRING

Create one product
    -body JSON: 
    -oauth STRING: 
    -jwt-token STRING: 

Example:
    %[1]s products create-product --body '{
      "product": {
         "category": "men",
         "cover": "https://i.ibb.co/ypkgK0X/blue-beanie.png",
         "name": "Guillaume",
         "price": 69
      }
   }' --oauth "Distinctio inventore." --jwt-token "Sit sit est libero dolor et earum."
`, os.Args[0])
}

func productsUpdateProductUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] products update-product -body JSON -id STRING -oauth STRING -jwt-token STRING

Update one product
    -body JSON: 
    -id STRING: 
    -oauth STRING: 
    -jwt-token STRING: 

Example:
    %[1]s products update-product --body '{
      "product": {
         "category": "men",
         "cover": "https://i.ibb.co/ypkgK0X/blue-beanie.png",
         "name": "Guillaume",
         "price": 69
      }
   }' --id "5dfb0bf7-597a-4250-b7ad-63a43ff59c25" --oauth "Enim qui recusandae assumenda eum." --jwt-token "Dolores aut."
`, os.Args[0])
}

func productsGetProductUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] products get-product -id STRING -oauth STRING

Get one product
    -id STRING: Unique ID of the product
    -oauth STRING: 

Example:
    %[1]s products get-product --id "5dfb0bf7-597a-4250-b7ad-63a43ff59c25" --oauth "Aliquam aspernatur."
`, os.Args[0])
}

// usersUsage displays the usage of the users command and its subcommands.
func usersUsage() {
	fmt.Fprintf(os.Stderr, `users of the api
Usage:
    %[1]s [globalflags] users COMMAND [flags]

COMMAND:
    get-user: Get one User

Additional help:
    %[1]s users COMMAND --help
`, os.Args[0])
}
func usersGetUserUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] users get-user -id STRING -oauth STRING -jwt-token STRING

Get one User
    -id STRING: Unique ID of the User
    -oauth STRING: 
    -jwt-token STRING: 

Example:
    %[1]s users get-user --id "5dfb0bf7-597a-4250-b7ad-63a43ff59c25" --oauth "Omnis veritatis et tempora reiciendis commodi inventore." --jwt-token "Omnis harum temporibus."
`, os.Args[0])
}
