package utils

import (
	"debris/types"
	"fmt"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

func GetConfigPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	path := filepath.Join(homeDir, ".config", "debris", "config.toml")
	return path
}

func LoadConfig() (types.Config, error) {
	configPath := GetConfigPath()

	data, err := os.ReadFile(configPath)
	if err != nil {
		return types.Config{}, fmt.Errorf("Config file not found!")
	}

	var config types.Config
	err = toml.Unmarshal(data, &config)
	if err != nil {
		return types.Config{}, fmt.Errorf("Failed to parsing TOML: %w", err)
	}

	return config, nil
}
