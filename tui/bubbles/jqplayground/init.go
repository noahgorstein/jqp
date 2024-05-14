package jqplayground

import (
	"github.com/charmbracelet/bubbletea"
)

func (b Bubble) Init() tea.Cmd {
	var cmds []tea.Cmd
	if b.queryinput.GetInputValue() != "" {
		b.executeQuery(&cmds)
	}
	cmds = append(cmds, b.queryinput.Init())
	return tea.Batch(cmds...)
}
