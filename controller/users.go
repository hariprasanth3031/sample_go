package controller

import (
	"github.com/gin-gonic/gin"
	"sample_go/service"
)

func GetAllUsers(c *gin.Context) {

	output := service.GetAllUsers()

	c.JSON(output.Code, output.Data)

}
