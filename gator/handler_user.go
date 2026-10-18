package main

import (
	"fmt"
)

func handlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Usage: %s <name>", cmd.Name)
	}

	name := cmd.Args[0]

	if err := s.Config.SetUser(name); err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}
