package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// Uses the `find` command to list directories up to a certain depth
func SearchDirectories(path string, maxDepth int, pattern string) (string, error) {
	// Construct `find` command
	findArgs := []string{path, "-type", "d", "-maxdepth", fmt.Sprintf("%d", maxDepth), "-name", fmt.Sprintf("*%s*", pattern), "!", "-name", "."}
	cmd := exec.Command("find", findArgs...)

	// Capture the output
	var out bytes.Buffer
	cmd.Stdout = &out

	// Run the command
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("error running the find command: %v", err)
	}

	// Check if no directories were found
	if out.Len() == 0 {
		return "", fmt.Errorf("no directories found")
	}

	return out.String(), nil
}

// Uses the `fzf` command to fuzzy select a directory
func FuzzySelectDirectory(directories string, pattern string) (string, error) {
	// Construct `fzf` command
	fzfArgs := []string{"--height", "50%", "--reverse", "--query", pattern, "--border", "rounded", "--header", "Select a directory"}
	cmd := exec.Command("fzf", fzfArgs...)

	// Set the input
	cmd.Stdin = bytes.NewBufferString(directories)

	// Capture the output
	var out bytes.Buffer
	cmd.Stdout = &out

	// Run the command
	if err := cmd.Run(); err != nil {
		// Check if error is due to cancelled selection
		if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() == 130 {
			return "", fmt.Errorf("selection cancelled")
		}

		return "", fmt.Errorf("error running the fzf command: %v", err)
	}

	// Remove leading and trailing whitespace
	return strings.TrimSpace(out.String()), nil
}

// Combines the search and fuzzy selection
func SearchAndSelect(path string, maxDepth int, pattern string) (string, error) {
	// Call the search function
	directories, err := SearchDirectories(path, maxDepth, pattern)
	if err != nil {
		return "", err
	}

	// Call the fuzzy selection function
	selectedDir, err := FuzzySelectDirectory(directories, pattern)
	if err != nil {
		return "", err
	}

	return selectedDir, nil
}
