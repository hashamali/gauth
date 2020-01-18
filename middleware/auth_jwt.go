package middleware

import (
	"context"
	"net/http"

  "github.com/hashamali/gauth"
)

// GetJWTAuthMiddleware returns a middleware function that runs JWT validation.
func GetJWTAuthMiddleware(jwt *gauth.JWTAuth) FuncHandler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenHeader := r.Header.Get("Authorization")
			token, _, err := jwt.Extract(tokenHeader)

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
