package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://127.0.0.1:3001"}
	// config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowMethods = []string{"GET", "POST", "DELETE", "OPTIONS", "PATCH"}
	config.AllowHeaders = []string{
		"Authorization", "Content-Type", "Upgrade", "Origin",
		"Connection", "Accept-Encoding", "Accept-Language", "Host",
		"Access-Control-Request-Method", "Access-Control-Request-Headers"}

	return cors.New(config)
}
