package pkg

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/gabrielg2020/goto/cmd/entities"
	"strings"
)

// Returns the usage message
func Usage() string {
	return `Usage: goto [options] <pattern>
Options:
	-d <depth> Specify the maximum depth to search (default: 5)
	-h         Display this help message
`
}

// Create tea.List items based off a string of directories
func CreateItems(rawDirectories string) []list.Item {
	// Split directories into a slice of strings
	directories := strings.Split(rawDirectories, "\n")

	var items []list.Item

	for _, directory := range directories {
		// If empty -> don't create Item
		if directory == "" || directory == " " {
			continue
		}

		item := entities.Item{Title: directory}
		items = append(items, item)
	}

	return items
}
