package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/diablovocado/declutr/modules/auth/repository"
)

type contextKey string

const UserIDKey contextKey = "userID"

func Auth(repo repository.UserRepository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			authHeader := r.Header.Get("Authorization")

			if authHeader == "" {
				http.Error(w, "missing authorization header", http.StatusUnauthorized)
				return
			}

			token := strings.TrimPrefix(authHeader, "Bearer ")

			session, err := repo.GetSessionByToken(token)
			if err != nil || session == nil {
				http.Error(w, "invalid session", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), UserIDKey, session.UserID)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
