package main

import pokecache "github.com/ohrelaxo/pokedexcli/internal"

type config struct {
	next     *string
	previous *string
	cache    *pokecache.Cache
}

type cliCommand struct {
	name        string
	description string
	callback    func(configptr *config) error
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
	}
	return commands
}
