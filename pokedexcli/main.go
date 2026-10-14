package main

import (
	"bufio"
	"fmt"
	//	"net/http"
	"github.com/timjdavis93/bootdev/pokedexCLI/pokeAPI"
	"os"
	"strings"
)

var cliCommands map[string]cliCommand

func main() {
	//bufio.Scanner <- os.Stdin
	conf := &config{
		commands: cliCommands,
		baseUrl:  "https://pokeapi.co",
	}

	scanner := bufio.NewScanner(os.Stdin)
	cliCommands = map[string]cliCommand{
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
		command.callback(conf)

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	commands map[string]cliCommand
	baseUrl  string
}

func cleanInput(text string) []string {
	var s []string
	text = strings.ToLower(text)
	s = strings.Fields(text)
	return s
}

func commandExit(conf *config) error {

	fmt.Println("Closing the Pokedex... Goodbye!\n")
	defer os.Exit(0)
	return nil
}

func commandHelp(conf *config) error {

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
