package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/NeroQue/pokedexcli/cache"
	"github.com/NeroQue/pokedexcli/cli"
	"github.com/NeroQue/pokedexcli/commands"
	"github.com/NeroQue/pokedexcli/models"
)

func main() {
	cfg := &models.Config{
		Cache:         cache.NewCache(10 * time.Second),
		CaughtPokemon: make(map[string]models.PokemonResponse),
	}
	r := io.Reader(os.Stdin)
	scanner := cli.NewScanner(r)

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		userInput := scanner.Text()

		words := cli.CleanInput(userInput)

		if len(words) == 0 {
			continue
		}

		command := words[0]

		commandExists := false
		for cmdName, cmd := range commands.Commands {
			if command == cmdName {
				err := cmd.Callback(cfg, words[1:]...)
				if err != nil {
					fmt.Println("Error:", err)
				}
				commandExists = true
				break
			}
		}

		if !commandExists {
			fmt.Println("Unknown command")
		}
	}
}
