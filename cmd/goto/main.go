package main

import (
	"fmt"
)

func main() {
	// Call the search and select function
	results, err := SearchAndSelect(".", 0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println(results)
}