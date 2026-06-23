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
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body.",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password.",
		})
		return
	}

	user := model.User{Email: body.Email, Password: string(hash)}

	connErr := database.DB.Ping()
	if connErr != nil {
		fmt.Println("PING ERRRR *********************************************")
	}
	query := "Insert into users (email, password) SELECT $1, $2 WHERE not exists (Select 1 from users where email = $1 ) RETURNING email"
	//query := "SELECT * FROM users"
	stmt, queryErr := database.DB.Prepare(query)
	//defer stmt.Close()

	fmt.Println("************* The user *********: ")

	if queryErr != nil {
		log.Fatalf("Failed to prepare statement: %v", queryErr)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error creating the query statement.",
		})
		return
	}

	var u model.User

	insertErr := stmt.QueryRow(user.Email, user.Password).Scan(&u.Email)

	if insertErr != nil {
		fmt.Println("Error is " + insertErr.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
