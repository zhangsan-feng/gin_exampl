package config

import (
	"gin_exampl/src/cache"
	"gin_exampl/src/data_store"
	"gin_exampl/src/file_store"
)

var GDb = data_store.InitEntStore()
var GCache = cache.InitRedis()
var GFileStore = file_store.InitMinio()

type ServerConfigImpl struct {
	Port string
}

var ServerConfig *ServerConfigImpl

func initLoggerConfig() {
	ServerConfig = &ServerConfigImpl{
		Port: "0.0.0.0:9999",
	}
}

func loadConfigFileYml() {

}

func InitConfig() {
	initLoggerConfig()
	loadConfigFileYml()
}
