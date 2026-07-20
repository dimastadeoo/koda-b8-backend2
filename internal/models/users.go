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
	Fullname  string `json:"fullname"`
	Email     string `json:"email"`
	Password  string `json:"-"`
}

type LoginUser struct {
	Email     string `json:"email"`
	Password  string `json:"-"`
}
