package gqlgen

import (
	"context"
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/tiramiseb/budbud-api/internal/ownership/model"
)

type tokenKey int
type userKey int
type writertKey int

var tokenCtxKey tokenKey
var userCtxKey userKey
var writerCtxKey writertKey

// AuthnMiddleware decodes the session cookie and packs the session into context
func authnMiddleware(a authnSrv) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), writerCtxKey, &w)
			c, err := r.Cookie("budbud-authn")
			if err == nil && c != nil {
				// No error and cookie exists, let's continue digging
				user, err := a.GetUserFromToken(c.Value)
				if err == nil {
					// User exists!
					ctx = context.WithValue(ctx, userCtxKey, &user)
					ctx = context.WithValue(ctx, tokenCtxKey, c.Value)
				}
				// If no user is found, it's not this middleware's job to block the request, continuing without a user...
			}
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// authenticationDirective executes the resolver only if the request comes from an authenticated user
func authenticationDirective(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	if CurrentUser(ctx) == nil {
		return nil, fmt.Errorf("Unauthenticated")
	}
	return next(ctx)
}

// CurrentUser returns the current user stored in the context
func CurrentUser(ctx context.Context) *model.User {
	raw, _ := ctx.Value(userCtxKey).(*model.User)
	return raw
}

// CurrentToken returns the current authentication token stored in the context
func CurrentToken(ctx context.Context) string {
	raw, _ := ctx.Value(tokenCtxKey).(string)
	return raw
}

// ResponseWriter returns the response writer stored in the context
func ResponseWriter(ctx context.Context) *http.ResponseWriter {
	raw, _ := ctx.Value(writerCtxKey).(*http.ResponseWriter)
	return raw
}
