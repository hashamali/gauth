package middleware

import (
	"net/http"

	"github.com/hashamali/gauth"
)

// GetBasicAuthConfigMiddleware returns a middleware function that validates that ensures basic auth.
func GetBasicAuthConfigMiddleware(auth gauth.BasicAuth, realm string) FuncHandler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, pass, ok := r.BasicAuth()
			if !ok || !auth.Validate(user, pass) {
				w.WriteHeader(http.StatusUnauthorized)
				w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
			} else {
				next.ServeHTTP(w, r)
			}
		})
	}
}
