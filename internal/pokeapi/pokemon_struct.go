package pokeapi

import (
	"net/http"
	"encoding/json"
	"fmt"
	"errors"
	"io"
)

type PokemonInfo struct {
	Name string 			`json:"name"`
	BaseExperience int		`json:"base_experience"`
	ElementType []PokeType		`json:"types"`
	Stats []PokeStats		`json:"stats"`
	Height int			`json:"height"`
	Weight int			`json:"weight"`
}

type PokeStats struct {
	BaseStat int			`json:"base_stat"`
	Effort int			`json:"effort"`
	Stat struct {
		Name string		`json:"name"`
		Url string		`json:"url"`
	} `json:"stat"`
}

type PokeType struct {
	Slot int			`json:"slot"`
	TypeInfo struct{
		Name string		`json:"name"`
		Url string		`json:"url"`
	} `json:"type"`
}


func (c *Client) GetPokemonInfo(name string) (PokemonInfo, error) {
	checkPokemon := "https://pokeapi.co/api/v2/pokemon/" + name
	if cachedBytes, found := c.cache.Get(checkPokemon); found {
		seenIt := PokemonInfo{}
		err := json.Unmarshal(cachedBytes, &seenIt)
		if err != nil {
			return PokemonInfo{}, errors.New("unable to fetch previous data")
		}
	return seenIt, nil
	}

	res, err := http.Get(checkPokemon)
	if err != nil {
		return PokemonInfo{}, errors.New("unable to retrieve data")
		}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonInfo{}, errors.New("unable to read data")
		}
	if res.StatusCode > 299 {
		return PokemonInfo{}, fmt.Errorf("bad status code: %v", res.StatusCode)
		}
	c.cache.Add(checkPokemon, body)

	pokemon := PokemonInfo{}
	err = json.Unmarshal(body, &pokemon)
		if err != nil {
			return PokemonInfo{}, errors.New("Unable to retrieve data")
			}
	return pokemon, nil
}
