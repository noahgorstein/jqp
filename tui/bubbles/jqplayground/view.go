package jqplayground

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/noahgorstein/jqp/tui/bubbles/state"
)

func (b Bubble) View() string {

	if b.state == state.Save {
		return lipgloss.JoinVertical(
			lipgloss.Left,
			b.queryinput.View(),
			lipgloss.JoinHorizontal(lipgloss.Top, b.inputdata.View(), b.output.View()),
			b.fileselector.View(),
			b.statusbar.View(),
			b.help.View())
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		b.queryinput.View(),
		lipgloss.JoinHorizontal(lipgloss.Top, b.inputdata.View(), b.output.View()),
		b.statusbar.View(),
		b.help.View())
}
