package main

import (
	"log"
	"vpn-operator/config"
	"vpn-operator/domain/infrastructure/api"
	"vpn-operator/domain/infrastructure/task"
)

func main() {
	config.InitConfig()

	err := api.InitHttp()
	if err != nil {
		log.Fatal(err)
	}

	task.InitTasks()
}
