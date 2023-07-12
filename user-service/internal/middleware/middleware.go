package middleware

import (
	"context"
	"net/http"
	"strings"
	"time"

	use_case "github.com/trunov/erply-assignement-task/user-service/internal/use-case"
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

func VerifyErplyUser(erply use_case.Erply, clientCode, username, password string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.Background()
			var ctxName interface{} = "sessionKey"

			reqCookie, _ := r.Cookie("sessionKey")

			if reqCookie != nil {
				ctx := context.WithValue(r.Context(), ctxName, reqCookie.Value)
				next.ServeHTTP(w, r.WithContext(ctx))
			}

			resp, err := erply.ErplyAuthentication(ctx, clientCode, username, password)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			sessionKey := resp.Records[0].SessionKey

			// documentation says token expiry is 3600 seconds
			expire := time.Now().Add(60 * time.Minute)
			cookie := &http.Cookie{Name: "sessionKey", Value: sessionKey, Expires: expire}
			http.SetCookie(w, cookie)

			passCtx := context.WithValue(r.Context(), ctxName, sessionKey)
			next.ServeHTTP(w, r.WithContext(passCtx))
		})
	}
}
