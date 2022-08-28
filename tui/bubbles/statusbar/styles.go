package statusbar

import (
	"github.com/charmbracelet/lipgloss"
	jqp_styles "github.com/noahgorstein/jqp/tui/styles"
)

type styles struct {
	containerStyle      lipgloss.Style
	errorMessageStyle   lipgloss.Style
	successMessageStyle lipgloss.Style
}

func defaultStyles() (s styles) {

	s.containerStyle = lipgloss.NewStyle().PaddingLeft(1)
	s.errorMessageStyle = lipgloss.NewStyle().Foreground(jqp_styles.RED)
	s.successMessageStyle = lipgloss.NewStyle().Foreground(jqp_styles.GREEN)

	return s

}
