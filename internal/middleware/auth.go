package middleware

import (
	"net/http"
)

func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Authentication check will be implemented
		// when the session manager is connected.

		next.ServeHTTP(w, r)
	})
}
