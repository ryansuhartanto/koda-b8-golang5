package menu

import (
	"bufio"
	"fmt"
	"strconv"
)

type System struct {
	LoggedIn bool
}

var systemOptions = []Menu{
	Register{},
	Login{},
	ForgotPassword{},
}

func (this System) Name() string {
	return "System"
}

func (this System) Handle(scanner *bufio.Scanner) *Menu {
	fmt.Println("--- Welcome to system ---")
	fmt.Println()

	defer HandlePanic(scanner)

	for index, option := range systemOptions {
		fmt.Printf("%v. %v\n", index+1, option.Name())
	}
	fmt.Println()

	fmt.Printf("%v. %v\n", 0, "Exit")
	fmt.Println()

	fmt.Print("Choose a menu: ")
	scanner.Scan()
	option, err := strconv.Atoi(scanner.Text())

	if err != nil {
		panic("Wrong selection, press enter to back... ")
	}

	option -= 1

	return nil
}
