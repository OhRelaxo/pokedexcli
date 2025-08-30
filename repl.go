package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ohrelaxo/pokedexcli/internal"
)

func repl() {
	url := "https://pokeapi.co/api/v2/location-area"
	cache, err := pokecache.NewCache(5 * time.Second)
	if err != nil {
		log.Fatal(err)
	}

	configptr := config{next: &url, previous: nil, cache: &cache}

	userInput := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		userInput.Scan() // gibt bool zurück, wenn der Scanner kein Token mehr hat. Scanner.Err() gibt einen error dann zurück siehe: https://pkg.go.dev/bufio#Scanner.Scan
		token := userInput.Text()
		if len(token) == 0 {
			continue
		}
		words := cleanInput(token)
		searchForCommand(words, &configptr)
	}
}

func searchForCommand(words []string, configptr *config) {
	var areaName string
	for i, word := range words {
		switch word {
		case "help":
			err := commandHelp(configptr, areaName)
			if err != nil {
				fmt.Printf("error while executing the help command: %v\n", err)
			}
		case "exit":
			err := commandExit(configptr, areaName)
			if err != nil {
				fmt.Printf("error while executing the exit command: %v\n", err)
				fmt.Println("exiting the programm")
				os.Exit(1)
			}
		case "map":
			err := commandMap(configptr, areaName)
			if err != nil {
				fmt.Printf("error while executing the map command: %v\n", err)
			}
		case "mapb":
			err := commandMapb(configptr, areaName)
			if err != nil {
				fmt.Printf("error while executing the mapb command: %v\n", err)
			}
		case "explore":
			areaName = words[i+1]
			err := commandExplore(configptr, areaName)
			if err != nil {
				fmt.Printf("error while executing the explore command: %v\n", err)
			}
		}
	}
}
