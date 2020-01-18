package middleware

import "net/http"

// FuncHandler is a type alias for middleware function prototypes.
type FuncHandler = func(next http.Handler) http.Handler
