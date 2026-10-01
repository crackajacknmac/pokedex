package main

import (

"errors"
"fmt"

)

func commandExplore(fig *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("unable to return location information")
		}
	areaName := args[0]
	resp, err := fig.pokeapiClient.Lookup(areaName)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", areaName)
	fmt.Println("Found Pokemon:")
	for _, encounter := range resp.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}
