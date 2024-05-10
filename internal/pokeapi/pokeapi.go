package pokeapi

import (
	"encoding/json"
	"io"

	"github.com/XyaelD/pokedexcli/internal/pokecache"
)

func (c *Client) GetLocations(locationUrl string, cache *pokecache.Cache) (ShallowLocations, error) {
	res, err := c.httpClient.Get(locationUrl)
	if err != nil {
		return ShallowLocations{}, err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if err != nil {
		return ShallowLocations{}, err
	}

	locationResults := ShallowLocations{}
	err = json.Unmarshal(body, &locationResults)
	if err != nil {
		return ShallowLocations{}, err
	}
	cache.Add(locationUrl, body)
	return locationResults, nil
}

func (c *Client) GetPokemonByLocation(cityName string, cache *pokecache.Cache) (PokemonByLocation, error) {
	searchURL := "https://pokeapi.co/api/v2/location-area/" + cityName
	res, err := c.httpClient.Get(searchURL)
	if err != nil {
		return PokemonByLocation{}, err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if err != nil {
		return PokemonByLocation{}, err
	}

	pokemonResults := PokemonByLocation{}
	err = json.Unmarshal(body, &pokemonResults)
	if err != nil {
		return PokemonByLocation{}, err
	}
	cache.Add(searchURL, body)
	return pokemonResults, nil
}

func (c *Client) CatchPokemon(pokemon string, cache *pokecache.Cache) (Pokemon, error) {
	searchURL := "https://pokeapi.co/api/v2/pokemon/" + pokemon
	res, err := c.httpClient.Get(searchURL)
	if err != nil {
		return Pokemon{}, err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if err != nil {
		return Pokemon{}, err
	}

	pokemonResults := Pokemon{}
	err = json.Unmarshal(body, &pokemonResults)
	if err != nil {
		return Pokemon{}, err
	}
	cache.Add(searchURL, body)
	return pokemonResults, nil
}
