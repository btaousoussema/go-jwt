package middleware

import (
	"fmt"
	"go-jwt/models"
	"go-jwt/services"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func ValidateAuth(c *gin.Context) {
	header := c.GetHeader("Authorization")
	token := strings.Split(header, " ")[1]

	if len(token) == 0 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var user models.User
	if c.BindJSON(&user) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body.",
		})
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	fmt.Printf("User: %v", user)

	user.Token = token
	userClaims := services.ValidateJwtToken(user)
	userFromDb, err := services.GetUser(user.Email)

	if err != nil {
		fmt.Println("Invalid user.")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	id, idErr := strconv.ParseUint(userClaims.Id, 10, 32)

	if userClaims.RegisteredClaims.Valid() != nil || idErr != nil || userFromDb.Id != uint(id) {
		fmt.Println("Invalid credentials.")
		c.AbortWithStatus(http.StatusOK)
		return
	}

	c.Set("user", userFromDb)
	c.Next()
}
