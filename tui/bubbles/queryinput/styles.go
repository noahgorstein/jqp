package queryinput

import (
	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	containerStyle lipgloss.Style
}

func DefaultStyles() (s Styles) {

	s.containerStyle = lipgloss.NewStyle().Border(lipgloss.RoundedBorder())
	return s
}
