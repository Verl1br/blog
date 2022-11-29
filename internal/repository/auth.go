package repository

import (
	"fmt"

	"github.com/dhevve/blog/internal/model"
	"github.com/jmoiron/sqlx"
)

type AuthorizationRepository struct {
	db *sqlx.DB
}

func NewAuthorizationRepository(db *sqlx.DB) *AuthorizationRepository {
	return &AuthorizationRepository{db: db}
}

func (r *AuthorizationRepository) CreateUser(user model.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (first_name, last_name, email, password_hash) VALUES ($1, $2, $3, $4) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.FirstName, user.LastName, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthorizationRepository) GetUser(email string) (model.User, error) {
	var user model.User

	query := fmt.Sprintf("SELECT id, password_hash FROM %s WHERE email = $1", usersTable)
	err := r.db.Get(&user, query, email)

	return user, err
}
