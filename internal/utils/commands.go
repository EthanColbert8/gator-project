package utils

import (
	"fmt"
)

type Commands struct {
	handlers map[string]func(*State, Command) error
}

type Command struct {
	Name string
	Args []string
}

func (c *Commands) Run(s *State, cmd Command) error {
	handler, exists := c.handlers[cmd.Name]
	if !exists {
		return fmt.Errorf("unknown command: %s", cmd.Name)
	}

	return handler(s, cmd)
}

func (c *Commands) Register(name string, handler func(*State, Command) error) {
	if c.handlers == nil {
		c.handlers = make(map[string]func(*State, Command) error)
	}

	c.handlers[name] = handler
}
