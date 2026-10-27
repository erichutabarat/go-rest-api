package repositories

import (
	"go-rest-api/internal/models"

	"gorm.io/gorm"
)

type BlogRepository interface {
	Create(post *models.Post) error
	ReadAll() ([]models.Post, error)
	ReadByID(id uint) (*models.Post, error)
	Delete(id uint) error
}

type blogRepository struct {
	db *gorm.DB
}

func NewBlogRepository(db *gorm.DB) BlogRepository {
	return &blogRepository{db: db}
}

func (b *blogRepository) Create(post *models.Post) error {
	return b.db.Create(post).Error
}

func (b *blogRepository) ReadAll() ([]models.Post, error) {
	var posts []models.Post
	err := b.db.Find(&posts).Error
	return posts, err
}

func (b *blogRepository) ReadByID(id uint) (*models.Post, error) {
	var post models.Post
	err := b.db.First(&post, id).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (b *blogRepository) Delete(id uint) error {
	return b.db.Delete(&models.Post{}, id).Error
}
