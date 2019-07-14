package main

import (
	"github.com/katoozi/gin-web-site/cmd/gin-web-site/commands"
	_ "github.com/katoozi/gin-web-site/cmd/gin-web-site/commands"
	_ "github.com/lib/pq"

	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command

func init() {
	rootCmd = &cobra.Command{
		Use:   "main.go [command]",
		Short: "gin web site main section cli",
	}
	rootCmd.AddCommand(commands.RunServerCommand)
	rootCmd.AddCommand(commands.CreateSuperUserCommand)
}

func main() {
	rootCmd.Execute()
}
