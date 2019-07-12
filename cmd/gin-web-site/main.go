package main

import (
	"fmt"
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
}

func main() {
	// fetch server and database configs from config.yaml file
	serverConfig := fetchServerConfig()
	databaseConfig := fetchDatabaseConfig()

	fmt.Println("database config: ", databaseConfig)

	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.Run(serverConfig.GetAddr())
	fmt.Println("Start Listning...")
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
