package commands

import (
	"fmt"
	"os"

	"github.com/NeroQue/pokedexcli/api"
	"github.com/NeroQue/pokedexcli/models"
)

var Commands map[string]models.Command

func init() {
	Commands = map[string]models.Command{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    CommandExit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    CommandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Shows the locations of the Pokemon world",
			Callback:    CommandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Shows the previous locations of the Pokemon world",
			Callback:    CommandMapb,
		},
		"explore": {
			Name:        "explore",
			Description: "Shows the available pokemon on the given map",
			Callback:    CommandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "Catches the Pokemon <Pokemon Name>",
			Callback:    CommandCatch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Shows you details about your pokemon",
			Callback:    CommandInspect,
		},
	}
}

func CommandExit(cfg *models.Config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return fmt.Errorf("Exited")
}

func CommandHelp(cfg *models.Config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	for _, command := range Commands {
		fmt.Printf("%s: %s\n", command.Name, command.Description)
	}
	return nil
}

func CommandMap(cfg *models.Config, args ...string) error {
	endpoint := ""
	if cfg.Next != "" {
		endpoint = cfg.Next
	} else {
		endpoint = "https://pokeapi.co/api/v2/location-area/"
	}
	return api.FetchAndDisplay(cfg, endpoint)
}

func CommandMapb(cfg *models.Config, args ...string) error {
	endpoint := ""
	if cfg.Previous != "" {
		endpoint = cfg.Previous
	} else {
		endpoint = "https://pokeapi.co/api/v2/location-area/"
	}
	return api.FetchAndDisplay(cfg, endpoint)
}

func CommandExplore(cfg *models.Config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: explore <location-area>")
	}
	area := args[0]
	endpoint := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", area)
	return api.FetchAndExplore(cfg, endpoint)
}

func CommandCatch(cfg *models.Config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: catch <location-area>")
	}
	pokemon := args[0]
	endpoint := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", pokemon)
	fmt.Printf("Throwing a Pokeball at %s...", pokemon)
	return api.FetchAndCatch(cfg, endpoint, pokemon)
}

func CommandInspect(cfg *models.Config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: inspect <pokemon-name>")
	}
	name := args[0]
	pokemon, ok := cfg.CaughtPokemon[name]
	if ok {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)

		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}

		fmt.Println("Types:")
		for _, typeInfo := range pokemon.Types {
			fmt.Printf("  - %s\n", typeInfo.Type.Name)
		}
	} else {
		fmt.Printf("You haven't caught %s yet!\n", name)
	}
	return nil
}
