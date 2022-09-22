package help

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	section   key.Binding
	back      key.Binding
	submit    key.Binding
	navigate  key.Binding
	page      key.Binding
	save      key.Binding
	copyQuery key.Binding
}

var keys = keyMap{
	section: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "section")),
	back: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "back")),
	submit: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "submit query")),
	navigate: key.NewBinding(
		key.WithKeys("↑↓"),
		key.WithHelp("↑↓", "scroll")),
	page: key.NewBinding(
		key.WithKeys("ctrl+u/ctrl+d"),
		key.WithHelp("ctrl+u/ctrl+d", "page up/down")),
	save: key.NewBinding(
		key.WithKeys("ctrl+s"),
		key.WithHelp("ctrl+s", "save output")),
	copyQuery: key.NewBinding(
		key.WithKeys("ctrl+y"),
		key.WithHelp("ctrl+y", "copy query")),
}
