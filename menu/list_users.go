package menu

import (
	"bufio"
	"encoding/hex"
	"fmt"

	"github.com/ryansuhartanto/koda-b8-golang5/data"
)

type ListAllUsers struct {
}

func (this ListAllUsers) Name() string {
	return "List All Users"
}

func (this ListAllUsers) Handle(scanner *bufio.Scanner, database *data.Database) (menu Menu) {
	fmt.Println("--- List All Users ---")
	fmt.Println()

	for index, user := range database.Users {
		fmt.Printf("%v.\n", index+1)
		fmt.Printf("Full Name: %v %v\n", user.First, user.Last)
		fmt.Printf("Email: %v\n", user.Email)
		fmt.Printf("Password: %v\n", hex.EncodeToString(user.Password[:]))
	}

	fmt.Println()
	fmt.Print("Press enter to go back...")
	scanner.Scan()

	return nil
}
