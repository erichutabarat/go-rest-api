package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"go-rest-api/internal/models"
	"go-rest-api/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type BlogService interface {
	CreatePost(post *models.Post, userId uint) error
	ReadAllPost() ([]models.Post, error)
	ReadByID(ctx context.Context, id uint) (*models.Post, error)
	Delete(ctx context.Context, id uint, userId uint) error
}

type blogService struct {
	repo repositories.BlogRepository
	rdb  *redis.Client
}

func NewBlogService(repo repositories.BlogRepository, rdb *redis.Client) BlogService {
	return &blogService{repo: repo, rdb: rdb}
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

func (b *blogService) ReadByID(ctx context.Context, id uint) (*models.Post, error) {
	cacheKey := fmt.Sprintf("post:%d", id)

	// Check Redis cache
	val, err := b.rdb.Get(ctx, cacheKey).Result()

	if err == nil {
		var post models.Post

		if json.Unmarshal([]byte(val), &post) == nil {
			if gin.Mode() == gin.DebugMode {
				log.Printf("[CACHE HIT] key=%s", cacheKey)
			}

			return &post, nil
		}

		if gin.Mode() == gin.DebugMode {
			log.Printf("[CACHE ERROR] Failed to deserialize cached post: key=%s", cacheKey)
		}
	} else if err == redis.Nil {
		if gin.Mode() == gin.DebugMode {
			log.Printf("[CACHE MISS] key=%s", cacheKey)
		}
	} else {
		if gin.Mode() == gin.DebugMode {
			log.Printf("[REDIS ERROR] Failed to get key=%s: %v", cacheKey, err)
		}
	}

	// Cache miss or Redis error: fetch from database
	post, err := b.repo.ReadByID(id)
	if err != nil {
		return nil, err
	}

	// Store result in Redis
	if jsonBytes, err := json.Marshal(post); err == nil {
		if err := b.rdb.Set(ctx, cacheKey, jsonBytes, 10*time.Minute).Err(); err != nil {
			if gin.Mode() == gin.DebugMode {
				log.Printf("[CACHE ERROR] Failed to set key=%s: %v", cacheKey, err)
			}
		}
	}

	return post, nil
}

func (b *blogService) Delete(ctx context.Context, id uint, userId uint) error {
	post, err := b.repo.ReadByID(id)
	if err != nil {
		return err
	}
	if post.UserID != userId {
		return errors.New("You are not authorized to perform this action")
	}

	if err := b.repo.Delete(id); err != nil {
		return err
	}

	// invalidate cache
	b.rdb.Del(ctx, fmt.Sprintf("post:%d", id))

	return nil
}
