package services

import (
	"github.com/dimastadeoo/backend1/internal/models"
	"github.com/dimastadeoo/backend1/internal/repo"
	"golang.org/x/crypto/bcrypt"
)


type UserService struct{
	repo *repo.UserRepo
}

func NewServiceUser (repo *repo.UserRepo) *UserService{
	return &UserService{repo: repo}
}

func (r *UserService) Register(user *models.RegisterUsers) (models.Users, error){
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
    )

    if err != nil {
        return models.Users{}, err
    }

    user.Password = string(hash)

	return r.repo.Create(user)
}