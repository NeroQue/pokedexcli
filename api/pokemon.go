package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/NeroQue/pokedexcli/models"
)

func FetchAndDisplay(cfg *models.Config, endpoint string) error {
	cachedData, ok := cfg.Cache.Get(endpoint)
	if ok {
		fmt.Println("(Using cached data)")
		var response models.LocationAreasResponse
		if err := json.Unmarshal(cachedData, &response); err == nil {
			printLocations(response)
			updateConfig(cfg, response)
			return nil
		}
	}

	res, err := http.Get(endpoint)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		fmt.Printf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return err
	}

	cfg.Cache.Add(endpoint, body)

	var response models.LocationAreasResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	printLocations(response)
	updateConfig(cfg, response)
	return nil
}

func printLocations(resp models.LocationAreasResponse) {
	for _, loc := range resp.Results {
		fmt.Println(loc.Name)
	}
}

func updateConfig(cfg *models.Config, resp models.LocationAreasResponse) {
	cfg.Previous = resp.Previous
	cfg.Next = resp.Next
}
