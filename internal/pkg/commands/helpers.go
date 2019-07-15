package commands

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/katoozi/gin-web-site/configs"
	"github.com/katoozi/gin-web-site/internal/pkg/auth"
	"github.com/katoozi/gin-web-site/internal/app/website"
	"github.com/spf13/viper"

	_ "github.com/lib/pq" // register postgresql driver
)

var dbCon *sqlx.DB

func init() {
	configs.SetDefaultValues()

	viper.SetConfigName("config")
	viper.SetConfigFile("configs/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config.yaml file: %v", err)
	}

	gin.SetMode(gin.DebugMode)

	// connect to postgresql
	databaseConfig := fetchDatabaseConfig()
	db, err := sqlx.Connect("postgres", configs.GetAddr(databaseConfig))
	if err != nil {
		log.Fatalf("Connect to db Failed: %v", err)
	}
	dbCon = db
	auth.MigrateTables(db)
	website.DbCon = db

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
