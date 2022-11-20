package database

import (
	"log"
	"vpn-operator/config"
	"vpn-operator/domain/operator/model"
)

func RunMigrations() {
	var err error

	err = config.DB.AutoMigrate(&model.Instance{})
	err = config.DB.AutoMigrate(&model.Client{})

	if err != nil {
		log.Fatal(err)
	}
}
