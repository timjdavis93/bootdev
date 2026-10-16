package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/timjdavis93/bootdev/pokedexCLI/pokeAPI"
	//"github.com/timjdavis93/bootdev/pokedexCLI/pokecache"
)

var cliCommands map[string]cliCommand

// Define ANSI color codes
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	White  = "\033[37m"
)

func main() {
	const cacheInterval = 5 * time.Minute
	const httpTimeout = 5 & time.Second

	apiClient := pokeAPI.NewClient(httpTimeout, cacheInterval)

	//bufio.Scanner <- os.Stdin
	conf := &config{
		commands: cliCommands,
		baseURL:  "https://pokeAPI.co/api/v2/location-area/",
		nextURL:  "https://pokeAPI.co/api/v2/location-area/?limit=20&offset=0",
		prevURL:  "",
	}

	cliCommands = map[string]cliCommand{
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

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(Red + "Pokedex >" + Reset)
		scanner.Scan()
		text := scanner.Text()
		tmp := cleanInput(text)
		if len(tmp) == 0 {
			continue
		}
		command, ok := cliCommands[tmp[0]]
		if !ok {
			fmt.Println("Not a command!")
			continue
		}
		command.callback(conf, apiClient, tmp[1:])

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, *pokeAPI.Client, []string) error
}

type config struct {
	commands map[string]cliCommand
	baseURL  string
	nextURL  string
	prevURL  string
}

func cleanInput(text string) []string {
	var s []string
	text = strings.ToLower(text)
	s = strings.Fields(text)
	return s
}

func commandExit(conf *config, apiClient *pokeAPI.Client, args []string) error {

	fmt.Println(Red + "Closing the Pokedex... Goodbye!" + Reset)
	defer os.Exit(0)
	return nil
}

func commandHelp(conf *config, apiClient *pokeAPI.Client, args []string) error {

	if cliCommands == nil {
		return fmt.Errorf(Red + "No commands!\n" + Reset)
	}

	if len(cliCommands) == 0 {
		return fmt.Errorf("No commands!\n")
	}

	fmt.Println(White + "Welcome to the Pokedex!")
	fmt.Println("Usage:" + Reset)
	fmt.Println("")

	for _, commands := range cliCommands {
		fmt.Printf("%s: %s\n", commands.name, commands.description)
	}

	return nil
}

func commandMap(conf *config, apiClient *pokeAPI.Client, args []string) error {
	//url := "https://pokeAPI.co/api/v2/location-area/?limit=20"
	url := conf.nextURL

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

func commandMapb(conf *config, apiClient *pokeAPI.Client, args []string) error {
	//url := "https://pokeAPI.co/api/v2/location-area/?limit=20"
	url := conf.prevURL

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

func commandExplore(conf *config, apiClient *pokeAPI.Client, args []string) error {
	//url := "https://pokeAPI.co/api/v2/location-area/" + "area name"

	if len(args) == 0 {
		fmt.Println("No location to explore")
		return nil
	}
	areaName := args[0]
	url := conf.baseURL + areaName

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

func commandCatch(conf *config, apiClient *pokeAPI.Client, args []string) error {

	if len(args) == 0 {
		fmt.Println("No Pokemon named")
	}

	pokemonName := args[0]
	url := conf.baseURL + pokemonName

	pokemon, err := apiClient.GetPokemon(url)
	if err != nil {
		fmt.Println("Error GetPokemon API")
		return err
	}

	fmt.Println(pokemon.Name)

	return nil
}
