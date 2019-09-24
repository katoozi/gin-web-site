package commands

import (
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/katoozi/gin-web-site/internal/app/website"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RunServerCommand is a cobra command that start web server
var RunServerCommand = &cobra.Command{
	Use:   "runserver",
	Short: "start web server",
	Long:  `start web server. configuration founds in config.yaml file`,
	Run:   runServer,
}

func runServer(cmd *cobra.Command, args []string) {
	Init()

	// use redis sentinel for high availability and failover
	addrs := viper.GetString("redis.sentinels")
	db := viper.GetInt("redis.db")
	pass := viper.GetString("redis.pass")

	sentinelAddrs := strings.Split(addrs, ",")
	redisClient := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "mymaster",
		SentinelAddrs: sentinelAddrs,
		Password:      pass,
		DB:            db,
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Fatalf("Error while connect to redis: %v\n", err)
	}
	website.RedisCon = redisClient

	// load static files, templates, register routers. start server
	r := gin.Default()
	r.Static("/static", "./web/build/static")
	r.StaticFile("/manifest.json", "./web/build/manifest.json")

	website.RegisterTemplateFuncs(r)

	// load html files
	r.LoadHTMLGlob("./web/build/*.html")

	website.RegisterRoutes(r)

	host := viper.GetString("server.host")
	port := viper.GetString("server.port")
	if port == "" {
		port = "8000"
	}
	r.Run(host + ":" + port)
	fmt.Println("Start Listning...")
}
