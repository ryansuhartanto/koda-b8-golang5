package menu

import (
	"bufio"
	"fmt"
	"os"
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

	defer HandlePanic(scanner, "Wrong selection, press enter to back...")

	for index, option := range systemOptions {
		fmt.Printf("%v. %v\n", index+1, option.Name())
	}
	fmt.Println()

	fmt.Printf("%v. %v\n", 0, "Exit")
	fmt.Println()

	fmt.Print("Choose a menu: ")
	scanner.Scan()
	selected, err := strconv.Atoi(scanner.Text())

	if err != nil {
		panic(err)
	}

	if selected == 0 {
		os.Exit(0)
	}

	return &systemOptions[selected-1]
}
