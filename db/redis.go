package db

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	ctx context.Context
	rdb *redis.Client
}

func (rc *RedisClient) RedisInit(){
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("error loading .env file: %w", err)
	}

	rc.ctx = context.Background()
	rc.rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_CLIENT_ADDR"),
		Username: "default",
		Password: os.Getenv("REDIS_CLIENT_PASS"),
		DB:       0,
	})

}

func (rc *RedisClient) SetVal(longURL string){

}

