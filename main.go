package main

import (
	"time"

	"github.com/XyaelD/pokedexcli/internal/pokeapi"
	"github.com/XyaelD/pokedexcli/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	pokeCache := pokecache.NewCache(600 * time.Second)
	pokeDex := pokeapi.NewPokedex()

	startingValue := "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"

	config := &config{
		pokeDex:       pokeDex,
		pokeCache:     pokeCache,
		pokeapiClient: pokeClient,
		Next:          &startingValue,
		Previous:      nil,
	}
	runPokedex(config)
}
