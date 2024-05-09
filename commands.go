package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/XyaelD/pokedexcli/internal/pokeapi"
	"github.com/XyaelD/pokedexcli/internal/pokecache"
)

type config struct {
	pokeCache     pokecache.Cache
	pokeapiClient pokeapi.Client
	Next          *string
	Previous      *string
}

type cliCommand struct {
	name     string
	desc     string
	callback func(*config) error
}

func helpCommand(config *config) error {
	fmt.Println()
	fmt.Printf("Welcome to the Pokedex!\n\nUsage:\n\n")
	userCommands := userCommands()
	for _, userCommand := range userCommands {
		fmt.Printf("%v: %v\n", userCommand.name, userCommand.desc)
	}
	fmt.Println()
	return nil
}

func exitCommand(config *config) error {
	os.Exit(1)
	return nil
}

func mapCommand(config *config) error {

	locationResults := pokeapi.ShallowLocations{}
	var err error

	exists, ok := config.pokeCache.Get(*config.Next)
	if ok {
		// For testing cache
		// fmt.Println("Cache accessed!")
		err = json.Unmarshal(exists, &locationResults)
	} else {
		locationResults, err = config.pokeapiClient.GetLocations(*config.Next, &config.pokeCache)
	}

	if err != nil {
		return nil
	}

	config.Next = locationResults.Next
	config.Previous = locationResults.Previous

	for _, location := range locationResults.Results {
		fmt.Printf("%v\n", location.Name)
	}
	return nil
}

func mapbCommand(config *config) error {
	if config.Previous == nil {
		return errors.New("cannot go back from the first page")
	}

	locationResults := pokeapi.ShallowLocations{}
	var err error

	exists, ok := config.pokeCache.Get(*config.Previous)
	if ok {
		// For testing cache
		// fmt.Println("Cache accessed!")
		err = json.Unmarshal(exists, &locationResults)
	} else {
		locationResults, err = config.pokeapiClient.GetLocations(*config.Previous, &config.pokeCache)
	}

	if err != nil {
		return nil
	}

	config.Next = locationResults.Next
	config.Previous = locationResults.Previous

	for _, location := range locationResults.Results {
		fmt.Printf("%v\n", location.Name)
	}
	return nil
}

func userCommands() map[string]cliCommand {

	userCommands := map[string]cliCommand{
		"help": {
			name:     "help",
			desc:     "Displays a help message",
			callback: helpCommand,
		},
		"exit": {
			name:     "exit",
			desc:     "Exits the Pokedex",
			callback: exitCommand,
		},
		"map": {
			name:     "map",
			desc:     "Shows map locations",
			callback: mapCommand,
		},
		"mapb": {
			name:     "mapb",
			desc:     "Shows previous map locations",
			callback: mapbCommand,
		},
	}
	return userCommands
}
