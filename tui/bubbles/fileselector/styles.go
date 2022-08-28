package fileselector

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/noahgorstein/jqp/tui/styles"
)

type Styles struct {
	containerStyle  lipgloss.Style
	inputLabelStyle lipgloss.Style
	promptStyle     lipgloss.Style
}

func DefaultStyles() (s Styles) {
	s.containerStyle = lipgloss.NewStyle().Align(lipgloss.Left).PaddingLeft(1)
	s.inputLabelStyle = lipgloss.NewStyle().Bold(true).Foreground(styles.BLUE)
	s.promptStyle = lipgloss.NewStyle().Bold(true).Foreground(styles.PINK)

	return s
}
