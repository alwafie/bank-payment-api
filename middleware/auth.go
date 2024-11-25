package middleware

import (
	"belajar/utils"
	"context"
	"net/http"
	"strings"
)

type contextKey string

const TokenContextKey contextKey = "token"

func TokenAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing authorization token", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		_, err := utils.ValidateToken(token)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), TokenContextKey, token)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
