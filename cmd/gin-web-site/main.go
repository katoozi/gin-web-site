package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/katoozi/gin-web-site/configs"
	"github.com/katoozi/gin-web-site/internal/app/website"
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

func main() {
	// fetch server and database configs from config.yaml file
	serverConfig := fetchServerConfig()
	databaseConfig := fetchDatabaseConfig()

	fmt.Println("database config: ", databaseConfig)

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
