package service

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/tiltedEnmu/puregrade_post/internal/entities"
	"github.com/tiltedEnmu/puregrade_post/internal/repository"
)

type PostService struct {
	repos *repository.Repository
}

func NewPostService(repos *repository.Repository) Post {
	return &PostService{repos: repos}
}

func (s *PostService) CreatePost(post *entities.Post) (string, error) {
	post.ID = uuid.New().String()
	post.CreatedAt = time.Now()
	if err := s.repos.Create(post); err != nil {
		return "", err
	}

	err := s.repos.Notifier.Push(post.ID, post.AuthorID)
	if err != nil {
		log.Println("An error was received while executing Notifier.Push(): ", err)
	}

	return post.ID, err
}

func (s *PostService) GetPost(id string) (*entities.Post, error) {
	return s.repos.Get(id)
}

func (s *PostService) DeletePost(id string) error {
	return s.repos.Delete(id)
}
