package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	UserID string
	jwt.StandardClaims
}

type contextKey string

var (
	UserCtxKey = contextKey("user")
)

// Middleware is used to handle auth logic
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		auth := r.Header.Get("Authorization")

		if auth != "" {
			splitAuthStrings := strings.Split(auth, "Bearer ")
			tokenString := ""
			if len(splitAuthStrings) >= 1 {
				tokenString = splitAuthStrings[1]
			}

			token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte("my_secret"), nil
			})

			if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
				fmt.Printf("claims--------> %v ", claims.UserID)
				ctx = context.WithValue(ctx, UserCtxKey, claims.UserID)
			} else {
				fmt.Println(err)
			}
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
