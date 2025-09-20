//nolint:revive // utils is an acceptable package name for shared utilities
package utils

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

var termenvChromaTerminal = map[termenv.Profile]string{
	termenv.Ascii:     "terminal",
	termenv.ANSI:      "terminal16",
	termenv.ANSI256:   "terminal256",
	termenv.TrueColor: "terminal16m",
}

// returns a string used for chroma syntax highlighting
func getTerminalColorSupport() string {
	if chroma, ok := termenvChromaTerminal[lipgloss.ColorProfile()]; ok {
		return chroma
	}
	return "terminal"
}
