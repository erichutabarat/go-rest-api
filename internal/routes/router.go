package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// API v1 route group
	v1 := r.Group("/api/v1")

	// Register user routes
	RegisterUserRoutes(v1, db)

	return r
}