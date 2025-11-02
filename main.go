package main

import (
	"fmt"

	"github.com/EthanColbert8/gator-project/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("Error reading config:", err)
		return
	}

	cfg.SetUser("Ethan")

	cfg, err = config.Read()
	if err != nil {
		fmt.Println("Error reading config:", err)
		return
	}

	fmt.Println("Current Username:", cfg.CurrentUsername)
	fmt.Println("Database URL:", cfg.DbUrl)
}
