package controllers

import (
	"net/http"

	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
)

func GetCaptcha(c *gin.Context) {
	captchaID := captcha.NewLen(6)
	c.JSON(http.StatusOK, captchaID) // gin.H{"captcha_id": captchaID}
}

func GetCaptchaImage(c *gin.Context) {
	captchaID := c.Param("captcha_id")
	c.Header("Content-Type", "image/png")
	captcha.WriteImage(c.Writer, captchaID, 240, 80)
}

// func Verify(c *gin.Context) {
// 	var requestBody struct {
// 		CaptchaID    string `json:"captcha_id"`
// 		CaptchaValue string `json:"captcha_value"`
// 	}
// 	if err := c.BindJSON(&requestBody); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid request body"})
// 		return
// 	}

// 	fmt.Println(requestBody.CaptchaID)
// 	fmt.Println(requestBody.CaptchaValue)
// 	if captcha.VerifyString(requestBody.CaptchaID, requestBody.CaptchaValue) {
// 		c.JSON(http.StatusOK, gin.H{"msg": "Captch verified!"})
// 	} else {
// 		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid captch!"})
// 	}
// }
