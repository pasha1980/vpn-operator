package auth

import (
	"vpn-operator/config"
	"vpn-operator/domain/infrastructure/utils"
)

func GenerateNewApiKey() error {
	newToken := utils.GenerateRandomString(config.SecretKeyLength)
	_, err := config.Redis.Set(config.Context, "secret_key", newToken, 0).Result()
	return err
}

func CheckApiToken(token string) (bool, error) {
	savedToken, err := config.Redis.Get(config.Context, "secret_key").Result()
	if err != nil {
		return false, err
	}

	return savedToken == token, nil
}
