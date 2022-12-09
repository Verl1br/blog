package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/dhevve/blog/internal/model"
	"github.com/dhevve/blog/internal/repository"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

type NewsFeedService struct {
	repo        repository.NewsFeed
	redisClient *redis.Client
}

func NewNewsFeedService(repo repository.NewsFeed, redisClient *redis.Client) *NewsFeedService {
	return &NewsFeedService{
		repo:        repo,
		redisClient: redisClient,
	}
}

func (s *NewsFeedService) GetNews(id int) ([]model.Post, error) {
	var news []model.Post

	//stringId := strconv.Itoa(id)
	stringId := fmt.Sprintf("news_%d", id)

	ctx := context.Background()
	cachedNews, err := s.redisClient.Get(ctx, stringId).Bytes()
	if err != nil {
		dbNews, err := s.repo.GetNews(id)
		if err != nil {
			return nil, err
		}

		cachedNews, err := json.Marshal(dbNews)
		if err != nil {
			return nil, err
		}

		err = s.redisClient.Set(ctx, stringId, cachedNews, 60*time.Second).Err()
		if err != nil {
			return nil, err
		}

		logrus.Info("REDIS!")

		return dbNews, nil
	}

	err = json.Unmarshal(cachedNews, &news)
	if err != nil {
		return nil, err
	}

	return news, nil
}
