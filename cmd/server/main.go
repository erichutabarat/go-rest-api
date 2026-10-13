package main

import (
	"log"
	"net/http"
	"os"

	"go-rest-api/internal/config"
	"go-rest-api/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error Loading .env file")
	}

	// Load PORT
	port := os.Getenv("PORT")

	// initialize db connection
	config.ConnectDB()

	// Setup gin server
	server := routes.SetupRouter(config.DB)

	server.GET("/ping", func (c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	log.Printf("Server running on http://localhost: %s", port)

	if err := server.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %s", err)
	}
}