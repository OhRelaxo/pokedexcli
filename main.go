package main

import (
	"strings"
)

func main() {
	repl()
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
