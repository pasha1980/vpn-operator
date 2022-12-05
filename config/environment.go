package config

import (
	"github.com/joho/godotenv"
	"os"
)

type environmentConfig struct {
	HttpAddress string

	StoragePath string

	DatabaseLink string

	RedisLink     string
	RedisPassword string
	RedisDatabase string

	DigitalOceanToken string
	DigitalOceanTag   string
}

var Config *environmentConfig

func initEnvironment() {
	_ = godotenv.Load(".env.local")

	Config = &environmentConfig{
		HttpAddress: os.Getenv("HTTP_ADDRESS"),

		StoragePath: os.Getenv("STORAGE_PATH"),

		DatabaseLink: os.Getenv("DATABASE_DSN"),

		RedisLink:     os.Getenv("REDIS_HOST"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
		RedisDatabase: os.Getenv("REDIS_DB"),

		DigitalOceanToken: os.Getenv("DO_TOKEN"),
		DigitalOceanTag:   os.Getenv("DO_DROPLET_TAG"),
	}
}
