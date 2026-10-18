package main

import (
	"bufio"
	"fmt"
	"github.com/timjdavis93/bootdev/pokedexcli/internal/pokeapi"
	//	"github.com/timjdavis93/bootdev/pokedexcli/internal/pokecache"
	//	"math/rand"
	"os"
	"strings"
	"time"
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

	apiClient := pokeapi.NewClient(httpTimeout, cacheInterval)
	baseURL := "https://pokeapi.co/api/v2/"
	nextURL := "https://pokeapi.co/api/v2/location-area/?limit=20&offset=0"
	prevURL := ""

	//bufio.Scanner <- os.Stdin
	conf := &config{
		commands: cliCommands,
		baseURL:  &baseURL,
		nextURL:  &nextURL,
		prevURL:  &prevURL,
	}

	// TODO: cliCommands file, read into program?
	cliCommands := getCommands()

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
		if err := command.callback(conf, apiClient, tmp[1:]); err != nil {
			fmt.Println("Error: ", err)
		}

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, *pokeapi.Client, []string) error
}

type config struct {
	commands map[string]cliCommand
	baseURL  *string
	nextURL  *string
	prevURL  *string
}

func cleanInput(text string) []string {
	var s []string
	text = strings.ToLower(text)
	s = strings.Fields(text)
	return s
}
