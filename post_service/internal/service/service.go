package service

import (
	"github.com/tiltedEnmu/puregrade_post/internal/entities"
)

type Service struct {
	Post
}

type Post interface {
	CreatePost(post *entities.Post) (string, error)
	GetPost(id string) (*entities.Post, error)
	DeletePost(id string) error
}
