package help

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/noahgorstein/jqp/tui/styles"
)

type Styles struct {
	helpbarStyle       lipgloss.Style
	helpKeyStyle       lipgloss.Style
	helpTextStyle      lipgloss.Style
	helpSeparatorStyle lipgloss.Style
}

func DefaultStyles() (s Styles) {

	s.helpbarStyle = lipgloss.NewStyle().MarginLeft(1).MarginBottom(1)

	s.helpKeyStyle = lipgloss.NewStyle().Bold(true).Foreground(
		lipgloss.AdaptiveColor{
			Light: string(styles.BLUE),
			Dark:  string(styles.BLUE),
		},
	)

	s.helpSeparatorStyle = lipgloss.NewStyle().Bold(true).Foreground(
		lipgloss.AdaptiveColor{
			Light: string(styles.GREY),
			Dark:  string(styles.GREY),
		},
	)

	s.helpTextStyle = lipgloss.NewStyle().Foreground(
		lipgloss.AdaptiveColor{
			Light: string(styles.PINK),
			Dark:  string(styles.PINK),
		},
	)

	return s
}
