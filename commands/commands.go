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
	}
}

func CommandExit(cfg *models.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return fmt.Errorf("Exited")
}

func CommandHelp(cfg *models.Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	for _, command := range Commands {
		fmt.Printf("%s: %s\n", command.Name, command.Description)
	}
	return nil
}

func CommandMap(cfg *models.Config) error {
	endpoint := ""
	if cfg.Next != "" {
		endpoint = cfg.Next
	} else {
		endpoint = "https://pokeapi.co/api/v2/location-area/"
	}
	return api.FetchAndDisplay(cfg, endpoint)
}

func CommandMapb(cfg *models.Config) error {
	endpoint := ""
	if cfg.Previous != "" {
		endpoint = cfg.Previous
	} else {
		endpoint = "https://pokeapi.co/api/v2/location-area/"
	}
	return api.FetchAndDisplay(cfg, endpoint)
}
