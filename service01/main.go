package main

import (
	"trab02/service01/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/usuarios", controllers.GetUsers)
	router.GET("/usuarios/:id", controllers.GetUser)
	router.POST("/usuarios", controllers.PostUser)
	router.GET("/usuarios/login", controllers.Login)
	router.DELETE("/usuarios/:id", controllers.DeleteUser)
	router.PUT("/usuarios/:id", controllers.UpdateUser)
	router.Run(":8080")
}
