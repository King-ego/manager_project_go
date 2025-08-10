package main

import (
	"manager_project/router"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	router.SetupAllRoutes(server)

	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	server.Run(":8989")
}
