package main

import (
	"bufio"
	"fmt"
	"os"
)

func repl() {
	userInput := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		userInput.Scan() // gibt bool zurück, wenn der Scanner kein Token mehr hat. Scanner.Err() gibt einen error dann zurück siehe: https://pkg.go.dev/bufio#Scanner.Scan
		token := userInput.Text()
		if len(token) == 0 {
			continue
		}
		words := cleanInput(token)
		searchForCommand(words)
	}
}

func searchForCommand(words []string) {
	for _, word := range words {
		switch word {
		case "help":
			err := commandHelp()
			if err != nil {
				fmt.Printf("error with executing the help command: %v\n", err)
			}
		case "exit":
			err := commandExit()
			if err != nil {
				fmt.Printf("error with executing the exit command: %v\n", err)
				fmt.Println("exiting the programm")
				os.Exit(1)
			}
		}
	}
}
