package repository

import (
	"github.com/dhevve/blog/internal/model"
	"github.com/jmoiron/sqlx"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type Post interface {
	CreatePost(post model.Post) (int, error)
	GetPost(postId, userId int) (model.Post, error)
	GetPosts(userId int) ([]model.Post, error)
	DeletePost(postId int) error
	UpdatePost(postId int, input model.UpdatePost) error
}

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(email string) (model.User, error)
}

type Repository struct {
	Authorization
	Post
}

func NewRepository(db *sqlx.DB, driver neo4j.Driver) *Repository {
	return &Repository{
		Authorization: NewAuthorizationRepository(db),
		Post:          NewPostRepository(db),
	}
}
