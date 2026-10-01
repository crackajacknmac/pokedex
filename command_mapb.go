package main

import (
"fmt"
"errors"
)

func commandMapb(fig *config, _ ...string) error {
	if fig.pageBack == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	resp, err := fig.pokeapiClient.GetLocations(fig.pageBack)
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
