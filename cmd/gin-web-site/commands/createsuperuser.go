package commands

import (
	"bufio"
	"fmt"
	"os"
	"syscall"

	"github.com/fatih/color"
	"github.com/katoozi/gin-web-site/internal/pkg/models"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
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
		bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
		pass = string(bytePassword)

		fmt.Print("\nConfirm Password: ")
		bytePassword, _ = terminal.ReadPassword(int(syscall.Stdin))
		passConfirm = string(bytePassword)

		if pass == passConfirm {
			break
		} else {
			color.Red("\nPasswords Does not Match!!!\n")
		}
	}

	user := models.NewUser("", "", username, email, pass)
	user.IsSperuser = true
	user.IsStaff = true

	_, err := dbCon.Exec(user.GenerateInsertQuery())
	if err != nil {
		color.Red("\nError while create super user: %v\n", err)
	}

	color.Green("\nSuperuser Seccessfully Created.")

}

func checkUsername(username string) bool {
	err := models.GetUser(username, dbCon)
	if err == nil {
		return true
	}
	return false
}
