package api

import (
	"context"

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

// contextWithAuthInfo adds the given JWT claims to the context and returns it.
func contextWithAuthInfo(ctx context.Context, auth authInfo) context.Context {
	return context.WithValue(ctx, ctxValueClaims, auth)
}

// contextAuthInfo returns the jwt.MapClaims from the given context.
func contextAuthInfo(ctx context.Context) (auth authInfo) {
	auth, _ = ctx.Value(ctxValueClaims).(authInfo)
	return
}
