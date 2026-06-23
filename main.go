package main

import (
	"go-jwt/controllers"
	"go-jwt/internal/database"

	"github.com/gin-gonic/gin"
)

func init() {
	database.ConnectToDb()
}

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World 2",
		})
	})

	r.POST("/signup", controllers.Signup)

	r.Run()
}
