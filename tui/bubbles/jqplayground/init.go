package jqplayground

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/noahgorstein/jqp/tui/utils"
)

// InvalidInputMsg signals that the user's data is not valid JSON or NDJSON
type InvalidInputMsg struct{}

func (b Bubble) Init() tea.Cmd {
	var cmds []tea.Cmd

	// validate input data
	_, isJSONLines, err := utils.IsValidInput(b.inputdata.GetInputJSON())
	if err != nil {
		return func() tea.Msg {
			return InvalidInputMsg{}
		}
	}

	// initialize rest of app
	b.isJSONLines = isJSONLines
	cmds = append(cmds, b.queryinput.Init(), b.inputdata.Init(isJSONLines))
	if b.queryinput.GetInputValue() != "" {
		b.executeQuery(&cmds)
	}
	return tea.Sequence(cmds...)
}
