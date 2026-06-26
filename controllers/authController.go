package controllers

import (
	"fmt"
	"go-jwt/models"
	"go-jwt/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var user models.User

	if c.Bind(&user) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body.",
		})
		return
	}

	userFromDb, err := services.GetUser(user.Email)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Failed to authenticate user.",
		})
		return
	}

	validatePasswordErr := bcrypt.CompareHashAndPassword([]byte(userFromDb.Password), []byte(user.Password))

	if validatePasswordErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Failed to authenticate user.",
		})
		return
	}

	services.InvalidateTokenForUser(user)

	token, err := services.GenerateJwtToken(userFromDb)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create a jwt token.",
		})
		return
	}

	refreshToken, err := services.CreateRefreshToken(userFromDb)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create a refresh token.",
		})
		return
	}

	c.SetCookie("refreshToken", refreshToken.Token, 86400, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Logout(c *gin.Context) {

	var user models.User

	if c.Bind(&user) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body.",
		})
		return
	}

	userClaims := services.ValidateJwtToken(user)

	if userClaims.RegisteredClaims.Valid() != nil {
		fmt.Println("Invalid credentials.")
		c.JSON(http.StatusOK, gin.H{})
		return
	}

	services.InvalidateTokenForUser(user)

	c.SetCookie("refreshToken", "", 0, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{})
}
