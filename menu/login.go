package menu

import (
	"bufio"
	"fmt"

	"github.com/ryansuhartanto/koda-b8-golang5/data"
)

type Login struct {
}

func (this Login) Name() string {
	return "Login"
}

func (this Login) Handle(scanner *bufio.Scanner, database *data.Database) (menu Menu) {
	fmt.Println("--- Login ---")
	fmt.Println()

	defer HandlePanic(scanner, "Wrong email or password, press enter to retry...")
	menu = this

	fmt.Print("Enter your email: ")
	scanner.Scan()
	email := scanner.Text()

	fmt.Print("Enter your password: ")
	scanner.Scan()
	password := scanner.Text()

	database.Login(email, password)

	fmt.Print("Logged-in successfully, press enter to go back...")
	scanner.Scan()

	return nil
}
