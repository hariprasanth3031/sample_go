package routes

import (
	"github.com/gin-gonic/gin"
	"sample_go/controller"
)

func CreateUserRoutes(server *gin.RouterGroup) {

	user := server.Group("/public/users")
	{
		user.GET("/get", controller.GetAllUsers)
	}
}
