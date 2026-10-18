package main

import (
	"fmt"
	"github.com/timjdavis93/bootdev/gator/internal/config"
	"log"
	"os"
)

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
	conf, err := config.ReadConfig()
	if err != nil {
		return
	}

	programState := State{
		Config: conf,
	}

	commands := Commands{
		registeredCommands: make(map[string]func(*State, Command) error),
	}

	commands.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatal("usage: cli <command> [args...]")
	}

	cmd := Command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	if err := commands.run(&programState, cmd); err != nil {
		fmt.Println("Failed to run given command")
		fmt.Println(err)
		os.Exit(1)
	}

}

type State struct {
	Config *config.Config
}

/*
	// NOTE: REPL

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print( "Gator >" )
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
*/

/*
func cleanInput(text string) []string {
	var s []string
	text = strings.ToLower(text)
	s = strings.Fields(text)
	return s
} */
