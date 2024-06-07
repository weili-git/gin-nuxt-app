package controllers

import (
	"net/http"

	"example.com/gin/models"
	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) { // GetUserByEmail
	userEmail, exists := c.Get("UserEmail")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User email not found in context"})
		return
	}
	// type assertion
	email, ok := userEmail.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data type of email"})
		return
	}
	// get user info
	user, err := models.GetUserInfoByEmail(email)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
