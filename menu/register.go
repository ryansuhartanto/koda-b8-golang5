package menu

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/ryansuhartanto/koda-b8-golang5/data"
)

type Register struct {
}

func (this Register) Name() string {
	return "Register"
}

func (this Register) Handle(scanner *bufio.Scanner, database *data.Database) (menu Menu) {
	fmt.Println("--- Register ---")
	fmt.Println()

	menu = this

	fmt.Print("What is your first name: ")
	scanner.Scan()
	first := scanner.Text()

	fmt.Print("What is your last name: ")
	scanner.Scan()
	last := scanner.Text()

	fmt.Print("What is your email: ")
	scanner.Scan()
	email := scanner.Text()

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

	fmt.Println()
	fmt.Println("You are about to register as...")
	fmt.Printf("First name: %v\n", first)
	fmt.Printf("Last name: %v\n", last)
	fmt.Printf("Email: %v\n", email)
	fmt.Println()

	fmt.Print("Is this correct [y/N]: ")
	scanner.Scan()
	if strings.ToLower(scanner.Text()) != "y" {
		return
	}

	database.Register(first, last, email, password)

	fmt.Print("Registered successfully, press enter to go back...")
	scanner.Scan()

	return nil
}
