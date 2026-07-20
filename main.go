package main

import (
	"log"

	"github.com/dimastadeoo/backend1/internal/di"
	"github.com/gin-gonic/gin"
)

func main() {
	container, err := di.NewContainer()
	if err != nil {
		log.Fatal(err)
	}
	user := container.Users()

	r := gin.Default()

	r.POST("/users", user.Register)
	r.GET("/users", user.GetAll)
	r.POST("/login", user.Login)

	r.Run("0.0.0.0:8080")
}
