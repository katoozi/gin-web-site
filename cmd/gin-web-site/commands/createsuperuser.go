package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

// CreateSuperUserCommand will create user with superuser permission in user table
var CreateSuperUserCommand = &cobra.Command{
	Use:   "createsuperuser",
	Short: "create user with superuser permissions",
	Run: func(cmd *cobra.Command, args []string) {
		createSuperuser()
	},
}

// CreateSuperuser will create user with superuser permission in user table
func createSuperuser() {
	fmt.Println("CreateSuperUser invoked...")
}
