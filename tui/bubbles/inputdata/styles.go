package inputdata

import (
	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	infoStyle      lipgloss.Style
	containerStyle lipgloss.Style
}

func DefaultStyles() (s Styles) {
	s.infoStyle = lipgloss.NewStyle().Bold(true).Border(lipgloss.RoundedBorder()).Padding(0, 1)
	s.containerStyle = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(1)

	return s
}
