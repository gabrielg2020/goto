package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gabrielg2020/goto/cmd/pkg"
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
	maxDepth := flag.Int("d", config.MaxDepth, "Maximum depth to search") // -d flag
	flag.Usage = func() {                                                 // -h flag
		fmt.Fprintf(os.Stderr, "%s\n", pkg.Usage())
	}

	flag.Parse()

	// Remaining arguments
	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
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
