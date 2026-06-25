package services

import (
	"go-jwt/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateJwtToken(user models.User) (string, error) {
	var jwtKey = []byte("your_super_secret_key_here")

	claims := jwt.RegisteredClaims{Issuer: "go-jwt", Subject: string(user.Id), ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute))}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtKey)
}
