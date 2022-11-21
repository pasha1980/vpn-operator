package task

import (
	"vpn-operator/domain/operator/job"
)

var tasks = []func(){
	job.InitHealthChecks,
	job.InitUpdateAuthKey,
}

func InitTasks() {
	for _, task := range tasks {
		go task()
	}
}
