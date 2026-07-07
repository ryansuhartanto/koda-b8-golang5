package menu

import (
	"bufio"
	"fmt"

	"github.com/ryansuhartanto/koda-b8-golang5/data"
)

type ForgotPassword struct {
}

func (this ForgotPassword) Name() string {
	return "Forgot Password"
}

func (this ForgotPassword) Handle(scanner *bufio.Scanner, database *data.Database) (menu Menu) {
	fmt.Println("--- Register ---")
	fmt.Println()

	menu = this

	fmt.Print("Enter your email: ")
	scanner.Scan()
	email := scanner.Text()

	user := database.ForgotPassword(email, true)

	var password string
	for {
		if func() bool {
			defer HandlePanic(scanner, "Failed to confirm password, press enter to retry...")

			fmt.Print("Enter a strong password: ")
			scanner.Scan()
			password = scanner.Text()

			fmt.Print("Confirm your password: ")
			scanner.Scan()

			if password != scanner.Text() {
				panic("")
			}

			return true
		}() {
			break
		}
	}

	user.ChangePassword(password)

	fmt.Print("Password successfully changed, press enter to go back...")
	scanner.Scan()

	return nil
}
