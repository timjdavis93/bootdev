package main

import (
	"fmt"
	"github.com/timjdavis93/bootdev/gator/internal/config"
)

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	cmds map[string]func(*State, Command) error
}

type State struct {
	Config *config.Config
}

func handlerLogin(s *State, cmd Command) error {

	if len(cmd.Args) == 0 {
		return fmt.Errorf("Args Error")
	}

	s.Config.Current_user_name = cmd.Args[0]
	fmt.Println("Username Set")

	return nil
}

// Runs a given command with the provided state if it exists
func (c *Commands) run(s *State, cmd Command) error {
	return nil
}

// registers a new handler function for a command name
func (c *Commands) register(name string, f func(*State, Command) error) error {
	return nil
}
