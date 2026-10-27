package services

import (
	"errors"

	"go-rest-api/internal/models"
	"go-rest-api/internal/repositories"
)

type BlogService interface {
	CreatePost(post *models.Post, userId uint) error
	ReadAllPost() ([]models.Post, error)
	ReadByID(id uint) (*models.Post, error)
	Delete(id uint, userId uint) error
}

type blogService struct {
	repo repositories.BlogRepository
}

func NewBlogService(repo repositories.BlogRepository) BlogService {
	return &blogService{repo: repo}
}

func (b *blogService) CreatePost(post *models.Post, userId uint) error {
	post.UserID = userId

	return b.repo.Create(post)
}

func (b *blogService) ReadAllPost() ([]models.Post, error) {
	var posts []models.Post

	posts, err := b.repo.ReadAll()
	if err != nil {
		return posts, err
	}

	return posts, nil
}

func (b *blogService) ReadByID(id uint) (*models.Post, error) {
	post, err := b.repo.ReadByID(id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (b *blogService) Delete(id uint, userId uint) error {
	var post *models.Post
	post, err := b.repo.ReadByID(id)
	if err != nil {
		return err
	}
	if post.UserID != userId {
		return errors.New("You are not authorized to perform this action")
	}

	return b.repo.Delete(id)
}
