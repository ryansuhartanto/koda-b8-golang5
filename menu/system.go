package menu

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/ryansuhartanto/koda-b8-golang5/data"
)

type System struct{}

var (
	loggedInOptions  = []Menu{}
	loggedOutOptions = []Menu{
		Register{},
		Login{},
		ForgotPassword{},
	}
)

func (this System) GetOptions(loggedIn bool) []Menu {
	if loggedIn {
		return loggedInOptions
	} else {
		return loggedOutOptions
	}
}

func (this System) Name() string {
	return "System"
}

func (this System) Handle(scanner *bufio.Scanner, database *data.Database) Menu {
	fmt.Println("--- Welcome to system ---")
	fmt.Println()

	defer HandlePanic(scanner, "Wrong selection, press enter to back...")

	current := database.Current
	options := this.GetOptions(current != nil)

	if current != nil {
		fmt.Printf("Hello, %v!\n", current.First)
	}
	for index, option := range options {
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

	return options[selected-1]
}
