package handlers

import (
	"fmt"
	"go-rest-api/internal/models"
	"go-rest-api/internal/services"
	"go-rest-api/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BlogHandler struct {
	service services.BlogService
}

func NewBlogHandler(service services.BlogService) *BlogHandler {
	return &BlogHandler{service: service}
}

func (h *BlogHandler) CreatePost(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		response.Error(
			c,
			http.StatusForbidden,
			"Not Auhtorized",
		)
		return
	}

	userID, ok := userId.(uint)
	if !ok {
		response.Error(
			c,
			http.StatusUnauthorized,
			"Invalid user identity",
		)
		return
	}

	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		response.Error(
			c,
			http.StatusBadRequest,
			"Post format invalid",
		)
		return
	}

	err := h.service.CreatePost(&post, userID)
	if err != nil {
		message := fmt.Sprintf("Failed create post: %v", err)
		response.Error(
			c,
			http.StatusInternalServerError,
			message,
		)
		return
	}

	response.Success(
		c,
		http.StatusCreated,
		"Create post successfully",
		post,
	)
	return
}

func (b *BlogHandler) ReadAll(c *gin.Context) {
	posts, err := b.service.ReadAllPost()
	if err != nil {
		Message := fmt.Sprintf("Failed get posts: %s", err)
		response.Error(
			c,
			http.StatusInternalServerError,
			Message,
		)
		return
	}
	response.Success(
		c,
		http.StatusOK,
		"Success get posts",
		posts,
	)
	return
}
