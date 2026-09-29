package main

import (
"errors"
"fmt"
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
