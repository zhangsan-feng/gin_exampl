package config

import (
	"gin_exampl/src/cache"
	"gin_exampl/src/cron_job"
	"gin_exampl/src/data_store"
	"gin_exampl/src/file_store"
)

type ServerConfigImpl struct {
	Port           string
	LoggerPath     string
	DB             string
	Redis          string
	Minio          string
	Mq             string
	ConfigFilePath string
}

var ServerConfig *ServerConfigImpl

func initServerConfig() {
	ServerConfig = &ServerConfigImpl{
		Port:           "0.0.0.0:9999",
		LoggerPath:     `./`,
		DB:             "",
		Redis:          "",
		Minio:          "",
		Mq:             "",
		ConfigFilePath: "./config.yaml",
	}
}

func loadConfigFileYml() {

}

var GDB = data_store.InitGormStore()
var GCache = cache.InitRedis()
var GFileStore = file_store.InitMinio()

func InitConfig() {
	initServerConfig()
	loadConfigFileYml()
	cron_job.InitCron()
}
