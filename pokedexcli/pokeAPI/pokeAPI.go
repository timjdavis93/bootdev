package pokeAPI

import (
	//	"fmt"
	"encoding/json"
	"net/http"
)

func getMap(url string) (Location, error) {
	//default returns ?limit=20 can also ?offset=## to get a different page
	//url := "https://pokeapi.co/api/v2/location/" // https://pokeapi.co/api/v2/location/{id or name}/

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer res.Body.Close()

	var loc location
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&loc)
	if err != nil {
		return Location{}, err
	}

	return loc, nil
}
