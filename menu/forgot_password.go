package menu

import "bufio"

type ForgotPassword struct {
}

func (this ForgotPassword) Name() string {
	return "Forgot Password"
}

func (this ForgotPassword) Handle(scanner *bufio.Scanner) *Menu {
	return nil
}
