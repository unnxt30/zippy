package db

import (
	"context"
	"fmt"
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

func (rc *RedisClient) RedisGet(shortURL string)(string,error) {
	var cursor uint64
	var found string
	
	for{
		keys, cursor, err := rc.rdb.Scan(rc.ctx, cursor, "*", 10).Result()
		if err != nil {
			return "", err
		}

		for _, key := range keys {
			val, err := rc.rdb.Get(rc.ctx, key).Result()
			if err == redis.Nil   {
				return "", fmt.Errorf("key does not exist")
			}else if err != nil {
				return "", err
			}

			if val == shortURL {
				found = key
				return found, nil 
			}

		}

		if cursor == 0 {
			break
		}
	}

	return "", fmt.Errorf("key does not exist")
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



