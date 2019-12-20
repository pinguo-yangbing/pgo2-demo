package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// Use the following code if you need to write the logs to file and console at the same time.
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.Default()
	router.GET("/test/index", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data":"hello world",
			"status":200,
			"message": "pong",
		})
	})

	router.Run(":8080")
}