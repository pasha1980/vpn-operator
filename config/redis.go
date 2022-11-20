package config

import (
	"github.com/go-redis/redis/v9"
	"log"
	"strconv"
)

var Redis *redis.Client

func initRedis() {
	database, _ := strconv.Atoi(Config.RedisDatabase)

	Redis = redis.NewClient(&redis.Options{
		Addr:     Config.RedisLink,
		Password: Config.RedisPassword,
		DB:       database,
	})

	_, err := Redis.Ping(Context).Result()
	if err != nil {
		log.Fatal(err)
	}
}
