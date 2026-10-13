package services

import (
	"log"

	"go-rest-api/internal/config"
	"go-rest-api/internal/repositories"
)

type UserService interface {
	Register(user *config.User) error
	Login(username string, password string) error
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Register(user *config.User) error {
	if err := s.repo.Register(user); err != nil {
		log.Printf("Error registering user: %v", err)
		return err
	}
	return nil
}

func (s *userService) Login(username string, password string) error {
	if err := s.repo.Login(username, password); err != nil {
		log.Printf("Error logging in user: %v", err)
		return err
	}
	return nil
}