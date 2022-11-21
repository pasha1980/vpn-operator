package task

import (
	"time"
	"vpn-operator/domain/operator/job"
)

var tasks = []func(){
	job.InitHealthChecks,
}

func InitTasks() {
	for range time.Tick(time.Minute) {
		for _, task := range tasks {
			go task()
		}
	}
}
