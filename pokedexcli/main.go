package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var cliCommands map[string]cliCommand

func main() {
	//bufio.Scanner <- os.Stdin
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
		command.callback()

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func cleanInput(text string) []string {
	var s []string
	text = strings.ToLower(text)
	s = strings.Fields(text)
	return s
}

func commandExit() error {

	fmt.Println("Closing the Pokedex... Goodbye!\n")
	defer os.Exit(0)
	return fmt.Errorf("Closing the Pokedex... Goodbye!\n")
}

func commandHelp() error {

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
