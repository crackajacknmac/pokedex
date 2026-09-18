package main

import (
"github.com/crackajacknmac/pokedex/internal/pokeapi"
)

func commandMap(fig *config) error {
        resp, err := pokeapi.GetLocations(config.pageNext)
	if err != nil {
		return error.New("incorrect location url")
	}
	config.pageBack = resp
	config.pageNext = 
	//come back to this statement for your notes
	return nil
}
