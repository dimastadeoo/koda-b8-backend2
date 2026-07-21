package middlewares

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func Cors() gin.HandlerFunc{
	return func(ctx *gin.Context)  {
		godotenv.Load()
		getPort := os.Getenv("PORT_FRONTEND")
		ctx.Header("Access-Control-Allow-Origin", "http://localhost:" + getPort)
		ctx.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PATCH, DELETE")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Origin, Authorization")
		if ctx.Request.Method == "OPTIONS"{
			ctx.Status(http.StatusNoContent)
			return 
		}
		ctx.Next()
	}
}

func Auth() gin.HandlerFunc{
	return func(ctx *gin.Context)  {
		token := ctx.GetHeader("Authorization")

		if token != "hello" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return 
		}
		ctx.Next()
	}
}