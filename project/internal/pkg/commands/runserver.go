package commands

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	// initial rabbitMQ
	rabbitClient := initialRabbitmq()
	rabbitCh, err := rabbitClient.Channel()
	if err != nil {
		log.Fatalf("Failed to open channel, %v", err)
	}

	// initial website package
	r := website.Initial(redisClient, psDB, rabbitCh)

	host := viper.GetString("server.host")
	port := viper.GetString("server.port")
	if port == "" {
		port = "8000"
	}
	r.Addr = host + ":" + port

	// create channel for shutdown server gracefully
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt, os.Kill)
	go func() {
		r.ListenAndServe()
	}()
	log.Println("Start Listning...")
	sig := <-sigs
	log.Println("Signal: ", sig)

	log.Println("Stopping Gin Server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := r.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Closing Postgres Connection...")
	psDB.Close()

	log.Println("Closing Redis Connection...")
	redisClient.Close()

	log.Println("Closing RabbitMQ Connection...")
	rabbitCh.Close()
	rabbitClient.Close()
}
