package service

import (
	"context"

	"github.com/dhevve/blog/internal/model"
	"github.com/dhevve/blog/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	ParseToken(accessToken string) (int, error)
	GenerateToken(email, password string) (string, error)
}

type Post interface {
	CreatePost(post model.Post) (int, error)
	GetPost(postId, userId int) (model.Post, error)
	GetPosts(userId int) ([]model.Post, error)
	DeletePost(postId int) error
	UpdatePost(postId int, input model.UpdatePost) error
}

type Сomment interface {
	CreateComment(comment model.Comment) (int, error)
	GetComment(commentId int) (model.Comment, error)
	GetComments(id int) ([]model.Comment, error)
	DeleteComment(commentId int) error
	UpdateComment(commentId int, input model.UpdateComment) error
}

type Photo interface {
	Upload(c *gin.Context, postId int) (int, error)
	DeletePhoto(photoId int) error
}

type Friend interface {
	GetFriends(id int, ctx context.Context) []model.User
	CreateFriends(myId, friendId int, ctx context.Context) error
	DeleteFriend(myId, friendId int, ctx context.Context) error
}

type NewsFeed interface {
	GetNews(id int, ctx context.Context) ([]model.Post, error)
}

type Service struct {
	Authorization
	Post
	Сomment
	Photo
	Friend
	NewsFeed
}

func NewService(repo *repository.Repository, redisClient *redis.Client) *Service {
	return &Service{
		Authorization: NewAuthorizationService(repo),
		Post:          NewPostService(repo),
		Сomment:       NewCommentService(repo),
		Photo:         NewPhotoService(repo),
		Friend:        NewFriendServce(repo),
		NewsFeed:      NewNewsFeedService(repo, redisClient),
	}
}
