package middlewares

import (
	"net/http"
	"os"
	"strings"

	"github.com/dimastadeoo/backend1/internal/lib"
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
		authHeader := ctx.GetHeader("Authorization")
		prefix := "Bearer "
		if !strings.HasPrefix(authHeader, prefix){
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
		token, _ := strings.CutPrefix(authHeader, prefix)

		if isValid, userId := lib.VerifyToken(token); isValid {
			ctx.Set("userId", userId)
			ctx.Next()
			return 
		}
		ctx.AbortWithStatus(http.StatusUnauthorized)

	}
}