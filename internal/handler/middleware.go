package handler

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

type contextKey string

const UserIDKey contextKey = "userID"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return []byte("VOTRE_SECRET_KEY"), nil
		})

		if err != nil {
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

			ctx := context.WithValue(r.Context(), UserIDKey, claims["user_id"])
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
		}
	})
}
