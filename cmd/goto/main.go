package main

import (
	"fmt"
)

func main() {
	// Call the search function
	results, err := SearchDirectories(".", 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println(results)
}