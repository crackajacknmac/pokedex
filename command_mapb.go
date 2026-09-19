package main

import (
"fmt"
"errors"
"github.com/crackajacknmac/pokedex/internal/pokeapi"
)

func commandMapb(fig *config) error {
	if fig.pageBack == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	resp, err := pokeapi.GetLocations(fig.pageBack)
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
