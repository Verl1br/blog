package service

import (
	"github.com/dhevve/blog/internal/model"
	"github.com/dhevve/blog/internal/repository"
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
	GetComments() ([]model.Comment, error)
	DeleteComment(commentId int) error
	UpdateComment(commentId int, input model.UpdateComment) error
}

type Service struct {
	Authorization
	Post
	Сomment
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthorizationService(repo),
		Post:          NewPostService(repo),
		Сomment:       NewCommentService(repo),
	}
}
