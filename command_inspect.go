package main

import (
	"errors"
	"fmt"
)

func commandInspect(fig *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("no pokemon provided")
	}

	request := args[0]
	pokemon, found := fig.pokedex[request]
	if found {
		
		fmt.Printf("Name: %v\n", pokemon.Name)
		fmt.Printf("Height: %v\n", pokemon.Height)
		fmt.Printf("Weight: %v\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  -%v: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, Type := range pokemon.ElementType {
			fmt.Printf("  - %v\n", Type.TypeInfo.Name)
		}
	}else{
		fmt.Println("you have not caught that pokemon")
	}
	return nil
}
