package handlers

import (
	"fmt"
	"go-rest-api/internal/models"
	"go-rest-api/internal/services"
	"go-rest-api/pkg/response"
	"net/http"
	"strconv"

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
			"Not Authorized",
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

func (h *BlogHandler) ReadAll(c *gin.Context) {
	posts, err := h.service.ReadAllPost()
	if err != nil {
		message := fmt.Sprintf("Failed get posts: %s", err)
		response.Error(
			c,
			http.StatusInternalServerError,
			message,
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

func (h *BlogHandler) ReadByID(c *gin.Context) {
	idParam := c.Param("id")
	id64, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		response.Error(
			c,
			http.StatusBadRequest,
			"Invalid post ID format",
		)
		return
	}

	post, err := h.service.ReadByID(uint(id64))
	if err != nil {
		response.Error(
			c,
			http.StatusNotFound,
			"Post not found",
		)
		return
	}

	response.Success(
		c,
		http.StatusOK,
		"Success get post",
		post,
	)
	return
}

func (h *BlogHandler) Delete(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		response.Error(
			c,
			http.StatusUnauthorized,
			"Not Authorized",
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

	idParam := c.Param("id")
	id64, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		response.Error(
			c,
			http.StatusBadRequest,
			"Invalid post ID format",
		)
		return
	}

	err = h.service.Delete(uint(id64), userID)
	if err != nil {
		if err.Error() == "You are not authorized to perform this action" {
			response.Error(
				c,
				http.StatusForbidden,
				err.Error(),
			)
			return
		}

		response.Error(
			c,
			http.StatusInternalServerError,
			fmt.Sprintf("Failed to delete post: %v", err),
		)
		return
	}

	response.Success(
		c,
		http.StatusOK,
		"Delete post successfully",
		nil,
	)
	return
}
