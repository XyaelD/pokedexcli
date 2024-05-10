package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"

	"github.com/XyaelD/pokedexcli/internal/pokeapi"
	"github.com/XyaelD/pokedexcli/internal/pokecache"
)

type config struct {
	pokeDex       pokeapi.Pokedex
	pokeCache     pokecache.Cache
	pokeapiClient pokeapi.Client
	Next          *string
	Previous      *string
}

type cliCommand struct {
	name     string
	desc     string
	callback func(*config, string) error
}

func helpCommand(config *config, _ string) error {
	fmt.Println()
	fmt.Printf("Welcome to the Pokedex!\n\nUsage:\n\n")
	userCommands := userCommands()
	for _, userCommand := range userCommands {
		fmt.Printf("%v: %v\n", userCommand.name, userCommand.desc)
	}
	fmt.Println()
	return nil
}

func exitCommand(config *config, _ string) error {
	os.Exit(1)
	return nil
}

func mapCommand(config *config, _ string) error {

	locationResults := pokeapi.ShallowLocations{}
	var err error

	exists, ok := config.pokeCache.Get(*config.Next)
	if ok {
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

func mapbCommand(config *config, _ string) error {
	if config.Previous == nil {
		return errors.New("cannot go back from the first page")
	}

	locationResults := pokeapi.ShallowLocations{}
	var err error

	exists, ok := config.pokeCache.Get(*config.Previous)
	if ok {
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

func exploreCommand(config *config, location string) error {

	pokemonResults := pokeapi.PokemonByLocation{}
	var err error

	searchURL := "https://pokeapi.co/api/v2/location-area/" + location

	exists, ok := config.pokeCache.Get(searchURL)
	if ok {
		err = json.Unmarshal(exists, &pokemonResults)
	} else {
		pokemonResults, err = config.pokeapiClient.GetPokemonByLocation(location, &config.pokeCache)
	}

	if err != nil {
		return nil
	}

	for _, explore := range pokemonResults.PokemonEncounters {
		fmt.Printf("%v\n", explore.Pokemon.Name)
	}
	return nil
}

func catchCommand(config *config, pokemon string) error {

	pokemonResults := pokeapi.Pokemon{}
	var err error

	searchURL := "https://pokeapi.co/api/v2/pokemon/" + pokemon

	exists, ok := config.pokeCache.Get(searchURL)
	if ok {
		err = json.Unmarshal(exists, &pokemonResults)
	} else {
		pokemonResults, err = config.pokeapiClient.CatchPokemon(pokemon, &config.pokeCache)
	}

	if err != nil {
		return nil
	}

	catchValue := rand.Intn(pokemonResults.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %v\n", pokemonResults.Name)

	if catchValue <= 50 {
		fmt.Printf("%v was caught!\n", pokemonResults.Name)
		fmt.Println("You may now inspect it with the inspect command.")
		config.pokeDex.Add(pokemonResults.Name, pokemonResults)
		// Testing
		// fmt.Printf("This is the weight from the Pokedex: %v\n", config.pokeDex.Pokedex[pokemonResults.Name].Weight)
	} else {
		fmt.Printf("%v escaped!\n", pokemonResults.Name)
	}

	// Testing
	// for name := range config.pokeDex.Pokedex {
	// 	fmt.Printf("%v is registered in the Pokedex\n", name)
	// }

	return nil
}

func inspectCommand(config *config, pokemon string) error {

	pokemonInfo, err := config.pokeDex.Lookup(pokemon)
	if err != nil {
		return err
	}

	fmt.Printf("Name: %v\n", pokemonInfo.Name)
	fmt.Printf("Height: %v\n", pokemonInfo.Height)
	fmt.Printf("Weight: %v\n", pokemonInfo.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemonInfo.Stats {
		fmt.Printf(" -%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokeType := range pokemonInfo.Types {
		fmt.Printf(" -%v\n", pokeType.Type.Name)
	}

	return nil
}

func pokedexCommand(config *config, _ string) error {

	if len(config.pokeDex.Pokedex) == 0 {
		fmt.Println("You have not caught any Pokemon yet!")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for name := range config.pokeDex.Pokedex {
		fmt.Printf(" - %v\n", name)
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
		"explore": {
			name:     "explore",
			desc:     "Shows all Pokemon encounters for a location",
			callback: exploreCommand,
		},
		"catch": {
			name:     "catch",
			desc:     "Try to catch a Pokemon",
			callback: catchCommand,
		},
		"inspect": {
			name:     "inspect",
			desc:     "Inspect a Pokemon in your Pokedex",
			callback: inspectCommand,
		},
		"pokedex": {
			name:     "pokedex",
			desc:     "Lists all entries in your Pokedex",
			callback: pokedexCommand,
		},
	}
	return userCommands
}
