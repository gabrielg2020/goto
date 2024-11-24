package main

import (
	"flag"
	"fmt"
	"os"

	gotoPkg "github.com/gabrielg2020/goto/cmd/pkg"
)

func main() {
	// Get the configuration
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	configPath := fmt.Sprintf("%s/.goto_config.yaml", homeDir)
	config, err := gotoPkg.LoadConfig(configPath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Define flags
	maxDepth := flag.Int("d", config.MaxDepth, "Maximum depth to search")
	flag.Parse()

	// Remaining arguments
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("Please provide a pattern")
		return
	}
	pattern := args[0]

	// Call the search and select function
	results, err := gotoPkg.SearchAndSelect(".", *maxDepth, pattern, config.ExcludeDirs)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println(results)
}
