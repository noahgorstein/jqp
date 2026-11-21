package queryinput

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/noahgorstein/jqp/tui/theme"
)

func TestLongQueryFitsOnOneLine(t *testing.T) {
	const width = 24
	query := "." + strings.Repeat("abcdefghijklmnop", 4)
	b := New(theme.Theme{})
	b.SetWidth(width)
	b, _ = b.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune(query)})

	checkQueryBox(t, b.View(), "end", width, "klmnop")
	b, _ = b.Update(tea.KeyMsg{Type: tea.KeyHome})
	checkQueryBox(t, b.View(), "home", width, ".abc")
	b, _ = b.Update(tea.KeyMsg{Type: tea.KeyRight})
	b.SetWidth(width / 2)
	checkQueryBox(t, b.View(), "resized with cursor near start", width/2, "abcd")
	if got := b.textinput.Position(); got != 1 {
		t.Errorf("cursor position after resize = %d, want 1", got)
	}
	b, _ = b.Update(tea.KeyMsg{Type: tea.KeyEnd})
	checkQueryBox(t, b.View(), "end after resize", width/2, "mnop")
	if got := b.GetInputValue(); got != query {
		t.Errorf("query changed while scrolling: got %q, want %q", got, query)
	}
}

func checkQueryBox(t *testing.T, view, label string, width int, visible string) {
	t.Helper()
	if got := lipgloss.Height(view); got != 3 {
		t.Errorf("%s: query box height = %d, want 3 (one row and borders)", label, got)
	}
	if got := lipgloss.Width(view); got != width {
		t.Errorf("%s: query box width = %d, want %d", label, got, width)
	}
	if !strings.Contains(view, visible) {
		t.Errorf("%s: query box does not show %q: %q", label, visible, view)
	}
}
