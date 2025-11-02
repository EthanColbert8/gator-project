package utils

import "fmt"

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("got %d args, expected 1 for command login", len(cmd.Args))
	}

	err := s.Cfg.SetUser(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("error setting user: %w", err)
	}

	fmt.Printf("Current user set to: %s\n", cmd.Args[0])
	return nil
}
