package menu

import "bufio"

type Register struct {
}

func (this Register) Name() string {
	return "Register"
}

func (this Register) Handle(scanner *bufio.Scanner) *Menu {
	return nil
}
