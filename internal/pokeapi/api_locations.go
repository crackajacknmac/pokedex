package pokeapi

import (
"errors"
"fmt"
"io"
"log"
"net/http"
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



func GetLocations(url string) (resp LocationAreaResponse, err error) {
	res, err := http.Get(url)
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
