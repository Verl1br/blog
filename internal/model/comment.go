package model

type Comment struct {
	Id     int    `json:"-" db:"id"`
	UserId int    `json:"user_id" validate:"required" db:"user_id"`
	PostId int    `json:"post_id" validate:"required" db:"post_id"`
	Text   string `json:"text" validate:"required" db:"comment"`
}

type UpdateComment struct {
	Text *string `json:"text" validate:"required" db:"comment"`
}
