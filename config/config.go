package config

func InitConfig() {
	initEnvironment()
	initContext()
	initDatabase()
	initRedis()
	initAuth()
}
