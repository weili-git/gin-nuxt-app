package controllers

import (
	"net/http"
	"time"

	"example.com/gin/models"
	"example.com/gin/utils"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
)

func LoginUser(c *gin.Context) {
	var requestBody struct {
		CaptchaID    string `json:"captcha_id"`
		CaptchaValue string `json:"captcha_value"`
		Email        string `json:"email"` // Capital!!!
		Password     string `json:"password"`
	}
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body!"})
		return
	}
	// verify captcha
	if !captcha.VerifyString(requestBody.CaptchaID, requestBody.CaptchaValue) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Captcha mismatch!"})
		return
	}
	// verify user info
	var existingUser models.User
	result := models.DB.Where("email=? AND password=?", requestBody.Email, requestBody.Password).First(&existingUser)
	if result.RowsAffected < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found."})
		return
	}
	// return token
	token, err := utils.GenerateToken(existingUser.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not generate token."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "Login OK!", "token": token})
}

func RegisterUser(c *gin.Context) {
	var RegisterInfo struct {
		FirstName    string    `json:"first_name"`
		LastName     string    `json:"last_name"`
		Email        string    `json:"email"`
		Password     string    `json:"password"`
		Birthday     time.Time `json:"birthday"`
		CaptchaID    string    `json:"captcha_id"`
		CaptchaValue string    `json:"captcha_value"`
	}
	if err := c.BindJSON(&RegisterInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// validate captcha
	if !captcha.VerifyString(RegisterInfo.CaptchaID, RegisterInfo.CaptchaValue) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Captcha mismatch!"})
		return
	}
	// create user
	user := models.User{
		FirstName: RegisterInfo.FirstName,
		LastName:  RegisterInfo.LastName,
		Email:     RegisterInfo.Email,
		Password:  RegisterInfo.Password, // You should hash the password before saving it
		Birthday:  RegisterInfo.Birthday,
	}
	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "logged out"})
}
