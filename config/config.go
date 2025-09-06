package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config holds the application configuration.
type Config struct {
	InputFile string `json:"inputFile"`
	OutputDir string `json:"outputDir"`
	LogLevel  int    `json:"logLevel"`
}

// LoadConfig reads configuration from the given file path.
func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	var cfg Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, fmt.Errorf("failed to decode config JSON: %w", err)
	}

	return &cfg, nil
}