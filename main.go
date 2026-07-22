package main

import (
	"log"
	"os"

	"github.com/dimastadeoo/backend1/docs"
	"github.com/dimastadeoo/backend1/internal/di"
	"github.com/dimastadeoo/backend1/internal/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @contact.name
// @contact.email  dimastadeoo@gmail.com

// @securityDefinitions.apikey Bearer
// @in                         header
// @name                       Authorization
// @description                Type "Bearer " followed by a space and your JWT token.
func main() {
	godotenv.Load()
	container, err := di.NewContainer()
	if err != nil {
		log.Fatal(err)
	}
	user := container.Users()

	// // programmatically set swagger info
	docs.SwaggerInfo.Title = "Backend CRUD"
	docs.SwaggerInfo.Description = "This is my first Backend CRUD"
	docs.SwaggerInfo.Version = "1.0.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := gin.Default()
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Static("/uploads", "./uploads")
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
		users.POST("", user.RegisterAdmin)
		users.GET("/:id", user.FindById)
		users.PATCH("/:id", user.Update)
		users.PATCH("/:id/picture", user.UpdatePicture)
		users.DELETE("/:id", user.Delete)
	}

	getPort := os.Getenv("PORT")
	r.Run("0.0.0.0:" + getPort)
}
