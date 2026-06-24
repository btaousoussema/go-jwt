package controllers

import (
	"fmt"
	"go-jwt/internal/database"
	model "go-jwt/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	var user model.User

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

	userToInsert := model.User{Email: user.Email, Password: string(hash)}

	query := "Insert into users (email, password) SELECT $1, $2 WHERE not exists (Select 1 from users where email = $1 ) RETURNING email"
	stmt, queryErr := database.DB.Prepare(query)

	if queryErr != nil {
		log.Fatalf("Failed to prepare statement: %v", queryErr)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error creating the query statement.",
		})
		return
	}

	insertErr := stmt.QueryRow(userToInsert.Email, userToInsert.Password)

	if insertErr.Err() != nil {
		fmt.Println("Error is ", insertErr.Err())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
