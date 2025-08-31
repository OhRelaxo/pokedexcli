package main

import pokecache "github.com/ohrelaxo/pokedexcli/internal"

type config struct {
	next     *string
	previous *string
	cache    *pokecache.Cache
	pokedex  map[string]catchAPI
}

type cliCommand struct {
	name        string
	description string
	callback    func(configptr *config, arg string) error
}

func getCommands() map[string]cliCommand {
	commands := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays 20 locations from the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations from the Pokemon world",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "list all the Pokemon located in an area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "catches a Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "see details about your Pokemon's",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "displays your caught Pokemon's",
			callback:    commandPokedex,
		},
	}
	return commands
}
