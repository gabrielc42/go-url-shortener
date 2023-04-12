package main

import (
	"fmt"

	"github.com/gabrielc42/go-url-shortener/handler"
	"github.com/gabrielc42/go-url-shortener/store"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Printf("Hello Go URL Shortener! 🐚 ")

	r := gin.Default()

	r.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Go URL Shortener! 🐚 ",
		})
	})

	r.POST("/create-short0url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	store.InitializeStore()

	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server [Error: %v] 🔖", err))
	}
}
