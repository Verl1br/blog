package model

type User struct {
	Id        int    `json:"-" db:"id"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"email,required"`
	Password  string `json:"password" validate:"min=4,max=20" db:"password_hash"`
}

type SingInInput struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"min=4,max=20"`
}
