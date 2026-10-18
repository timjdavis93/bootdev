package main

import (
	"fmt"
	"github.com/timjdavis93/bootdev/gator/internal/config"
	"os"
)

func main() {
	conf, err := config.Read()
	if err != nil {
		return
	}

	state := State{
		Config: conf,
	}

	var cmds map[string]func(*State, Command) error
	commands := Commands{
		cmds: cmds,
	}

	commands.register("login", handlerLogin)

	args := os.Args
	if len(args) < 2 {
		fmt.Println("Error Args")
		return
	}

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

func cleanInput(text string) []string {
	var s []string
	text = strings.ToLower(text)
	s = strings.Fields(text)
	return s
}
