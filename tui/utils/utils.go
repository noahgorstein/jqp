package utils

import (
	"io"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/formatters"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

func HighlightJson(w io.Writer, source string, style *chroma.Style) error {
	l := lexers.Get("json")
	if l == nil {
		l = lexers.Fallback
	}
	l = chroma.Coalesce(l)

	f := formatters.Get(getTerminalColorSupport())
	if f == nil {
		f = formatters.Fallback
	}

	if style == nil {
		style = styles.Fallback
	}

	it, err := l.Tokenise(nil, source)
	if err != nil {
		return err
	}
	return f.Format(w, style, it)
}

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
