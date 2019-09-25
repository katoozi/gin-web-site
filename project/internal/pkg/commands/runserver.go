package commands

import (
	"fmt"

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

	// connect to redis
	redisClient := initialRedis()

	r := website.Initial(redisClient, psDB)

	host := viper.GetString("server.host")
	port := viper.GetString("server.port")
	if port == "" {
		port = "8000"
	}
	r.Run(host + ":" + port)
	fmt.Println("Start Listning...")
}
