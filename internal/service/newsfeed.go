package service

import (
	"github.com/dhevve/blog/internal/model"
	"github.com/dhevve/blog/internal/repository"
)

type NewsFeedService struct {
	repo repository.NewsFeed
}

func NewNewsFeedService(repo repository.NewsFeed) *NewsFeedService {
	return &NewsFeedService{repo: repo}
}

func (s *NewsFeedService) GetNews(id int) ([]model.Post, error) {
	return s.repo.GetNews(id)
}
