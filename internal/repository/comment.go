package repository

import (
	"fmt"
	"strings"

	"github.com/dhevve/blog/internal/model"
	"github.com/jmoiron/sqlx"
)

type CommentRepository struct {
	db *sqlx.DB
}

func NewCommentRepository(db *sqlx.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (r *CommentRepository) CreateComment(comment model.Comment) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (user_id, post_id, comment) VALUES ($1, $2, $3) RETURNING id", commentsTable)
	row := r.db.QueryRow(query, comment.UserId, comment.PostId, comment.Text)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *CommentRepository) GetComment(commentId int) (model.Comment, error) {
	var comment model.Comment

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", commentsTable)
	err := r.db.Get(&comment, query, commentId)

	return comment, err
}

func (r *CommentRepository) GetComments() ([]model.Comment, error) {
	var comments []model.Comment

	query := fmt.Sprintf("SELECT * FROM %s", commentsTable)
	err := r.db.Select(&comments, query)

	return comments, err
}

func (r *CommentRepository) DeleteComment(commentId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", commentsTable)
	_, err := r.db.Exec(query, commentId)
	return err
}

func (r *CommentRepository) UpdateComment(commentId int, input model.UpdateComment) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argsId := 1

	if input.Text != nil {
		setValues = append(setValues, fmt.Sprintf("comment=$%d", argsId))
		args = append(args, *input.Text)
		argsId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE posts_comments SET %s WHERE id = $%d", setQuery, argsId)

	args = append(args, commentId)

	_, err := r.db.Exec(query, args...)

	return err
}
