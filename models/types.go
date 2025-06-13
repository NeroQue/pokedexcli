package models

import "github.com/NeroQue/pokedexcli/cache"

type Config struct {
	Next     string
	Previous string
	Cache    *cache.Cache
}

type LocationAreasResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type Command struct {
	Name        string
	Description string
	Callback    func(*Config) error
}
