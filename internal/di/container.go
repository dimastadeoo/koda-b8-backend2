package di

import (
	"github.com/dimastadeoo/backend1/internal/handler"
	"github.com/dimastadeoo/backend1/internal/models"
	"github.com/dimastadeoo/backend1/internal/repo"
	"github.com/dimastadeoo/backend1/internal/services"
)

type Container struct {
	userModels  *[]models.Users
	userRepo    *repo.UserRepo
	userService *services.UserService
	userHandler *handler.UserHandler
}

func (c *Container) initDeps() {
	c.userRepo = repo.NewUserRepo(c.userModels)
	c.userService = services.NewServiceUser(c.userRepo)
	c.userHandler = handler.NewHandlerUser(c.userService)
}

func (c *Container) Users() *handler.UserHandler {
	return c.userHandler
}

func NewContainer() (*Container, error) {

	users := []models.Users{}

	container := &Container{
		userModels: &users,
	}

	container.initDeps()

	return container, nil
}
