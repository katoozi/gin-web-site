package commands

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/katoozi/gin-web-site/configs"
	"github.com/katoozi/gin-web-site/internal/app/website"

	"github.com/spf13/cobra"
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
	// connect to redis
	redisConfig := fetchRedisConfig()
	// basic redis connection
	// redisClient := redis.NewClient(&redis.Options{
	// 	Addr:     configs.GetAddr(redisConfig),
	// 	Password: redisConfig.Password,
	// 	DB:       redisConfig.DB,
	// })

	// use redis sentinel for high availability and failover
	addrs := os.Getenv("REDIS_SENTINEL_INSTANCES")
	sentinelAddrs := strings.Split(addrs, ",")
	redisClient := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "mymaster",
		SentinelAddrs: sentinelAddrs,
		Password:      redisConfig.Password,
		DB:            redisConfig.DB,
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
	// r.StaticFS("/more_static", http.Dir("my_file_system"))
	// r.StaticFile("/favicon.ico", "./resources/favicon.ico")

	website.RegisterTemplateFuncs(r)

	// load html files
	// r.LoadHTMLGlob("./web/templates/components/*")
	r.LoadHTMLGlob("./web/build/*.html")
	//r.LoadHTMLFiles("templates/template1.html", "templates/template2.html")

	website.RegisterRoutes(r)

	// fetch server configs from config.yaml file
	serverConfig := fetchServerConfig()
	r.Run(configs.GetAddr(serverConfig))
	fmt.Println("Start Listning...")
}
