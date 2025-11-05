package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/EthanColbert8/gator-project/internal/database"
	"github.com/google/uuid"
)

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("got %d args, expected 1 for command login", len(cmd.Args))
	}

	_, err := s.Db.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("error retrieving user: %w", err)
	}

	err = s.Cfg.SetUser(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("error setting user: %w", err)
	}

	fmt.Printf("Current user set to: %s\n", cmd.Args[0])
	return nil
}

func HandlerRegister(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("got %d args, expected 1 for command register", len(cmd.Args))
	}

	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
	}

	newUser, err := s.Db.CreateUser(context.Background(), userParams)
	if err != nil {
		return fmt.Errorf("error registering user: %w", err)
	}

	s.Cfg.SetUser(newUser.Name)
	fmt.Printf("User %s registered and set as current user.\n", newUser.Name)

	fmt.Printf("\nUser Details:\nID: %s\nCreated At: %s\nUpdated At: %s\n\n", newUser.ID, newUser.CreatedAt, newUser.UpdatedAt)
	return nil
}

func HandlerResetUsers(s *State, cmd Command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("got %d args, expected 0 for command reset", len(cmd.Args))
	}

	err := s.Db.ResetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error resetting users table: %w", err)
	}

	return nil
}
