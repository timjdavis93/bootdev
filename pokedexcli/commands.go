package main

import (
	"fmt"
	"github.com/timjdavis93/bootdev/pokedexcli/internal/pokeapi"
	"os"
)

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"map": {
			name:        "map",
			description: "Iterate through the locations, 20 each time",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "map... the other way... b wards",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "explore <city name> to look at an area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "catch things",
			callback:    commandCatch,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func commandMap(conf *config, apiClient *pokeapi.Client, args []string) error {
	//url := "https://pokeapi.co/api/v2/location-area/?limit=20"
	url := *conf.nextURL

	loc, err := apiClient.GetMap(url)
	if err != nil {
		fmt.Println("Error GetMap API")
		return err
	}

	for _, locations := range loc.Results {
		fmt.Println(locations.Name)
	}

	conf.nextURL = loc.Next
	conf.prevURL = loc.Previous
	return nil
}

func commandMapb(conf *config, apiClient *pokeapi.Client, args []string) error {
	//url := "https://pokeapi.co/api/v2/location-area/?limit=20"
	url := *conf.prevURL

	loc, err := apiClient.GetMap(url)
	if err != nil {
		fmt.Println("Error GetMap API")
		return err
	}
	for _, locations := range loc.Results {
		fmt.Println(locations.Name)
	}

	conf.nextURL = loc.Next
	conf.prevURL = loc.Previous
	return nil
}

func commandExplore(conf *config, apiClient *pokeapi.Client, args []string) error {
	//url := "https://pokeapi.co/api/v2/location-area/" + "area name"

	if len(args) == 0 {
		fmt.Println("No location to explore")
		return nil
	}
	areaName := args[0]
	url := *conf.baseURL + "location-area/" + areaName

	area, err := apiClient.GetArea(url)
	if err != nil {
		fmt.Println("Error GetMap API")
		return err
	}
	// TODO: Need to cache the found pokemon so that I can tag them as "found" before allowing catch
	fmt.Println(White + "Exploring: " + areaName)
	fmt.Println("Found Pokemon: ")
	for _, encounters := range area.PokemonEncounters {
		fmt.Println(" - " + encounters.Pokemon.Name)
	}
	fmt.Println(Reset)

	return nil
}

func commandCatch(conf *config, apiClient *pokeapi.Client, args []string) error {

	if len(args) == 0 {
		fmt.Println("No Pokemon named")
	}

	pokemonName := args[0]
	url := *conf.baseURL + "pokemon/" + pokemonName

	pokemon, err := apiClient.GetPokemon(url)
	if err != nil {
		fmt.Println("Error GetPokemon API")
		return err
	}
	fmt.Println(White, "Throwing a Pokeball at "+pokemon.Name+"..."+Reset)

	// TODO: I really don't want to come up with a catch percentage chance right now, should come back refactor and add this when I've got more time and experience.
	fmt.Println("Caught", pokemon.Name, "successfully!")

	return nil
}

func commandHelp(conf *config, apiClient *pokeapi.Client, args []string) error {

	fmt.Println(White + "Welcome to the Pokedex!")
	fmt.Println("Usage:" + Reset)
	fmt.Println("")
	cliCommands := getCommands()

	for _, commands := range cliCommands {
		fmt.Printf("%s: %s\n", commands.name, commands.description)
	}

	return nil
}

func commandExit(conf *config, apiClient *pokeapi.Client, args []string) error {

	fmt.Println(Red + "Closing the Pokedex... Goodbye!" + Reset)
	defer os.Exit(0)
	return nil
}
