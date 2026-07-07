package menu

import (
	"bufio"
	"fmt"
)

type Menu interface {
	Name() string
	Handle(scanner *bufio.Scanner) *Menu
}

func HandlePanic(scanner *bufio.Scanner, message string) {
	if r := recover(); r != nil {
		fmt.Print(message, " ")
		scanner.Scan()
	}
}
