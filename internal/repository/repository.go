package repository

import (
	"github.com/dhevve/blog/internal/model"
	"github.com/jmoiron/sqlx"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(email string) (model.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB, driver neo4j.Driver) *Repository {
	return &Repository{
		Authorization: NewAuthorizationRepository(db),
	}
}
