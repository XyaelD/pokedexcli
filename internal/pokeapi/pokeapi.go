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
	// For testing cache
	// fmt.Printf("Adding this URL: %v", locationUrl)
	cache.Add(locationUrl, body)
	return locationResults, nil
}
