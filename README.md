# GOA X GM

### What is that ?
A starter-pack to make its APIs in golang.
We'll use Goa, it's a framework for building micro-services and APIs in Go using a unique design-first approach.

## Why I use GOA DESIGN ?

### Reason n°1: Postman updated every time the design file is modified
Just do:
```sh
make api-goa
```
File: openapi.json
---------------------
You will have a new postman with all the expected requests, payloads and authentication system

### Reason n°2: Automatically updated online documentation without writing a single line of HTML, CSS and Javascript
Documentation is based on generated postman file
Just do:
```sh
make api-doc
```
In this documentation you will have:
- All possible requests
![](documentation/doc.png)

--------------------------

- All expected payloads for each request with examples
![](documentation/parameters.png)


--------------------------

- All responses to each request
![](documentation/response.png)

--------------------------

- All expected structures with their types (int, float, string, boolean)
![](documentation/struct.png)

### Reason n°3: 70% of your code is generated
As the Goa design language is Go DSL, it is easy to customize and understandable by anyone.
It is from this GOA DSL that all the code will be generated.

```go
	Method("createBook", func() {
		Description("Create one item")
		Security(OAuth2, JWTAuth)
		Payload(func() {
			Attribute("name", String, func() {
				MinLength(3)
				MaxLength(10)
				Description("Name of the book")
				Example("Guillaume")
			})
			Attribute("price", Float64, func() {
				Description("Price of the book")
				Minimum(0.1)
			})
			TokenField(1, "jwtToken", String, func() {
				Description("JWT used for authentication after Signin/Signup")
			})
			AccessTokenField(2, "oauth", String, func() {
				Description("Use to generate Oauth with /authorization")
			})
			Required("name", "price", "oauth", "jwtToken")
		})
		HTTP(func() {
			POST("/book/add")
			Header("oauth:Authorization", String, "OAuth token", func() {
				Pattern("^Bearer [^ ]+$")
			})
			Header("jwtToken:jwtToken", String, "Jwt token", func() {
				Pattern("^Bearer [^ ]+$")
			})
			Response(StatusCreated)
		})
		Result(func() {
            Attribute("id", String, func() {
				Format(FormatUUID)
				Example("5dfb0bf7-597a-4250-b7ad-63a43ff59c25")
			})
			Attribute("name", String, func() {
				Example("Guillaume")
			})
			Attribute("price", Float64, func() {
				Example(69.0)
			})
			Attribute("success", Boolean)
			Required("id", "name", "price", "success")
		})
	})
})
```
**For example:**
We understand very clearly that this query:
=> Create an item
=> Attend Name, price, jwtToken, oauth as payload
=> The route is /book/add
=> The response will return id, name, price and a success



