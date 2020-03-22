package middleware

import (
	"context"
	"net/http"

	"github.com/hashamali/gauth"
)

// GetJWTAuthMiddleware returns a middleware function that runs JWT validation.
func GetJWTAuthMiddleware(
	jwt *gauth.JWTAuth,
	accessTokenCookie string,
	refreshTokenCookie string,
	onValidRefreshToken func(w http.ResponseWriter, refreshToken interface{},
	) (interface{}, error)) FuncHandler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var token interface{}
			var err error

			accessCookie, err := r.Cookie(accessTokenCookie)
			if err != nil {
				tokenHeader := r.Header.Get("Authorization")

				err = nil
				token, err = jwt.ExtractFromHeader(tokenHeader)
			} else {
				err = nil
				token, err = jwt.Extract(accessCookie.Value)
			}

			if err != nil {
				// Try to process refresh cookie.
				var refreshCookie *http.Cookie
				var refreshToken interface{}

				err = nil
				refreshCookie, err = r.Cookie(refreshTokenCookie)
				if err == nil {
					refreshToken, err = jwt.Extract(refreshCookie.Value)
					if err == nil {
						token, err = onValidRefreshToken(w, refreshToken)
					}
				}

				// Otherwise return 401.
				if err != nil {
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
			}

			c := context.WithValue(r.Context(), gauth.ContextJWTMetaKey, token)
			r = r.WithContext(c)
			next.ServeHTTP(w, r)
		})
	}
}
