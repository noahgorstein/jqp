package jqplayground

import (
	"github.com/charmbracelet/lipgloss"

	"github.com/noahgorstein/jqp/tui/bubbles/state"
)

func (b Bubble) View() string {
	var inputoutput []string

	if b.showInputPanel {
		inputoutput = []string{b.inputdata.View()}
		if b.width%2 != 0 {
			inputoutput = append(inputoutput, " ")
		}
		inputoutput = append(inputoutput, b.output.View())
	} else {
		inputoutput = []string{b.output.View()}
	}

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
