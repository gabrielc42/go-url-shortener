package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Printf("Hello Go URL Shortener! 🐚 ")

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Go URL Shortener! 🐚 ",
		})
	})

	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server [Error: %v] 🔖", err))
	}
}
