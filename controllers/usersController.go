package controllers

import (
	"fmt"
	"go-jwt/models"
	"net/http"

	"github.com/gin-gonic/gin"

	"go-jwt/services"

	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	var user models.User

	if c.Bind(&user) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body.",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password.",
		})
		return
	}

	userToInsert := models.User{Email: user.Email, Password: string(hash)}
	insertErr := services.InsertUser(userToInsert)

	if insertErr != nil {
		fmt.Printf("Error is %v", insertErr.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
