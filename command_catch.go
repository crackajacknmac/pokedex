package main

import (
	"fmt"
	"errors"
	"math/rand"
)

func commandCatch(fig *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("unable to return pokemon information")
		}
	pokename := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokename)
	resp, err := fig.pokeapiClient.GetPokemonInfo(pokename)
	if err != nil {
		return err
	}

	difficulty := resp.BaseExperience
	roll := rand.Intn(320)
	if roll >= difficulty {
		fmt.Printf("%v was caught!\n", pokename)
		fig.pokedex[pokename] = resp
	} else {
		fmt.Printf("%v escaped!\n", pokename)
	}

	return nil
}
