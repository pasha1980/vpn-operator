package config

import (
	"github.com/joho/godotenv"
	"os"
)

type EnvironmentConfig struct {
	HttpAddress string

	StoragePath string

	DatabaseLink string

	RedisLink     string
	RedisPassword string
	RedisDatabase string
}

var Config *EnvironmentConfig

func initEnvironment() {
	_ = godotenv.Load(".env.local")

	Config = &EnvironmentConfig{
		HttpAddress: os.Getenv("HTTP_ADDRESS"),

		StoragePath: os.Getenv("STORAGE_PATH"),

		DatabaseLink: os.Getenv("DATABASE_DSN"),

		RedisLink:     os.Getenv("REDIS_HOST"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
		RedisDatabase: os.Getenv("REDIS_DB"),
	}
}
