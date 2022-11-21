package config

import (
	"log"
	"vpn-operator/domain/infrastructure/utils"
)

func initAuth() {
	newToken := utils.GenerateRandomString(500)
	_, err := Redis.Set(Context, "secret_key", newToken, 0).Result()
	if err != nil {
		log.Fatal(err)
	}
}
