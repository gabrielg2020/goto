package pkg

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

// Config struct
type Config struct {
	MaxDepth    int      `yaml:"maxDepth"`
	ExcludeDirs []string `yaml:"excludeDirs"`
}

// LoadConfig loads the configuration from $HOME/.goto_config.yaml
func LoadConfig(configPath string) (*Config, error) {
	// Open the configuration file
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
