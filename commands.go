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

type exploreAPI struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func commandExit(_ *config, _ string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(_ *config, _ string) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	supportedCommands := getCommands()
	for _, v := range supportedCommands {
		fmt.Printf("%v: %v\n", v.name, v.description)
	}
	return nil
}

func commandMap(configptr *config, _ string) error {
	if configptr.next == nil {
		return fmt.Errorf("error no next field in the configptr")
	}
	url := *configptr.next
	err := mapCommandsHelper(url, configptr)
	if err != nil {
		return err
	}
	return nil
}

func commandMapb(configptr *config, _ string) error {
	if configptr.previous == nil {
		return fmt.Errorf("error no pervious field in the configptr")
	}
	url := *configptr.previous
	err := mapCommandsHelper(url, configptr)
	if err != nil {
		return err
	}
	return nil
}

func mapCommandsHelper(url string, configptr *config) error {
	rawjson, ok := configptr.cache.Get(url)

	if !ok {
		result, err := fetchPokeAPI(url, configptr)
		if err != nil {
			return err
		}
		rawjson = result
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

func commandExplore(configptr *config, areaName string) error {
	baseUrl := "https://pokeapi.co/api/v2/location-area/"
	url := baseUrl + areaName
	rawjson, ok := configptr.cache.Get(url)

	if !ok {
		result, err := fetchPokeAPI(url, configptr)
		if err != nil {
			return err
		}
		rawjson = result
	}
	var result exploreAPI
	if err := json.Unmarshal(rawjson, &result); err != nil {
		return fmt.Errorf("error while unmarshal the read response boy: %w", err)
	}

	for _, encounters := range result.PokemonEncounters {
		fmt.Println(encounters.Pokemon.Name)
	}
	return nil
}

func fetchPokeAPI(url string, configptr *config) ([]byte, error) {
	var result []byte

	res, err := http.Get(url)
	if err != nil {
		return result, fmt.Errorf("error while making a http request to the pokeAPI: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)
	result, err = io.ReadAll(res.Body)
	if err != nil {
		return result, fmt.Errorf("error while reading the response body: %w", err)
	}
	configptr.cache.Add(url, result)
	return result, nil
}
