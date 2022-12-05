package task

import (
	"vpn-operator/domain/operator/job"
)

var tasks = []func(){
	job.InitHealthChecks,
	job.InitUpdateAuthKey,
	job.InitDigitalOceanHealthCheck,
}

func InitTasks() {
	for _, task := range tasks {
		go task()
	}
}
