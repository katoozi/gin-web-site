package commands

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/katoozi/gin-web-site/configs"
	"github.com/spf13/viper"
)

func init() {
	configs.SetDefaultValues()

	viper.SetConfigName("config")
	viper.SetConfigFile("configs/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config.yaml file: %v", err)
	}

	gin.SetMode(gin.DebugMode)
}

func fetchServerConfig() *configs.ServerConfig {
	serverConfig := &configs.ServerConfig{}
	viper.UnmarshalKey("server", &serverConfig)
	return serverConfig
}

func fetchDatabaseConfig() *configs.DatabaseConfig {
	databaseConfig := &configs.DatabaseConfig{}
	viper.UnmarshalKey("database", &databaseConfig)
	return databaseConfig
}

func fetchRedisConfig() *configs.RedisConfig {
	redisConfig := &configs.RedisConfig{}
	viper.UnmarshalKey("redis", &redisConfig)
	return redisConfig
}
