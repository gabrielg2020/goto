package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gabrielg2020/goto/cmd/entities"
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

	// Grab directories
	directories, err := gotoPkg.SearchDirectories(".", *maxDepth, pattern, config.ExcludeDirs)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Create a slice of tea items
	items := gotoPkg.CreateItems(directories)

	// Start model
	model := entities.Model{List: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	model.List.Title = "Pick a Directory"

	program := tea.NewProgram(model, tea.WithAltScreen())

	finalModel, err := program.Run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if fm, ok := finalModel.(entities.Model); ok && fm.Choice != "" {
		fmt.Println(fm.Choice)
	} else {
		fmt.Println("")
	}
}
