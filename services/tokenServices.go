package services

import (
	"go-jwt/models"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type UserClaims struct {
	Id    string
	Email string
	jwt.RegisteredClaims
}

func GenerateJwtToken(user models.User) (string, error) {
	var jwtKey = []byte("your_super_secret_key_here")

	id := strconv.FormatUint(uint64(user.Id), 10)
	claims := UserClaims{
		Id:    id,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-jwt",
			Subject:   id,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtKey)
}

func ValidateJwtToken(user models.User) *UserClaims {

	parsedAccessToken, _ := jwt.ParseWithClaims(user.Token, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("your_super_secret_key_here")), nil
	})

	return parsedAccessToken.Claims.(*UserClaims)
}
