package pokeAPI

import (
	"encoding/json"
	"fmt"
	"io"
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

func (c *Client) GetArea(url string) (Area, error) {

	var loc Area
	data, ok := c.cache.Get(url)
	if ok {
		if err := json.Unmarshal(data, &loc); err != nil {
			return Area{}, err
		}
		return loc, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Area{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Area{}, err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&loc)
	if err != nil {
		return Area{}, err
	}

	return loc, nil
}

func (c *Client) GetPokemon(url string) (Pokemon, error) {

	var pokemon Pokemon
	data, err := c.fetchAndCache(url)
	if err != nil {
		fmt.Println("Failure to fetch pokemon")
		return Pokemon{}, err
	}
	if err := json.Unmarshal(data, &pokemon); err != nil {
		fmt.Println("Failure to Unmarshal pokemon")
		return Pokemon{}, err
	}

	fmt.Println("Sending Pokemon to main")
	return pokemon, nil

}

func (c *Client) fetchAndCache(url string) ([]byte, error) {

	if data, ok := c.cache.Get(url); ok {
		fmt.Println("Pokemon-data found in cache")
		return data, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Http.NewRequest failed")
		return []byte{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Println("c.httpClient.Do(req) failed")
		return []byte{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("io.ReadAll() failed")
		return []byte{}, err
	}

	c.cache.Add(url, data)

	fmt.Println("data found online")
	return data, nil
}
