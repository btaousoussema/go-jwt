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

	r.POST("/signup", controllers.Signup)

	r.POST("/login", controllers.Login)

	r.Run()
}
