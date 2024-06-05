package jqplayground

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/noahgorstein/jqp/tui/utils"
)

// invalidInputMsg signals that the user's data is not valid JSON or NDJSON
type invalidInputMsg struct{}

type setupMsg struct {
	isJSONLines bool
}

// initialQueryMsg represents a message containing an initial query string to execute when
// the app is loaded.
type initialQueryMsg struct {
	query string
}

func setupCmd(isJSONLines bool) tea.Cmd {
	return func() tea.Msg {
		return setupMsg{isJSONLines: isJSONLines}
	}
}

// initialQueryCmd creates a command that returns an initialQueryMsg with the provided query string.
func initialQueryCmd(query string) tea.Cmd {
	return func() tea.Msg {
		return initialQueryMsg{query: query}
	}
}

func (b Bubble) Init() tea.Cmd {
	var cmds []tea.Cmd

	// validate input data
	_, isJSONLines, err := utils.IsValidInput(b.inputdata.GetInputJSON())
	if err != nil {
		return func() tea.Msg {
			return invalidInputMsg{}
		}
	}

	// initialize rest of app
	cmds = append(cmds, b.queryinput.Init(), b.inputdata.Init(isJSONLines), setupCmd(isJSONLines))
	if b.queryinput.GetInputValue() != "" {
		cmds = append(cmds, initialQueryCmd(b.queryinput.GetInputValue()))
	}
	return tea.Sequence(cmds...)
}
