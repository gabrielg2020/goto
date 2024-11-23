package pkg

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Config struct
type Config struct {
	MaxDepth    int      `yaml:"maxDepth"`
	ExcludeDirs []string `yaml:"excludeDirs"`
}

// LoadConfig loads the configuration from $HOME/.goto_config.yaml
func LoadConfig() (*Config, error) {
	// Get the user home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("error getting the user home directory: %v", err)
	}

	// Open the configuration file
	configPath := filepath.Join(homeDir, ".goto_config.yaml")
	file, err := os.Open(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			// Default configuration
			return &Config{
				MaxDepth:    5,
				ExcludeDirs: []string{".git", "node_modules"},
			}, nil
		}
		return nil, fmt.Errorf("error opening the configuration file: %v", err)
	}
	defer file.Close()

	// Decode the configuration
	var config Config
	if err := yaml.NewDecoder(file).Decode(&config); err != nil {
		return nil, fmt.Errorf("error decoding the configuration file: %v", err)
	}

	return &config, nil
}
