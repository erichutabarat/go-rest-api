package services

import (
	"log"
	"errors"
	
	"go-rest-api/internal/models"
	"go-rest-api/internal/repositories"
	"go-rest-api/pkg/hashing"
)

type UserService interface {
	Register(user *models.User) error
	Login(username string, password string) error
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Register(user *models.User) error {
	// 1. Hash the password using the pkg/hashing utility
	hashedPassword, err := hashing.HashPassword(user.Password)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return err
	}
	
	// 2. Assign the hashed password back to the user model
	user.Password = hashedPassword

	// 3. Save to repository
	if err := s.repo.Register(user); err != nil {
		log.Printf("Error registering user: %v", err)
		return err
	}
	return nil
}

func (s *userService) Login(username string, password string) error {
	// 1. Find user by username
	user, err := s.repo.Login(username)
	if err != nil {
		log.Printf("User not found: %v", err)
		return errors.New("invalid username or password")
	}

	// 2. Verify password
	if !hashing.CheckPasswordHash(password, user.Password) {
		log.Println("Password mismatch")
		return errors.New("invalid username or password")
	}

	return nil
}