package middleware

import (
	"net/http"
	"strings"

	"clinic-app/utils"

	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(next http.HandlerFunc, allowedRoles ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Missing or invalid token", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims := &utils.Claims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return utils.JwtKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		authorized := false
		for _, role := range allowedRoles {
			if claims.Role == role {
				authorized = true
				break
			}
		}

		if !authorized {
			http.Error(w, "Access denied", http.StatusForbidden)
			return
		}

		// Store email and role in headers for downstream use
		r.Header.Set("UserEmail", claims.Email)
		r.Header.Set("UserRole", claims.Role)
		next(w, r)
	}
}
