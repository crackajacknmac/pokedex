package pokeapi

import (
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
