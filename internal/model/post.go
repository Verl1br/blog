package model

import "time"

type Post struct {
	Id          int       `json:"-" db:"id"`
	Title       string    `json:"title" validate:"required" db:"title"`
	Content     string    `json:"content" validate:"required" db:"content"`
	PublishedAt time.Time `json:"published_at" validate:"email,required" db:"published_at"`
	UserId      int       `json:"user_id" validate:"required" db:"user_id"`
}

type UpdatePost struct {
	Title   *string `json:"title" validate:"required"`
	Content *string `json:"content" validate:"required"`
}
