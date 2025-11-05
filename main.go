package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/EthanColbert8/gator-project/internal/config"
	"github.com/EthanColbert8/gator-project/internal/database"
	"github.com/EthanColbert8/gator-project/internal/utils"
	_ "github.com/lib/pq"
)

func main() {
	var state utils.State

	cfg, err := config.Read()
	if err != nil {
		fmt.Println("Error reading config:", err)
		os.Exit(1)
	}
	state.Cfg = &cfg

	db, err := sql.Open("postgres", state.Cfg.DbUrl)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		os.Exit(1)
	}
	dbQueries := database.New(db)
	state.Db = dbQueries

	commands := &utils.Commands{}
	commands.Register("login", utils.HandlerLogin)
	commands.Register("register", utils.HandlerRegister)
	commands.Register("users", utils.HandlerGetUsers)
	commands.Register("reset", utils.HandlerResetUsers)

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
