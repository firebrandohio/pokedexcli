package main

import "fmt"

func commandHelp() {
	availableCommands := getCommands()
	fmt.Println("Welcome to the pokedex help menu!")
	fmt.Println("Available commands:")
	fmt.Println("")

	for _, command := range availableCommands {
		fmt.Printf("  %s: %s\n", command.name, command.description)
	}
	fmt.Println("")
}
