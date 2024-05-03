package utils

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

// returns a string used for chroma syntax highlighting
func getTerminalColorSupport() string {
	switch lipgloss.ColorProfile() {
	case termenv.Ascii:
		return "terminal"
	case termenv.ANSI:
		return "terminal16"
	case termenv.ANSI256:
		return "terminal256"
	case termenv.TrueColor:
		return "terminal16m"
	default:
		return "terminal"
	}
}
