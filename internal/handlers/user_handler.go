package handlers

import (
	"net/http"

	"go-rest-api/internal/models"
	"go-rest-api/pkg/response"
	"go-rest-api/internal/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		response.Error(
			c,
			http.StatusBadRequest,
			"Invalid request body",
		)
		return
	}

	if err := h.service.Register(&user); err != nil {
		response.Error(
			c,
			http.StatusInternalServerError,
			"Failed to register user",
		)
		return
	}

	response.Success(
		c,
		http.StatusCreated,
		"User registered successfully",
		nil,
	)
}

func (h *UserHandler) Login(c *gin.Context) {
	var req models.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(
			c,
			http.StatusBadRequest,
			"Username and password are required",
		)
		return
	}

	token, err := h.service.Login(req.Username, req.Password)
	if err != nil {
		response.Error(
			c,
			http.StatusUnauthorized,
			"Invalid username or password",
		)
		return
	}

	response.Success(
		c,
		http.StatusOK,
		"Login successful",
		gin.H{
			"token": token,
		},
	)
}