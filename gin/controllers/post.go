package controllers

import (
	"net/http"
	"time"

	"example.com/gin/models"
	"github.com/gin-gonic/gin"
)

func GetPost(c *gin.Context) {
	var post models.Post
	if err := models.DB.Where("id=?", c.Param("id")).First(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Post not found"})
		return
	}
	c.JSON(http.StatusOK, post)
}

func CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post.PubDate = time.Now()
	post.ModDate = time.Now()
	models.DB.Create(&post)
	c.JSON(http.StatusCreated, gin.H{"msg": "Successfully created post"})
}
