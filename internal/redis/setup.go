package redis

import (
	"os"

	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client

func ConnectRedis() {
	redisHost := os.Getenv("REDIS_HOST")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisUsername := os.Getenv("REDIS_USERNAME")

	Rdb = redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: redisPassword,
		DB:       0, // use default DB
		Username: redisUsername,
	})
}
