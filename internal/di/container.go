package di

import (
	"github.com/dimastadeoo/backend1/internal/handler"
	"github.com/dimastadeoo/backend1/internal/lib"
	"github.com/dimastadeoo/backend1/internal/repo"
	"github.com/dimastadeoo/backend1/internal/services"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Container struct {
	userDb  *pgxpool.Pool

	userRepo    *repo.UserRepo
	userService *services.UserService
	userHandler *handler.UserHandler
}

func (c *Container) initDeps() {
	c.userRepo = repo.NewUserRepo(c.userDb)
	c.userService = services.NewServiceUser(c.userRepo)
	c.userHandler = handler.NewHandlerUser(c.userService)
}

func (c *Container) Users() *handler.UserHandler {
	return c.userHandler
}

func NewContainer() (*Container, error) {


	users, err := lib.Conn()

	if err != nil{
		return nil, err
	}

	container := &Container{
		userDb: users,
	}

	container.initDeps()

	return container, nil
}
