package services

import (
	"errors"

	"github.com/dimastadeoo/backend1/internal/models"
	"github.com/dimastadeoo/backend1/internal/repo"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repo.UserRepo
}

func NewServiceUser(repo *repo.UserRepo) *UserService {
	return &UserService{repo: repo}
}

func (r *UserService) Register(user *models.RegisterUsers) (models.Users, error) {
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

func (s *UserService) GetAll() []models.Users {
    return s.repo.GetAll()
}

func (s *UserService) Login(req *models.LoginUser) (*models.Users, error) {

	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("email atau password salah")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	)

	if err != nil {
		return nil, errors.New("email atau password salah")
	}

	return user, nil
}
