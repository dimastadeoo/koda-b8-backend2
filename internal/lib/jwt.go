package lib

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)


type JwtToken struct {
	userId int 
	jwt.RegisteredClaims
}

func GeneratedToken(id int)string {
	godotenv.Load()

	// Create claims with multiple fields populated
	claims := JwtToken{
		id,
		jwt.RegisteredClaims{
			ExpiresAt:  jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, _:= token.SignedString([]byte(os.Getenv("JWT_KEY")))
	return "Bearer " + result
}

func VerifyToken(token string) (bool, *int) {
	payload, err := jwt.ParseWithClaims(token, &JwtToken{}, func(token *jwt.Token) (any, error) {
		return []byte(os.Getenv("JWT_KEY")), nil
	})
	result, _ := payload.Claims.(*JwtToken)

	if err != nil{
		return false, nil
	}
	return true, &result.userId
}
