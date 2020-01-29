package middleware

import (
	"context"
	"net/http"

	"github.com/hashamali/gauth"
)

// GetJWTAuthMiddleware returns a middleware function that runs JWT validation.
func GetJWTAuthMiddleware(jwt *gauth.JWTAuth, tokenName string) FuncHandler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var token interface{}
			var err error

			cookie, err := r.Cookie(tokenName)
			if err != nil {
				tokenHeader := r.Header.Get("Authorization")
				token, err = jwt.ExtractFromHeader(tokenHeader)
			} else {
				token, err = jwt.Extract(cookie.Value)
			}

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
			} else {
				c := context.WithValue(r.Context(), gauth.ContextJWTMetaKey, token)
				r = r.WithContext(c)
				next.ServeHTTP(w, r)
			}
		})
	}
}
