package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type pokeAPI struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
	} `json:"results"`
}

func commandExit(_ *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(_ *config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	supportedCommands := getCommands()
	for _, v := range supportedCommands {
		fmt.Printf("%v: %v\n", v.name, v.description)
	}
	return nil
}

func commandMap(configptr *config) error {
	if configptr.next == nil {
		return fmt.Errorf("error no next field in the configptr")
	}
	url := *configptr.next
	err := mapCommandshelper(url, configptr)
	if err != nil {
		return err
	}
	return nil
}

func commandMapb(configptr *config) error {
	if configptr.previous == nil {
		return fmt.Errorf("error no pervious field in the configptr")
	}
	url := *configptr.previous
	err := mapCommandshelper(url, configptr)
	if err != nil {
		return err
	}
	return nil
}

func mapCommandshelper(url string, configptr *config) error {
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error while making a http request to the pokeAPI: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)
	rawjson, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error while reading the response body: %w", err)
	}
	var result pokeAPI
	if err := json.Unmarshal(rawjson, &result); err != nil {
		return fmt.Errorf("error while unmarshal the read response boy: %w", err)
	}
	configptr.next = result.Next
	configptr.previous = result.Previous
	for _, areas := range result.Results {
		fmt.Println(areas.Name)
	}
	return nil
}
