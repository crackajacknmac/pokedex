package pokeapi

import (
"fmt"
"errors"
"io"
"net/http"
"encoding/json"
)

type LocationAreaResponse struct {
	Count int		`json:"count"`
	Next *string		`json:"next"`
	Previous *string	`json:"previous"`
	Results []struct {
		Name string	`json:"name"`
		Url string	`json:"url"`
	} `json:"results"`
}


type LocationLookup struct {
    PokemonEncounters []struct {
        Pokemon struct {
            Name string `json:"name"`
            URL  string `json:"url"`
        } `json:"pokemon"`
    } `json:"pokemon_encounters"`
}


func (c Client) GetLocations(url *string) (resp LocationAreaResponse, err error) {
	var endpoint string
	if url == nil {
		endpoint = "https://pokeapi.co/api/v2/location-area"
		} else {
			endpoint = *url
		}
	res, err := http.Get(endpoint)
	if err != nil {
		return LocationAreaResponse{}, errors.New("no valid url")
		}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		return LocationAreaResponse{}, errors.New("Response failed")
		}

	locations := LocationAreaResponse{}
	err = json.Unmarshal(body, &locations)
		if err != nil {
			return LocationAreaResponse{}, errors.New("Response failed")
		}

	return locations, nil
}


func (c *Client) Lookup(location string) (LocationLookup, error) {
	checkLocation := "https://pokeapi.co/api/v2/location-area/" + location
	if cachedBytes, found := c.cache.Get(checkLocation); found {
		beenThere := LocationLookup{}
		err := json.Unmarshal(cachedBytes, &beenThere)
                if err != nil {
                        return LocationLookup{}, errors.New("unable to fetch previous data")
               }
		return beenThere, nil
	}

        res, err := http.Get(checkLocation)
        if err != nil {
                return LocationLookup{}, errors.New("no valid url")
                }
	defer res.Body.Close()
        body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationLookup{}, errors.New("unable to read data")
		}

        if res.StatusCode > 299 {
                return LocationLookup{}, fmt.Errorf("bad status code: %v", res.StatusCode)
                }

	c.cache.Add(checkLocation, body)

        feedback := LocationLookup{}
        err = json.Unmarshal(body, &feedback)
                if err != nil {
                        return LocationLookup{}, errors.New("unable to fetch new data")
                }

        return feedback, nil
}

