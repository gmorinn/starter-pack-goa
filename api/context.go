package api

import (
	"context"
	"fmt"
	oauth "starter-pack-goa/gen/o_auth"

	"github.com/dgrijalva/jwt-go"
)

type authInfo struct {
	oAuth    jwt.MapClaims
	jwtToken jwt.MapClaims
}

type ctxValue int

const (
	ctxValueClaims ctxValue = iota
)

var (
	ErrNullPayload        error = fmt.Errorf("payload is null")
	ErrRessourceNotFound  error = fmt.Errorf("ressource is not found")
	ErrEmailExist         error = fmt.Errorf("email already exists")
	ErrInvalidTokenScopes error = fmt.Errorf("invalid scope token")
	ErrInvalidToken       error = fmt.Errorf("invalid format token")
	ErrInvalidPassword    error = fmt.Errorf("invalid password")
	ErrBadRole            error = fmt.Errorf("user has bad role")
	ErrWrongIdFormat      error = fmt.Errorf("wrong id format")
	ErrUserNotExist       error = fmt.Errorf("user not exist")

	// Oauth errors
	ErrUnsupportedGrantType error = oauth.Unauthorized("unsupported grant")
	ErrInvalidRequest       error = oauth.Unauthorized("invalid request")

	// Paypal error
	ErrInfoPaypal           error = fmt.Errorf("payload paypal is wrong")
	ErrUserAlreadySubscribe error = fmt.Errorf("user already have a subscription")
)

// contextWithAuthInfo adds the given JWT claims to the context and returns it.
func contextWithAuthInfo(ctx context.Context, auth authInfo) context.Context {
	return context.WithValue(ctx, ctxValueClaims, auth)
}

// contextAuthInfo returns the jwt.MapClaims from the given context.
func contextAuthInfo(ctx context.Context) (auth authInfo) {
	auth, _ = ctx.Value(ctxValueClaims).(authInfo)
	return
}
