package api

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand/v2"
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

func updateConfig(cfg *models.Config, m interface{}) {
	switch model := m.(type) {
	case models.LocationAreasResponse:
		cfg.Previous = model.Previous
		cfg.Next = model.Next
	default:
		cfg.Previous = ""
		cfg.Next = ""
	}
}

func FetchAndExplore(cfg *models.Config, endpoint string) error {
	cachedData, ok := cfg.Cache.Get(endpoint)
	if ok {
		fmt.Println("(Using cached data)")
		var response models.EncountersResponse
		if err := json.Unmarshal(cachedData, &response); err == nil {
			printEncounters(response)
		}
		return nil
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
	var response models.EncountersResponse
	if err := json.Unmarshal(body, &response); err == nil {
		printEncounters(response)
	}
	updateConfig(cfg, response)
	return nil
}

func printEncounters(resp models.EncountersResponse) {
	for _, i := range resp.PokemonEncounters {
		fmt.Println(i.Pokemon.Name)
	}
}

func FetchAndCatch(cfg *models.Config, endpoint string, pokemonName string) error {
	cachedData, ok := cfg.Cache.Get(endpoint)
	if ok {
		fmt.Println("(Using cached data)")
		var response models.PokemonResponse
		if err := json.Unmarshal(cachedData, &response); err == nil {
			return handleCatchAttempt(cfg, &response, pokemonName)
		}
	}
	res, err := http.Get(endpoint)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		fmt.Printf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return err
	}
	cfg.Cache.Add(endpoint, body)
	var response models.PokemonResponse
	if err := json.Unmarshal(body, &response); err == nil {
		return handleCatchAttempt(cfg, &response, pokemonName)
	}
	return nil
}

func handleCatchAttempt(cfg *models.Config, pokemon *models.PokemonResponse, pokemonName string) error {
	fmt.Println()

	catchRate := 1.0 - (float64(pokemon.BaseExperience) / 1275.0)
	if catchRate < 0.1 {
		catchRate = 0.1
	}

	caught := rand.Float64() < catchRate

	if caught {
		cfg.CaughtPokemon[pokemonName] = *pokemon
		fmt.Printf("%s was caught!\n", pokemonName)
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}
