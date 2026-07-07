package menu

import (
	"bufio"
	"fmt"

	"github.com/ryansuhartanto/koda-b8-golang5/data"
)

type Logout struct {
}

func (this Logout) Name() string {
	return "Logout"
}

func (this Logout) Handle(scanner *bufio.Scanner, database *data.Database) (menu Menu) {
	database.Logout()

	fmt.Print("Logged-out successfully, press enter to go back...")
	scanner.Scan()

	return nil
}
