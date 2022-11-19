package config

func InitConfig() {
	initEnvironment()
	initDatabase()
	initRedis()
}
