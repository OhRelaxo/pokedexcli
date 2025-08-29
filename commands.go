package main

import (
	"fmt"
	"os"
)

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	supportedCommands := getCommands()
	for _, v := range supportedCommands {
		fmt.Printf("%v: %v\n", v.name, v.description)
	}
	return nil
}
