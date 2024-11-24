package test

import (
	"os"
	"testing"

	gotoPkg "github.com/gabrielg2020/goto/cmd/pkg"
)

// Helper function to create a temporary configuration file
func createTempConfigFile(content string) string {
	// Create a temporary file
	file, err := os.CreateTemp("", "goto_config_*.yaml")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Write the content to the file
	if _, err := file.WriteString(content); err != nil {
		panic(err)
	}

	return file.Name()
}

func TestLoadConfig(t *testing.T) {
	mockConfig := `maxDepth: 3
excludeDirs:
  - node_modules
  - .git
  - .vscode
`

	// Create a temporary configuration file
	tempConfig := createTempConfigFile(mockConfig)
	defer os.Remove(tempConfig) // Clean up

	//Load the configuration
	config, err := gotoPkg.LoadConfig(tempConfig)
	if err != nil {
		t.Errorf("Error loading the configuration: %v", err)
	}

	if config.MaxDepth != 3 {
		t.Errorf("Expected maxDepth to be 3, got %d", config.MaxDepth)
	}

	if len(config.ExcludeDirs) != 3 {
		t.Errorf("Expected 2 excludeDirs, got %d", len(config.ExcludeDirs))
	}
}

func TestLoadConfigDefault(t *testing.T) {
	//Load the configuration
	config, err := gotoPkg.LoadConfig("nonexistent.yaml")
	if err != nil {
		t.Errorf("Error loading the configuration: %v", err)
	}

	if config.MaxDepth != 5 {
		t.Errorf("Expected maxDepth to be 3, got %d", config.MaxDepth)
	}

	if len(config.ExcludeDirs) != 2 {
		t.Errorf("Expected 2 excludeDirs, got %d", len(config.ExcludeDirs))
	}
}
