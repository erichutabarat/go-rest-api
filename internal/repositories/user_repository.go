package repositories

import (
	"go-rest-api/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Register(user *models.User) error
	Login(username string, password string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Register(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) Login(username string, password string) error {
	var user models.User
	if err := r.db.Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
		return err
	}
	return nil
}