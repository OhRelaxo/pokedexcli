package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	userInput := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		userInput.Scan() // gibt bool zurück, wenn der Scanner kein Token mehr hat. Scanner.Err() gibt einen error dann zurück siehe: https://pkg.go.dev/bufio#Scanner.Scan
		token := userInput.Text()
		if len(token) == 0 {
			continue
		}
		words := cleanInput(token)
		fmt.Printf("Your command was: %v\n", words[0])
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
