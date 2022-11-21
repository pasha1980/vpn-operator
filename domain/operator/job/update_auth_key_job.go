package job

import (
	"time"
	"vpn-operator/domain/infrastructure/auth"
)

func InitUpdateAuthKey() {
	for range time.Tick(time.Hour) {
		auth.GenerateNewApiKey()
	}
}
