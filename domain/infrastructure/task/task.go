package task

var tasks = []func(){}

func InitTasks() {
	for _, job := range tasks {
		go job()
	}
}
