package commands

import (
	"fmt"
	"log"

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
	// connect to postgresql
	psDB := initialPostgres()
	defer psDB.Close()

	// connect to redis
	redisClient := initialRedis()
	defer redisClient.Close()

	// initial rabbitMQ
	rabbitClient := initialRabbitmq()
	rabbitCh, err := rabbitClient.Channel()
	if err != nil {
		log.Fatalf("Failed to open channel, %v", err)
	}
	defer rabbitCh.Close()
	defer rabbitClient.Close()

	r := website.Initial(redisClient, psDB, rabbitCh)

	host := viper.GetString("server.host")
	port := viper.GetString("server.port")
	if port == "" {
		port = "8000"
	}
	r.Run(host + ":" + port)
	fmt.Println("Start Listning...")
}
