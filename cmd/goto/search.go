package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

// Uses the `find` command to list directories up to a certain depth
func SearchDirectories(path string, maxDepth int) (string, error) {
	// Construct `find` command
	cmd := exec.Command("find", path, "-type", "d", "-maxdepth", fmt.Sprintf("%d", maxDepth))

	// Capture the output
	var out bytes.Buffer
	cmd.Stdout = &out

	// Run the command
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("error running the find command: %v", err)
	}

	return out.String(), nil
}
