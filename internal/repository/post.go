package repository

import (
	"fmt"
	"strings"

	"github.com/dhevve/blog/internal/model"
	"github.com/jmoiron/sqlx"
)

type PostRepository struct {
	db *sqlx.DB
}

func NewPostRepository(db *sqlx.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) CreatePost(post model.Post) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (title, content, user_id) VALUES ($1, $2, $3) RETURNING id", "posts")

	row := r.db.QueryRow(query, post.Title, post.Content, post.UserId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *PostRepository) GetPost(postId, userId int) (model.Post, error) {
	var post model.Post

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1 AND user_id = $2", "posts")
	err := r.db.Get(&post, query, postId, userId)
	return post, err
}

func (r *PostRepository) GetPosts(userId int) ([]model.Post, error) {
	var posts []model.Post

	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1", "posts")
	err := r.db.Select(&posts, query, userId)
	return posts, err
}

func (r *PostRepository) DeletePost(postId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", "posts")
	_, err := r.db.Exec(query, postId)
	return err
}

func (r *PostRepository) UpdatePost(postId int, input model.UpdatePost) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Content != nil {
		setValues = append(setValues, fmt.Sprintf("content=$%d", argId))
		args = append(args, *input.Content)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE posts SET %s WHERE id = $%d", setQuery, argId)

	args = append(args, postId)

	_, err := r.db.Exec(query, args...)
	return err
}
