package menu

import (
	"bufio"
	"fmt"
)

type Menu interface {
	Name() string
	Handle(scanner *bufio.Scanner) *Menu
}

func HandlePanic(scanner *bufio.Scanner) {
	if r := recover(); r != nil {
		fmt.Print(r, " ")
		scanner.Scan()
	}
}
