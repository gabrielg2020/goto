package test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	gotoPkg "github.com/gabrielg2020/goto/cmd/pkg"
)

func TestSearchDirectory(t *testing.T) {
	tempDir := t.TempDir()
	// Create a temporary directory structure
	os.Mkdir(filepath.Join(tempDir, "src"), 0755)
	os.Mkdir(filepath.Join(tempDir, ".git"), 0755)
	os.Mkdir(filepath.Join(tempDir, "node_modules"), 0755)

	// Create a temporary configuration variables
	excludeDirs := []string{".git", "node_modules"}
	maxDepth := 1
	
	// Call the search function
	output, err := gotoPkg.SearchDirectories(tempDir, maxDepth, "src", excludeDirs)
	if err != nil {
		t.Errorf("Error searching the directory: %v", err)
	}

	if !strings.Contains(output, "src") {
		t.Errorf("Expected 'src' in the output, got %s", output)
	}

	if strings.Contains(output, ".git") {
		t.Errorf("Expected '.git' to be excluded, got %s", output)
	}

	if strings.Contains(output, "node_modules") {
		t.Errorf("Expected 'node_modules' to be excluded, got %s", output)
	}
}
