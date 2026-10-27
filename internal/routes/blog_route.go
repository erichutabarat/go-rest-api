package routes

import (
	"go-rest-api/internal/handlers"
	"go-rest-api/internal/middlewares"
	"go-rest-api/internal/repositories"
	"go-rest-api/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterBlogRoute(router *gin.RouterGroup, db *gorm.DB) {
	blogRepo := repositories.NewBlogRepository(db)
	blogService := services.NewBlogService(blogRepo)
	blogHandler := handlers.NewBlogHandler(blogService)

	// blogRoutes := router.Group("/blog")

	protectedBlogRoutes := router.Group("/blog")
	protectedBlogRoutes.Use(middlewares.AuthMiddleware())
	{
		protectedBlogRoutes.POST("/post", blogHandler.CreatePost)
	}
}
