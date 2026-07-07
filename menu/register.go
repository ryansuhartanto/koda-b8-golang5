package menu

import (
	"bufio"
	"fmt"

	"github.com/ryansuhartanto/koda-b8-golang5/data"
)

type Register struct {
}

func (this Register) Name() string {
	return "Register"
}

func (this Register) Handle(scanner *bufio.Scanner, database data.Database) *Menu {
	fmt.Println("--- Register ---")
	fmt.Println()

	return nil
}
