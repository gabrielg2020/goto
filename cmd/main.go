package main

import (
	"flag"
	"fmt"

	gotoPkg "github.com/gabrielg2020/goto/cmd/pkg"
)

func main() {
	// Get the configuration
	config, err := gotoPkg.LoadConfig()
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
