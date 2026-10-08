package main

import (
	"fmt"
	"errors"
)

func commandPokedex(fig *config, _ ...string) error {
	if len(fig.pokedex) == 0 {
		return errors.New("no pokemon caught yet")
	}
	fmt.Println("Your Pokedex:")
	for _, pokemon := range fig.pokedex {
		fmt.Printf(" - %v\n", pokemon.Name)
	}
	return nil
}
