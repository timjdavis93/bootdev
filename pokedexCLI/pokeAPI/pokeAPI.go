package pokeAPI

import (
	//	"fmt"
	"encoding/json"
	"net/http"
)

func (c *Client) GetMap(url string) (Location, error) {

	var loc Location
	data, ok := c.cache.Get(url)
	if ok {
		if err := json.Unmarshal(data, &loc); err != nil {
			return Location{}, err
		}
		return loc, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&loc)
	if err != nil {
		return Location{}, err
	}

	return loc, nil
}
