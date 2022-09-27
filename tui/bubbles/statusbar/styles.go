package statusbar

import (
	"github.com/charmbracelet/lipgloss"
)

type styles struct {
	containerStyle      lipgloss.Style
	errorMessageStyle   lipgloss.Style
	successMessageStyle lipgloss.Style
}

func defaultStyles() (s styles) {

	s.containerStyle = lipgloss.NewStyle().PaddingLeft(1)
	s.errorMessageStyle = lipgloss.NewStyle()
	s.successMessageStyle = lipgloss.NewStyle()

	return s

}
