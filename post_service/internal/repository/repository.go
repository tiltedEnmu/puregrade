package repository

import (
	"github.com/tiltedEnmu/puregrade_post/internal/entities"
)

type Repository struct {
	Post
	Notifier
}

type Post interface {
	Create(post *entities.Post) error
	Get(id string) (*entities.Post, error)
	Delete(id string) error
}

type Notifier interface {
	Push(postID, authorID string) error
}
