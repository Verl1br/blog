package model

import "time"

type Post struct {
	Id          int       `json:"-" db:"id"`
	Title       string    `json:"title" db:"title"`
	Content     string    `json:"content" db:"content"`
	PublishedAt time.Time `json:"published_at" db:"published_at"`
	UserId      int       `json:"user_id" db:"user_id"`
}

type UpdatePost struct {
	Title   *string `json:"title" validate:"required"`
	Content *string `json:"content" validate:"required"`
}
