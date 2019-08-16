package main

import (
	"github.com/katoozi/gin-web-site/internal/pkg/commands"

	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command

func init() {
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
