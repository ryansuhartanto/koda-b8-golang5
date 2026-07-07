package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ryansuhartanto/koda-b8-golang5/data"
	"github.com/ryansuhartanto/koda-b8-golang5/menu"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	database := data.Database{}

	system := menu.System{}
	var menu menu.Menu

	for true {
		fmt.Print("\033c")

		if menu == nil {
			menu = system
		}

		menu.Handle(scanner, database)
	}
}
