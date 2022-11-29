package repository

import (
	"github.com/dhevve/blog/internal/model"
	"github.com/jmoiron/sqlx"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

const (
	usersTable    = "users"
	postTable     = "posts"
	commentsTable = "posts_comments"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(email string) (model.User, error)
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

type Repository struct {
	Authorization
	Post
	Сomment
}

func NewRepository(db *sqlx.DB, driver neo4j.Driver) *Repository {
	return &Repository{
		Authorization: NewAuthorizationRepository(db),
		Post:          NewPostRepository(db),
		Сomment:       NewCommentRepository(db),
	}
}
