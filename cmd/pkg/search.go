package pkg

import (
	"bytes"
	"fmt"
	"os/exec"
)

// Uses the `find` command to list directories up to a certain depth
func SearchDirectories(path string, maxDepth int, pattern string, excludeDirs []string) (string, error) {
	// Construct `find` command
	findArgs := []string{path, "-type", "d", "-maxdepth", fmt.Sprintf("%d", maxDepth), "-name", fmt.Sprintf("*%s*", pattern), "!", "-name", "."}
	for _, dir := range excludeDirs {
		findArgs = append(findArgs, "!", "-path", fmt.Sprintf("*%s*", dir))
	}
	cmd := exec.Command("find", findArgs...)

	// Capture the output
	var out bytes.Buffer
	cmd.Stdout = &out

	// Run the command
	err := cmd.Run()
	if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() != 1 {
		return "", fmt.Errorf("error running the find command: %v", err)
	}

	// Check if no directories were found
	if out.Len() == 0 {
		return "", fmt.Errorf("no directories found")
	}

	return out.String(), nil
}
