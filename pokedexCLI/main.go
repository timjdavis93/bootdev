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

func main() {
	const cacheInterval = 5 * time.Minute
	const httpTimeout = 5 & time.Second

	apiClient := pokeAPI.NewClient(httpTimeout, cacheInterval)

	//bufio.Scanner <- os.Stdin
	conf := &config{
		commands: cliCommands,
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
		fmt.Print("Pokedex >")
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
		command.callback(conf, apiClient)

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, *pokeAPI.Client) error
}

type config struct {
	commands map[string]cliCommand
	nextURL  string
	prevURL  string
}

func cleanInput(text string) []string {
	var s []string
	text = strings.ToLower(text)
	s = strings.Fields(text)
	return s
}

func commandExit(conf *config, apiClient *pokeAPI.Client) error {

	fmt.Println("Closing the Pokedex... Goodbye!")
	defer os.Exit(0)
	return nil
}

func commandHelp(conf *config, apiClient *pokeAPI.Client) error {

	if cliCommands == nil {
		return fmt.Errorf("No commands!\n")
	}

	if len(cliCommands) == 0 {
		return fmt.Errorf("No commands!\n")
	}

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	for _, commands := range cliCommands {
		fmt.Printf("%s: %s\n", commands.name, commands.description)
	}

	return nil
}

func commandMap(conf *config, apiClient *pokeAPI.Client) error {
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

func commandMapb(conf *config, apiClient *pokeAPI.Client) error {
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
