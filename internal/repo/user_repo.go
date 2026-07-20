package repo

import (
	"context"

	"github.com/dimastadeoo/backend1/internal/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepo struct {
	data *pgxpool.Pool
}

func NewUserRepo(data *pgxpool.Pool) *UserRepo {
	return &UserRepo{data: data}
}

func (u *UserRepo) Create(create *models.RegisterUsers) (models.Users, error) {

	query := `
			INSERT INTO users(fullname, email, password)
			VALUES($1, $2, $3)
			RETURNING id, fullname, email, password, created_at, updated_at;
	`
	var user models.Users

	err := u.data.QueryRow(
		context.Background(), query,
		create.Fullname,
		create.Email,
		create.Password,
	).Scan(
		&user.Id,
		&user.Fullname,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return models.Users{}, err
	}

	return user, nil
}

func (u *UserRepo) GetAll() ([]models.Users, error) {
    query := `
			SELECT id, fullname, email, created_at, updated_at
			FROM users
	`
	data, err := u.data.Query(context.Background(), query)
	if err != nil{
		return nil, err
	}

	defer data.Close()

	users := []models.Users{}

	for data.Next(){
		var user models.Users

		err := data.Scan(
			&user.Id,
			&user.Fullname,
			&user.Email,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, data.Err()
}

func (u *UserRepo) FindByEmail(email string) (*models.Users, error) {
	var user models.Users

	query := `
			SELECT id, fullname, email, password, created_at, updated_at
			FROM users
			WHERE email=$1
	`

	err := u.data.QueryRow(
		context.Background(),
		query,
		email,
	).Scan(
		&user.Id,
		&user.Fullname,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows{
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}