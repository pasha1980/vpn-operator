package job

import (
	"time"
	"vpn-operator/config"
	"vpn-operator/domain/operator/model"
)

func InitHealthChecks() {
	for range time.Tick(20 * time.Second) {
		var instances []model.Instance
		config.DB.Find(&instances)
		for _, instance := range instances {
			go healthCheck(instance)
		}
	}
}

func healthCheck(instance model.Instance) {
	defer config.DB.Save(&instance)

	pong := instance.Ping()
	if !pong {
		instance.IsActive = false
		return
	}

	instance.IsActive = true

	instance.GetStatus()
}
