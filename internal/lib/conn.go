package lib

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func Conn() (*pgxpool.Pool, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Cannot read file ekstensions .env")
	}

	getCon := os.Getenv("DATABASE_URL")

	pool, err := pgxpool.New(context.Background(), getCon)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(context.Background()); err != nil{
		pool.Close()
		return nil, err
	}

	return pool, nil
}
