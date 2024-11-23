package main

import (
	"flag"
	"fmt"
)

func main() {
	// Define flags
	maxDepth := flag.Int("d", 5, "Maximum depth to search")
	flag.Parse()

	// Remaining arguments
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("Please provide a pattern")
		return
	}
	pattern := args[0]

	// Call the search and select function
	results, err := SearchAndSelect(".", *maxDepth, pattern)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println(results)
}