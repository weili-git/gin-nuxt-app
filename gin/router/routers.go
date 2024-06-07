package router

import (
	"example.com/gin/controllers"
	"example.com/gin/middleware"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors())
	r.Use(middleware.Session())

	r.GET("/", controllers.Welcome)

	basic := r.Group("/basic") // crud, to be deleted
	{
		basic.GET("/", controllers.GetUsers) // func begins with capital character
		basic.GET("/:id", controllers.GetUser)
		basic.POST("/update/:id", controllers.UpdateUser)
		basic.DELETE("/delete/:id", controllers.DeleteUser)
	}
	user := r.Group("/user")
	{
		user.POST("/login", controllers.LoginUser)
		user.POST("/register", controllers.RegisterUser)
	}
	captcha := r.Group("/captcha")
	{
		captcha.GET("/", controllers.GetCaptcha) // final slash is necessary
		captcha.GET("/:captcha_id", controllers.GetCaptchaImage)
	}
	api := r.Group("/api")
	api.Use(middleware.Auth())
	{
		api.GET("/profile", controllers.GetProfile)
	}
	return r
}
