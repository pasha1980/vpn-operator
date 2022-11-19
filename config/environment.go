package config

import "os"

type EnvironmentConfig struct {
	HttpAddress string

	DatabaseLink string

	RedisLink     string
	RedisPassword string
	RedisDatabase string
}

var Config *EnvironmentConfig

func initEnvironment() {
	Config = &EnvironmentConfig{
		HttpAddress: os.Getenv("HTTP_ADDRESS"),

		DatabaseLink: os.Getenv("DATABASE_DSN"),

		RedisLink:     os.Getenv("REDIS_HOST"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
		RedisDatabase: os.Getenv("REDIS_DB"),
	}
}
