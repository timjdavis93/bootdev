package main

import (
	"errors"
	//	"fmt"
	// "github.com/timjdavis93/bootdev/gator/internal/config"
)

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	registeredCommands map[string]func(*State, Command) error
}

// Runs a given command with the provided state if it exists
func (c *Commands) run(s *State, cmd Command) error {
	f, ok := c.registeredCommands[cmd.Name]
	if !ok {
		return errors.New("command not found")
	}
	return f(s, cmd)
}

// registers a new handler function for a command name
func (c *Commands) register(name string, f func(*State, Command) error) {
	c.registeredCommands[name] = f
}
