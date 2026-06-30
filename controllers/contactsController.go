package controllers

import (
	"go-jwt/models"
	"go-jwt/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetContacts(c *gin.Context) {
	var user models.User

	if c.Bind(&user) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body.",
		})
		return
	}

	contacts := services.GetContacts()

	if contacts == nil {
		contacts = []services.Contact{}
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"contacts": contacts,
	})

}
