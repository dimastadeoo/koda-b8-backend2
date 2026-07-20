package models

import "time"

type Users struct {
	Id        int `json:"id"`
	Fullname  string `json:"fullname"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type RegisterUsers struct {
	Fullname  string `json:"fullname" form:"fullname" binding:"required"`
	Email     string `json:"email" form:"email" binding:"required,email"`
	Password  string `json:"-" form:"password" binding:"required"`
}

type LoginUser struct {
	Email     string `json:"email" form:"email" binding:"required,email"`
	Password  string `json:"-" form:"password" binding:"required"`
}
