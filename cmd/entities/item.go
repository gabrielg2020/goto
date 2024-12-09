package entities

type Item struct {
	TitleStr string
}

func (i Item) Title() string { return i.TitleStr }

func (i Item) Description() string { return "" }

// Create FilterValue method to fit with List.Item interface
func (i Item) FilterValue() string { return i.TitleStr }
