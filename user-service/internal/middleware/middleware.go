package middleware

import (
	"net/http"
	"strings"
)

func TokenAuthorization(token string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			prefix := "Bearer "
			authHeader := r.Header.Get("Authorization")
			reqToken := strings.TrimPrefix(authHeader, prefix)

			if reqToken != token {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
