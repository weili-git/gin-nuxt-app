package main

import (
	"example.com/gin/models"
	"example.com/gin/router"
)

func main() {
	models.ConnectDatabase()
	// models.Migrate()
	r := router.Router()
	r.Run(":8086")
}
