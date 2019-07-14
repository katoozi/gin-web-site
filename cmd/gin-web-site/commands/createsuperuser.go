package commands

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/katoozi/gin-web-site/internal/pkg/models"
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
	color.Green("Start Creating Superuser...")
	var (
		username    string
		pass        string
		passConfirm string
	)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Username: ")
		username, _ = reader.ReadString('\n')
		if checkUsername(username) {
			break
		} else {
			color.Red("User with this username already exists")
		}
	}
	fmt.Print("Email: ")
	email, _ := reader.ReadString('\n')

	for {
		fmt.Print("Password: ")
		pass, _ = reader.ReadString('\n')
		fmt.Print("Confirm Password: ")
		passConfirm, _ = reader.ReadString('\n')
		if pass == passConfirm {
			break
		} else {
			color.Red("Passwords Does not Match!!!")
		}
	}

	user := models.NewUser("", "", username, email, pass)
	user.IsSperuser = true
	user.IsStaff = true

	_, err := dbCon.Exec(user.GenerateInsertQuery())
	if err != nil {
		color.Red("Error while create super user: %v", err)
	}

	color.Green("Superuser Seccessfully Created.")

}

func checkUsername(username string) bool {
	err := models.GetUser(username, dbCon)
	if err == nil {
		return true
	}
	return false
}
