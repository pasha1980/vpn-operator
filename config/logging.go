package config

import "time"

var Log *redisLogger

func initLogging() {
	Log = &redisLogger{}
}

type redisLogger struct {
}

func (l redisLogger) Write(message string, logType string) {
	logs, _ := Redis.Get(Context, "operator_log").Result()
	logs = logs + "  ||  " + time.Now().Format("2006-01-02 15:04:05") + " - [" + logType + "] " + message
	Redis.Set(Context, "operator_log", logs, 0)
}
