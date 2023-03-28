package main

import (
	"fmt"
	"net/http"
	"sample_go/config"

	"sample_go/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	config.InitializeEnv()
	config.InitializeDb()

	fmt.Println("Starting new service!!!")

	server := gin.New()

	server.Use(gin.Logger())

	group := server.Group("/sample_go/")

	routes.CreateRoutes(group)

	group.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ping": "pong"})
	})

	if err := server.Run(":8090").Error(); err != "" {
		panic(err)
	}

}
