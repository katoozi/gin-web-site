package main

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/katoozi/gin-web-site/internal/pkg/commands"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd *cobra.Command

func init() {
	// customize logging style
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// config the viper
	// only read from environmet variables
	viper.SetConfigName("config")
	viper.SetEnvPrefix("test_project")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	viper.ReadInConfig()

	// set the gin mode
	mode := viper.Get("mode")
	if mode == nil {
		mode = "debug"
	}
	gin.SetMode(mode.(string))

	// configure the cobra commands
	rootCmd = &cobra.Command{
		Use:   "main [command]",
		Short: "gin web site main section cli",
	}
	rootCmd.AddCommand(commands.RunServerCommand)
	rootCmd.AddCommand(commands.CreateSuperUserCommand)
}

func main() {
	rootCmd.Execute()
}
