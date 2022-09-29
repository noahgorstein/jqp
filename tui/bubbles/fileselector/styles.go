package fileselector

import (
	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	containerStyle  lipgloss.Style
	inputLabelStyle lipgloss.Style
	promptStyle     lipgloss.Style
}

func DefaultStyles() (s Styles) {
	s.containerStyle = lipgloss.NewStyle().Align(lipgloss.Left).PaddingLeft(1)
	s.inputLabelStyle = lipgloss.NewStyle().Bold(true)
	s.promptStyle = lipgloss.NewStyle().Bold(true)

	return s
}
