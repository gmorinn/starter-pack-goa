// Code generated by goa v3.5.2, DO NOT EDIT.
//
// crud HTTP client CLI support package
//
// Command:
// $ goa gen api_crud/design

package cli

import (
	crudc "api_crud/gen/http/crud/client"
	jwttokenc "api_crud/gen/http/jwt_token/client"
	oauthc "api_crud/gen/http/o_auth/client"
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
	return `crud (get-book|update-book|get-all-books|delete-book|create-book)
jwt-token (signup|signin)
o-auth o-auth
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` crud get-book --id "5dfb0bf7-597a-4250-b7ad-63a43ff59c25" --jwt-token "Voluptatem et reiciendis."` + "\n" +
		os.Args[0] + ` jwt-token signup --body '{
      "email": "guillaume@epitech.eu",
      "firstname": "Guillaume",
      "lastname": "Morin",
      "password": "map"
   }'` + "\n" +
		os.Args[0] + ` o-auth o-auth --client-id "Error beatae accusantium qui accusantium voluptates et." --client-secret "Nisi molestiae." --grant-type "Et voluptas rerum tempore."` + "\n" +
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
		crudFlags = flag.NewFlagSet("crud", flag.ContinueOnError)

		crudGetBookFlags        = flag.NewFlagSet("get-book", flag.ExitOnError)
		crudGetBookIDFlag       = crudGetBookFlags.String("id", "REQUIRED", "")
		crudGetBookJWTTokenFlag = crudGetBookFlags.String("jwt-token", "", "")

		crudUpdateBookFlags    = flag.NewFlagSet("update-book", flag.ExitOnError)
		crudUpdateBookBodyFlag = crudUpdateBookFlags.String("body", "REQUIRED", "")
		crudUpdateBookIDFlag   = crudUpdateBookFlags.String("id", "REQUIRED", "")

		crudGetAllBooksFlags = flag.NewFlagSet("get-all-books", flag.ExitOnError)

		crudDeleteBookFlags  = flag.NewFlagSet("delete-book", flag.ExitOnError)
		crudDeleteBookIDFlag = crudDeleteBookFlags.String("id", "REQUIRED", "")

		crudCreateBookFlags    = flag.NewFlagSet("create-book", flag.ExitOnError)
		crudCreateBookBodyFlag = crudCreateBookFlags.String("body", "REQUIRED", "")

		jwtTokenFlags = flag.NewFlagSet("jwt-token", flag.ContinueOnError)

		jwtTokenSignupFlags    = flag.NewFlagSet("signup", flag.ExitOnError)
		jwtTokenSignupBodyFlag = jwtTokenSignupFlags.String("body", "REQUIRED", "")

		jwtTokenSigninFlags    = flag.NewFlagSet("signin", flag.ExitOnError)
		jwtTokenSigninBodyFlag = jwtTokenSigninFlags.String("body", "REQUIRED", "")

		oAuthFlags = flag.NewFlagSet("o-auth", flag.ContinueOnError)

		oAuthOAuthFlags            = flag.NewFlagSet("o-auth", flag.ExitOnError)
		oAuthOAuthClientIDFlag     = oAuthOAuthFlags.String("client-id", "REQUIRED", "")
		oAuthOAuthClientSecretFlag = oAuthOAuthFlags.String("client-secret", "REQUIRED", "")
		oAuthOAuthGrantTypeFlag    = oAuthOAuthFlags.String("grant-type", "REQUIRED", "")
	)
	crudFlags.Usage = crudUsage
	crudGetBookFlags.Usage = crudGetBookUsage
	crudUpdateBookFlags.Usage = crudUpdateBookUsage
	crudGetAllBooksFlags.Usage = crudGetAllBooksUsage
	crudDeleteBookFlags.Usage = crudDeleteBookUsage
	crudCreateBookFlags.Usage = crudCreateBookUsage

	jwtTokenFlags.Usage = jwtTokenUsage
	jwtTokenSignupFlags.Usage = jwtTokenSignupUsage
	jwtTokenSigninFlags.Usage = jwtTokenSigninUsage

	oAuthFlags.Usage = oAuthUsage
	oAuthOAuthFlags.Usage = oAuthOAuthUsage

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
		case "crud":
			svcf = crudFlags
		case "jwt-token":
			svcf = jwtTokenFlags
		case "o-auth":
			svcf = oAuthFlags
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
		case "crud":
			switch epn {
			case "get-book":
				epf = crudGetBookFlags

			case "update-book":
				epf = crudUpdateBookFlags

			case "get-all-books":
				epf = crudGetAllBooksFlags

			case "delete-book":
				epf = crudDeleteBookFlags

			case "create-book":
				epf = crudCreateBookFlags

			}

		case "jwt-token":
			switch epn {
			case "signup":
				epf = jwtTokenSignupFlags

			case "signin":
				epf = jwtTokenSigninFlags

			}

		case "o-auth":
			switch epn {
			case "o-auth":
				epf = oAuthOAuthFlags

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
		case "crud":
			c := crudc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "get-book":
				endpoint = c.GetBook()
				data, err = crudc.BuildGetBookPayload(*crudGetBookIDFlag, *crudGetBookJWTTokenFlag)
			case "update-book":
				endpoint = c.UpdateBook()
				data, err = crudc.BuildUpdateBookPayload(*crudUpdateBookBodyFlag, *crudUpdateBookIDFlag)
			case "get-all-books":
				endpoint = c.GetAllBooks()
				data = nil
			case "delete-book":
				endpoint = c.DeleteBook()
				data, err = crudc.BuildDeleteBookPayload(*crudDeleteBookIDFlag)
			case "create-book":
				endpoint = c.CreateBook()
				data, err = crudc.BuildCreateBookPayload(*crudCreateBookBodyFlag)
			}
		case "jwt-token":
			c := jwttokenc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "signup":
				endpoint = c.Signup()
				data, err = jwttokenc.BuildSignupPayload(*jwtTokenSignupBodyFlag)
			case "signin":
				endpoint = c.Signin()
				data, err = jwttokenc.BuildSigninPayload(*jwtTokenSigninBodyFlag)
			}
		case "o-auth":
			c := oauthc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "o-auth":
				endpoint = c.OAuth()
				data, err = oauthc.BuildOAuthPayload(*oAuthOAuthClientIDFlag, *oAuthOAuthClientSecretFlag, *oAuthOAuthGrantTypeFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// crudUsage displays the usage of the crud command and its subcommands.
func crudUsage() {
	fmt.Fprintf(os.Stderr, `The principe of CRUD API with GET, PUT, POST, DELETE
Usage:
    %[1]s [globalflags] crud COMMAND [flags]

COMMAND:
    get-book: Get one item
    update-book: Update one item
    get-all-books: Read All items
    delete-book: Delete one item by ID
    create-book: Create one item

Additional help:
    %[1]s crud COMMAND --help
`, os.Args[0])
}
func crudGetBookUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] crud get-book -id STRING -jwt-token STRING

Get one item
    -id STRING: 
    -jwt-token STRING: 

Example:
    %[1]s crud get-book --id "5dfb0bf7-597a-4250-b7ad-63a43ff59c25" --jwt-token "Voluptatem et reiciendis."
`, os.Args[0])
}

func crudUpdateBookUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] crud update-book -body JSON -id STRING

Update one item
    -body JSON: 
    -id STRING: 

Example:
    %[1]s crud update-book --body '{
      "name": "Guillaume",
      "price": 0.1253396498855919
   }' --id "5dfb0bf7-597a-4250-b7ad-63a43ff59c25"
`, os.Args[0])
}

func crudGetAllBooksUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] crud get-all-books

Read All items

Example:
    %[1]s crud get-all-books
`, os.Args[0])
}

func crudDeleteBookUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] crud delete-book -id STRING

Delete one item by ID
    -id STRING: 

Example:
    %[1]s crud delete-book --id "5dfb0bf7-597a-4250-b7ad-63a43ff59c25"
`, os.Args[0])
}

func crudCreateBookUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] crud create-book -body JSON

Create one item
    -body JSON: 

Example:
    %[1]s crud create-book --body '{
      "name": "Guillaume",
      "price": 0.1436730351084788
   }'
`, os.Args[0])
}

// jwt-tokenUsage displays the usage of the jwt-token command and its
// subcommands.
func jwtTokenUsage() {
	fmt.Fprintf(os.Stderr, `Use Token to authenticate. Signin and Signup
Usage:
    %[1]s [globalflags] jwt-token COMMAND [flags]

COMMAND:
    signup: signup
    signin: signin

Additional help:
    %[1]s jwt-token COMMAND --help
`, os.Args[0])
}
func jwtTokenSignupUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] jwt-token signup -body JSON

signup
    -body JSON: 

Example:
    %[1]s jwt-token signup --body '{
      "email": "guillaume@epitech.eu",
      "firstname": "Guillaume",
      "lastname": "Morin",
      "password": "map"
   }'
`, os.Args[0])
}

func jwtTokenSigninUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] jwt-token signin -body JSON

signin
    -body JSON: 

Example:
    %[1]s jwt-token signin --body '{
      "email": "guillaume@epitech.eu",
      "password": "sa2"
   }'
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
	fmt.Fprintf(os.Stderr, `%[1]s [flags] o-auth o-auth -client-id STRING -client-secret STRING -grant-type STRING

oAuth
    -client-id STRING: 
    -client-secret STRING: 
    -grant-type STRING: 

Example:
    %[1]s o-auth o-auth --client-id "Error beatae accusantium qui accusantium voluptates et." --client-secret "Nisi molestiae." --grant-type "Et voluptas rerum tempore."
`, os.Args[0])
}
