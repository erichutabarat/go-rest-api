package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	backendPort = "8080"
)

func main() {
	server := gin.Default()

	server.GET("/ping", func (c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	log.Printf("Server running on http://localhost: %s", backendPort)

	if err := server.Run(":" + backendPort); err != nil {
		log.Fatalf("Failed to run server: %s", err)
	}
}