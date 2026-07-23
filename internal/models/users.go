package models

import "time"

type Users struct {
	Id        int       `json:"id"`
	Fullname  string    `json:"fullname"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Picture   *string   `json:"picture"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedBy *int      `json:"created_by"`
}

type RegisterUsers struct {
	Fullname string `json:"fullname" form:"fullname" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"-" form:"password" binding:"required"`
	CreatedBy *int
}

type UpdateUser struct {
	Fullname string `json:"fullname" form:"fullname" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
}

type UpdatePicture struct {
	Picture string `json:"picture" form:"picture"`
}

type LoginUser struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"-" form:"password" binding:"required"`
}

type Sort struct {
	Column string
	Order  string
}
