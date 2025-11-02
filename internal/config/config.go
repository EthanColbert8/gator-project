package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUsername string `json:"current_user_name"`
}

func Read() (Config, error) {
	configPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, fmt.Errorf("error getting config file path: %w", err)
	}

	config, err := os.ReadFile(configPath)
	if err != nil {
		return Config{}, fmt.Errorf("error reading config file: %w", err)
	}

	var cfg Config
	err = json.Unmarshal(config, &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("error unmarshaling config file: %w", err)
	}

	return cfg, nil
}

func (c Config) SetUser(username string) error {
	c.CurrentUsername = username

	jsonConfig, err := json.Marshal(c)
	if err != nil {
		return fmt.Errorf("error marshaling configuration: %w", err)
	}

	configPath, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("error getting config file path: %w", err)
	}

	err = os.WriteFile(configPath, jsonConfig, 0644)
	if err != nil {
		return fmt.Errorf("error writing configuration file: %w", err)
	}

	return nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error finding home directory: %w", err)
	}

	return path.Join(homeDir, configFileName), nil
}
