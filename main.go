package main

import (
	"log"
	"vpn-operator/config"
	"vpn-operator/domain/infrastructure/api"
	"vpn-operator/domain/infrastructure/database"
	"vpn-operator/domain/infrastructure/task"
)

func main() {
	config.InitConfig()

	database.RunMigrations()

	go task.InitTasks()

	err := api.InitHttp()
	if err != nil {
		log.Fatal(err)
	}

}
