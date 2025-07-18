package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()

	server.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	server.Run(":8989")
}
