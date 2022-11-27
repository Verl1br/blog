package service

import (
	"github.com/dhevve/blog/internal/model"
	"github.com/dhevve/blog/internal/repository"
)

type PostService struct {
	repo repository.Post
}

func NewPostService(repo repository.Post) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) CreatePost(post model.Post) (int, error) {
	return s.repo.CreatePost(post)
}

func (s *PostService) GetPost(postId, userId int) (model.Post, error) {
	return s.repo.GetPost(postId, userId)
}

func (s *PostService) GetPosts(userId int) ([]model.Post, error) {
	return s.repo.GetPosts(userId)
}

func (s *PostService) DeletePost(postId int) error {
	return s.repo.DeletePost(postId)
}

func (s *PostService) UpdatePost(postId int, input model.UpdatePost) error {
	return s.repo.UpdatePost(postId, input)
}
