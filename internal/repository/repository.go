package repository

import (
	"github.com/dhevve/blog/internal/model"
	"github.com/jmoiron/sqlx"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

const (
	usersTable    = "users"
	postTable     = "posts"
	commentsTable = "posts_comments"
	photoTable    = "posts_photo"
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
	GetComments(id int) ([]model.Comment, error)
	DeleteComment(commentId int) error
	UpdateComment(commentId int, input model.UpdateComment) error
}

type Photo interface {
	Upload(postId int, fullFileName string) (int, error)
	DeletePhoto(photoId int) error
}

type Friend interface {
	GetFriends(id int) []model.User
	CreateFriends(myId, friendId int) error
	DeleteFriend(myId, friendId int) error
}

type NewsFeed interface {
	GetNews(id int) ([]model.Post, error)
}

type Repository struct {
	Authorization
	Post
	Сomment
	Photo
	Friend
	NewsFeed
}

func NewRepository(db *sqlx.DB, driver neo4j.DriverWithContext) *Repository {
	return &Repository{
		Authorization: NewAuthorizationRepository(db),
		Post:          NewPostRepository(db),
		Сomment:       NewCommentRepository(db),
		Photo:         NewPhotoRepository(db),
		Friend:        NewFriendRepository(db, driver),
		NewsFeed:      NewNewsFeedRepository(db, driver),
	}
}
