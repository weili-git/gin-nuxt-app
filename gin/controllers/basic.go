package controllers

import (
	"net/http"

	"example.com/gin/models"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) { // show all users
	var users []models.User
	models.DB.Find(&users)

	// c.JSON(http.StatusOK, gin.H{"data": users})
	c.JSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	var user []models.User
	models.DB.Where("id = ?", c.Param("id")).First(&user)

	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Updates(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteUser(c *gin.Context) {
	var user models.User
	if err := models.DB.Where("id =? ", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}
