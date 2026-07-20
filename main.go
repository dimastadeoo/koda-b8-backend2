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

	r := gin.Default()

	r.POST("/users", container.Users().Register)
	r.GET("/users", container.Users().GetAll)
	r.POST("/login", container.Users().Login)

	r.Run("0.0.0.0:8080")
}
