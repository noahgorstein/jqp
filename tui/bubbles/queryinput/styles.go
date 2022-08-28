package queryinput

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/noahgorstein/jqp/tui/styles"
)

type Styles struct {
	containerStyle lipgloss.Style
}

func DefaultStyles() (s Styles) {

	s.containerStyle = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(styles.BLUE)
	return s
}
