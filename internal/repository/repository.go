package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type Repository struct {
}

func NewRepository(db *sqlx.DB, driver neo4j.Driver) *Repository {
	return &Repository{}
}
