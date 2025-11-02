package main

import (
	"fmt"
	"os"

	"github.com/EthanColbert8/gator-project/internal/config"
	"github.com/EthanColbert8/gator-project/internal/utils"
)

func main() {
	var state utils.State
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("Error reading config:", err)
		os.Exit(1)
	}
	state.Cfg = &cfg

	commands := &utils.Commands{}
	commands.Register("login", utils.HandlerLogin)

	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("No command provided")
		os.Exit(1)
	}

	cmd := utils.Command{
		Name: args[0],
		Args: args[1:],
	}

	err = commands.Run(&state, cmd)
	if err != nil {
		fmt.Println("Error executing command:", err)
		os.Exit(1)
	}
}
