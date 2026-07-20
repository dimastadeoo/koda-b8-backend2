package repo

import (
	"errors"
	"time"

	"github.com/dimastadeoo/backend1/internal/models"
)

type UserRepo struct {
	data *[]models.Users
}

func NewUserRepo(data *[]models.Users) *UserRepo {
	return &UserRepo{data: data}
}

func (u *UserRepo) Create(create *models.RegisterUsers) (models.Users, error) {

	for _, user := range *u.data {
		if user.Email == create.Email {
			return models.Users{}, errors.New("email already exists")
		}
	}
	user := models.Users{
		Id:        len(*u.data) + 1,
		Fullname:  create.Fullname,
		Email:     create.Email,
		Password:  create.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	*u.data = append(*u.data, user)

	return user, nil
}

func (r *UserRepo) GetAll() []models.Users {
    return *r.data
}