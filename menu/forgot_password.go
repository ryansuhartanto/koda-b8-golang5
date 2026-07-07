package menu

import (
	"bufio"

	"github.com/ryansuhartanto/koda-b8-golang5/data"
)

type ForgotPassword struct {
}

func (this ForgotPassword) Name() string {
	return "Forgot Password"
}

func (this ForgotPassword) Handle(scanner *bufio.Scanner, database data.Database) Menu {
	return nil
}
