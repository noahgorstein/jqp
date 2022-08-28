package output

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/noahgorstein/jqp/tui/styles"
)

type Styles struct {
	infoStyle      lipgloss.Style
	containerStyle lipgloss.Style
}

func DefaultStyles() (s Styles) {
	s.infoStyle = lipgloss.NewStyle().Bold(true).Border(lipgloss.RoundedBorder()).BorderForeground(styles.GREY).Padding(0, 1)
	s.containerStyle = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("240")).Padding(1)

	return s

}
