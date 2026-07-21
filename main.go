package main

import (
	"log"
	"os"

	"github.com/dimastadeoo/backend1/internal/di"
	"github.com/dimastadeoo/backend1/internal/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	container, err := di.NewContainer()
	if err != nil {
		log.Fatal(err)
	}
	user := container.Users()

	r := gin.Default()
	r.Use(middlewares.Cors())

	{
		auth := r.Group("/auth")
		auth.POST("/register", user.Register)
		auth.POST("/login", user.Login)
	}

	{
		users := r.Group("/users")
		users.Use(middlewares.Auth())
		users.GET("", user.GetAll)
		users.POST("", user.Register)
		users.GET("/:id", user.FindById)
		users.PATCH("/:id", user.Update)
		users.DELETE("/:id", user.Delete)

	}

	getPort := os.Getenv("PORT")
	r.Run("0.0.0.0:" + getPort)
}
