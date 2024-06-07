package middleware

import (
	"net/http"
	"strings"

	"example.com/gin/utils"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// read token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer:") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token!"})
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer:")
		// parse token
		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Could not parse token!"})
			return
		}
		// check login user
		// redis
		// check expire

		c.Set("UserEmail", claims.Email)
		c.Next()
	}
}
