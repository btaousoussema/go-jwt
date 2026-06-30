package main

import (
	"go-jwt/controllers"
	"go-jwt/internal/database"
	"go-jwt/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	database.ConnectToDb()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.Signup)

	r.POST("/login", controllers.Login)

	r.POST("/logout", middleware.ValidateAuth, controllers.Logout)

	r.GET("/contact", middleware.ValidateAuth, controllers.GetContacts)

	r.Run()
}
