package config

import (
	"log"
	"vpn-operator/domain/infrastructure/auth"
)

func initAuth() {
	err := auth.GenerateNewApiKey()
	if err != nil {
		log.Fatal(err)
	}
}
