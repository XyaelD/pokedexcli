package pokeapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/XyaelD/pokedexcli/internal/pokecache"
)

func TestGetLocations(t *testing.T) {
	nextURL := "http://example.com/next"
	previousURL := "http://example.com/prev"

	expectedResult := ShallowLocations{
		Count:    3,
		Next:     &nextURL,
		Previous: &previousURL,
		Results: []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		}{
			{
				Name: "Location 1",
				URL:  "http://example.com/location1",
			},
			{
				Name: "Location 2",
				URL:  "http://example.com/location2",
			},
		},
	}

	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Convert the expected results to []byte and return those when the server is hit
		resp, err := json.Marshal(expectedResult)
		if err != nil {
			t.Fatalf("Unable to marshal mock response: %v", err)
		}
		w.Write(resp)
	}))

	defer server.Close()

	c := &Client{
		httpClient: *server.Client(),
	}

	cache := pokecache.NewCache(10 * time.Second)

	actualResult, err := c.GetLocations(server.URL, &cache)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !reflect.DeepEqual(expectedResult, actualResult) {
		t.Errorf("Expected result %v, got %v", expectedResult, actualResult)
	}
}

// Need to alter code to mock the server properly
// func TestGetPokemonByLocation(t *testing.T) {

// 	expectedResult := PokemonByLocation{
// 		PokemonEncounters: []struct {
// 			Pokemon struct {
// 				Name string `json:"name"`
// 			} `json:"pokemon"`
// 		}{
// 			{
// 				Pokemon: struct {
// 					Name string `json:"name"`
// 				}{
// 					Name: "pikachu",
// 				},
// 			},
// 			{
// 				Pokemon: struct {
// 					Name string `json:"name"`
// 				}{
// 					Name: "charmander",
// 				},
// 			},
// 			{
// 				Pokemon: struct {
// 					Name string `json:"name"`
// 				}{
// 					Name: "combee",
// 				},
// 			},
// 		},
// 	}

// 	// Start a local HTTP server
// 	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// Convert the expected results to []byte and return those when the server is hit
// 		resp, err := json.Marshal(expectedResult)
// 		if err != nil {
// 			t.Fatalf("Unable to marshal mock response: %v", err)
// 		}
// 		w.Write(resp)
// 	}))

// 	defer server.Close()

// 	c := &Client{
// 		httpClient: *server.Client(),
// 	}

// 	cache := pokecache.NewCache(10 * time.Second)

// 	actualResult, err := c.GetPokemonByLocation(server.URL, &cache)
// 	if err != nil {
// 		t.Fatalf("Expected no error, got %v", err)
// 	}

// 	if !reflect.DeepEqual(expectedResult, actualResult) {
// 		t.Errorf("Expected result %v, got %v", expectedResult, actualResult)
// 	}
// }
