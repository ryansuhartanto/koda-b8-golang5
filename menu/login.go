package menu

import (
	"bufio"

	"github.com/ryansuhartanto/koda-b8-golang5/data"
)

type Login struct {
}

func (this Login) Name() string {
	return "Login"
}

func (this Login) Handle(scanner *bufio.Scanner, database data.Database) *Menu {
	return nil
}
