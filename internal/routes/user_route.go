package routes

import (
	"go-rest-api/internal/handlers"
	"go-rest-api/internal/repositories"
	"go-rest-api/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUserRoutes(router *gin.RouterGroup, db *gorm.DB) {
	// Dependency injection
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Register user routes
	userRoutes := router.Group("/user")
	{
		userRoutes.POST("/register", userHandler.Register)
		userRoutes.POST("/login", userHandler.Login)
	}
}