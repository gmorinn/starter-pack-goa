// Code generated by goa v3.5.2, DO NOT EDIT.
//
// crud HTTP client CLI support package
//
// Command:
// $ goa gen api_crud/design

package cli

import (
	crudc "api_crud/gen/http/crud/client"
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
	return `crud (get-book|update-book|get-all-books|delete-book|create-book|signup)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` crud get-book --id "74C53540-E974-ABFF-2565-6BF99F9017B2"` + "\n" +
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

		crudGetBookFlags  = flag.NewFlagSet("get-book", flag.ExitOnError)
		crudGetBookIDFlag = crudGetBookFlags.String("id", "REQUIRED", "")

		crudUpdateBookFlags    = flag.NewFlagSet("update-book", flag.ExitOnError)
		crudUpdateBookBodyFlag = crudUpdateBookFlags.String("body", "REQUIRED", "")
		crudUpdateBookIDFlag   = crudUpdateBookFlags.String("id", "REQUIRED", "")

		crudGetAllBooksFlags = flag.NewFlagSet("get-all-books", flag.ExitOnError)

		crudDeleteBookFlags  = flag.NewFlagSet("delete-book", flag.ExitOnError)
		crudDeleteBookIDFlag = crudDeleteBookFlags.String("id", "REQUIRED", "")

		crudCreateBookFlags    = flag.NewFlagSet("create-book", flag.ExitOnError)
		crudCreateBookBodyFlag = crudCreateBookFlags.String("body", "REQUIRED", "")

		crudSignupFlags    = flag.NewFlagSet("signup", flag.ExitOnError)
		crudSignupBodyFlag = crudSignupFlags.String("body", "REQUIRED", "")
	)
	crudFlags.Usage = crudUsage
	crudGetBookFlags.Usage = crudGetBookUsage
	crudUpdateBookFlags.Usage = crudUpdateBookUsage
	crudGetAllBooksFlags.Usage = crudGetAllBooksUsage
	crudDeleteBookFlags.Usage = crudDeleteBookUsage
	crudCreateBookFlags.Usage = crudCreateBookUsage
	crudSignupFlags.Usage = crudSignupUsage

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

			case "signup":
				epf = crudSignupFlags

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
				data, err = crudc.BuildGetBookPayload(*crudGetBookIDFlag)
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
			case "signup":
				endpoint = c.Signup()
				data, err = crudc.BuildSignupPayload(*crudSignupBodyFlag)
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
    get-book: Read Book
    update-book: Update One Book
    get-all-books: Read All Books
    delete-book: Delete Book
    create-book: Create Book
    signup: signup

Additional help:
    %[1]s crud COMMAND --help
`, os.Args[0])
}
func crudGetBookUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] crud get-book -id STRING

Read Book
    -id STRING: 

Example:
    %[1]s crud get-book --id "74C53540-E974-ABFF-2565-6BF99F9017B2"
`, os.Args[0])
}

func crudUpdateBookUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] crud update-book -body JSON -id STRING

Update One Book
    -body JSON: 
    -id STRING: 

Example:
    %[1]s crud update-book --body '{
      "name": "Sint voluptas in.",
      "price": 0.6399601382054029
   }' --id "2C91774A-6D4A-67B6-DC43-6AED849DACC1"
`, os.Args[0])
}

func crudGetAllBooksUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] crud get-all-books

Read All Books

Example:
    %[1]s crud get-all-books
`, os.Args[0])
}

func crudDeleteBookUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] crud delete-book -id STRING

Delete Book
    -id STRING: 

Example:
    %[1]s crud delete-book --id "1265498D-5A84-134A-1C7A-ED5B4B92788E"
`, os.Args[0])
}

func crudCreateBookUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] crud create-book -body JSON

Create Book
    -body JSON: 

Example:
    %[1]s crud create-book --body '{
      "name": "j71",
      "price": 0.023143062949744986
   }'
`, os.Args[0])
}

func crudSignupUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] crud signup -body JSON

signup
    -body JSON: 

Example:
    %[1]s crud signup --body '{
      "email": "cordia.muller@ferry.name",
      "firstname": "o2q",
      "lastname": "ei1",
      "password": "wi5"
   }'
`, os.Args[0])
}
