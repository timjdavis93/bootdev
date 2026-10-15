package pokeAPI

import (
	"net/http"
	"time"
)

import "github.com/timjdavis93/bootdev/pokedexCLI/pokecache"

type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

func NewClient(httpTimeout, cacheInterval time.Duration) *Client {
	return &Client{
		httpClient: http.Client{Timeout: httpTimeout},
		cache:      pokecache.NewCache(cacheInterval),
	}
}

type Location struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
