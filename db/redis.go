package db

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	hash "github.com/unnxt30/zippy/hash"
)

type RedisClient struct {
	ctx context.Context
	rdb *redis.Client
}

func (rc *RedisClient) Init() RedisClient{
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

	return *rc

}

func (rc *RedisClient) RedisGet(longURL string)(string,error) {
	val, err := rc.rdb.Get(rc.ctx, longURL).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}

func (rc *RedisClient) RedisSet(longURL string) (string,error){
	shortenedURL := hash.GetShortURL(5)

	if _, err := rc.rdb.Get(rc.ctx, longURL).Result(); err == redis.Nil {
		rc.rdb.Set(rc.ctx, longURL, shortenedURL, 0)
		return shortenedURL, nil
	}

	val, _ := rc.rdb.Get(rc.ctx, longURL).Result()

	return val, nil
}



