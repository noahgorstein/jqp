package help

import (
	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	helpbarStyle       lipgloss.Style
	helpKeyStyle       lipgloss.Style
	helpTextStyle      lipgloss.Style
	helpSeparatorStyle lipgloss.Style
}

func DefaultStyles() (s Styles) {

	s.helpbarStyle = lipgloss.NewStyle().MarginLeft(1).MarginBottom(1)

	s.helpKeyStyle = lipgloss.NewStyle().Bold(true)

	s.helpSeparatorStyle = lipgloss.NewStyle().Bold(true)

	s.helpTextStyle = lipgloss.NewStyle()

	return s
}
