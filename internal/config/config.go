package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

type Config struct {
	ApiToken string `json:"api_token"`
}

func Get() (Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return Config{}, fmt.Errorf("failed to get user home directory %w", err)
	}

	data, err := os.ReadFile(path.Join(home, ".config", "todo", "config.json"))
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil

}
