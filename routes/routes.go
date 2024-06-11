package routes

import (
	"eft-private-server/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/users/:id", controllers.GetUser)
	r.POST("/users", controllers.CreateUser)

	return r
}
