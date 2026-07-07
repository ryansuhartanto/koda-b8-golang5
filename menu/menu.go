package menu

import (
	"bufio"
	"fmt"

	"github.com/ryansuhartanto/koda-b8-golang5/data"
)

type Menu interface {
	Name() string
	Handle(scanner *bufio.Scanner, database data.Database) Menu
}

func HandlePanic(scanner *bufio.Scanner, message string) {
	if r := recover(); r != nil {
		fmt.Print(message, " ")
		scanner.Scan()
	}
}
