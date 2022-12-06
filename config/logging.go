package config

import "time"

var Log *redisLogger

func initLogging() {
	Log = &redisLogger{}
}

type redisLogger struct {
}

func (l redisLogger) Write(message string, logType string) {
	logs := Redis.Get(Context, "operator_log").String()
	logs = logs + "\n" + time.Now().String() + " - [" + logType + "] " + message
	Redis.Set(Context, "operator_log", logs, 0)
}
