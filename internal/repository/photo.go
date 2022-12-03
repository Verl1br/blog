package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type PhotoRepository struct {
	db *sqlx.DB
}

func NewPhotoRepository(db *sqlx.DB) *PhotoRepository {
	return &PhotoRepository{db: db}
}

func (r *PhotoRepository) Upload(postId int, fullFileName string) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (photo_location, post_id) VALUES ($1, $2) RETURNING id", photoTable)

	row := r.db.QueryRow(query, fullFileName, postId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *PhotoRepository) DeletePhoto(photoId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", photoTable)
	_, err := r.db.Exec(query, photoId)
	return err
}
