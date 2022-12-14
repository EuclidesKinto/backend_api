package config

import (
	"net/http"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

func CreateToken(userID uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add((time.Hour * 24) * 30).Unix()
	permissions["userId"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	return token.SignedString([]byte(SecretKey))
}

func ValidateToken(r *http.Request) error {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString)
}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}
