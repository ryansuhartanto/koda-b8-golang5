package menu

import "bufio"

type Login struct {
}

func (this Login) Name() string {
	return "Login"
}

func (this Login) Handle(scanner *bufio.Scanner) *Menu {
	return nil
}
