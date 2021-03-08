package main

import (
	"Gin-api-server/config"
	"Gin-api-server/controller"
	"Gin-api-server/database"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	config.InitConfig()
	database.Connect()

	r := gin.Default()

	r.GET("/test", controller.ServerTest)
	r.POST("/signup", controller.SignUp)
	r.POST("signin", controller.Signin)

	r.Run(":3000")
}
