package pokeapi

import "errors"

type Pokedex struct {
	Pokedex map[string]Pokemon
}

func NewPokedex() Pokedex {

	NewPokedex := Pokedex{
		Pokedex: make(map[string]Pokemon),
	}
	return NewPokedex
}

func (p *Pokedex) Add(pokemon string, data Pokemon) {
	p.Pokedex[pokemon] = data
}

func (p *Pokedex) Lookup(pokemon string) (Pokemon, error) {
	_, ok := p.Pokedex[pokemon]
	if !ok {
		return Pokemon{}, errors.New("you have not caught that pokemon")
	}
	return p.Pokedex[pokemon], nil
}
