package jqplayground

import (
	"github.com/charmbracelet/bubbletea"
)

func (b Bubble) Init() tea.Cmd {
	var cmds []tea.Cmd
	if b.queryinput.GetInputValue() != "" {
		b.executeQuery(&cmds)
	}

	setInputDataContentCmd := b.setInputDataContentCommand(b.inputdata.GetHighlightedInputJSON())
	cmds = append(cmds, b.queryinput.Init(), setInputDataContentCmd)
	return tea.Batch(cmds...)
}
