package jqplayground

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/noahgorstein/jqp/tui/bubbles/state"
)

func (b Bubble) View() string {
	inputoutput := []string{b.inputdata.View()}
	if b.width%2 != 0 {
		inputoutput = append(inputoutput, " ")
	}
	inputoutput = append(inputoutput, b.output.View())

	if b.state == state.Save {
		return lipgloss.JoinVertical(
			lipgloss.Left,
			b.queryinput.View(),
			lipgloss.JoinHorizontal(lipgloss.Top, inputoutput...),
			b.fileselector.View(),
			b.statusbar.View(),
			b.help.View())
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		b.queryinput.View(),
		lipgloss.JoinHorizontal(lipgloss.Top, inputoutput...),
		b.statusbar.View(),
		b.help.View())
}
