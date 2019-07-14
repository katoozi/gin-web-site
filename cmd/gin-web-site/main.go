package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"github.com/katoozi/gin-web-site/configs"
	"github.com/katoozi/gin-web-site/internal/app/website"
	"github.com/katoozi/gin-web-site/internal/pkg/models"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"
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

func main() {
	// connect to postgresql
	databaseConfig := fetchDatabaseConfig()
	dbConnectionStr := fmt.Sprintf(
		"user=%s dbname=%s password=%s sslmode=disable",
		databaseConfig.User,
		databaseConfig.DatabaseName,
		databaseConfig.Password,
	)
	db, err := sqlx.Connect("postgres", dbConnectionStr)
	if err != nil {
		log.Fatalf("Connect to db Failed: %v", err)
	}
	models.MigrateTables(db)
	website.DbCon = db

	// connect to redis
	redisConfig := fetchRedisConfig()
	redisClient := redis.NewClient(&redis.Options{
		Addr:     configs.GetAddr(redisConfig),
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})
	_, err = redisClient.Ping().Result()
	if err != nil {
		log.Fatalf("Error while connect to redis: %v\n", err)
	}
	website.RedisCon = redisClient

	// load static files, templates, register routers. start server
	r := gin.Default()
	r.Static("/static", "./web/assets")
	// r.StaticFS("/more_static", http.Dir("my_file_system"))
	// r.StaticFile("/favicon.ico", "./resources/favicon.ico")

	website.RegisterTemplateFuncs(r)

	// load html files
	// r.LoadHTMLGlob("./web/templates/components/*")
	r.LoadHTMLGlob("./web/templates/*.html")
	//r.LoadHTMLFiles("templates/template1.html", "templates/template2.html")

	website.RegisterRoutes(r)

	// fetch server configs from config.yaml file
	serverConfig := fetchServerConfig()
	r.Run(configs.GetAddr(serverConfig))
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

func fetchRedisConfig() *configs.RedisConfig {
	redisConfig := &configs.RedisConfig{}
	viper.UnmarshalKey("redis", &redisConfig)
	return redisConfig
}
