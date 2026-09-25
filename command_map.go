package main

import (
"errors"
"fmt"
"github.com/crackajacknmac/pokedex/internal/pokeapi"
)

func commandMap(fig *config) error {
        resp, err := fig.pokeapiClient.GetLocations(fig.pageNext)
	if err != nil {
		return errors.New("incorrect location url")
	}
	fig.pageBack = resp.Previous
	fig.pageNext = resp.Next
	for _, location := range resp.Results {
		fmt.Printf("%s\n", location.Name)
	}
	return nil
}
