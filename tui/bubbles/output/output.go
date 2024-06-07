package output

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/noahgorstein/jqp/tui/theme"
)

type Bubble struct {
	Ready    bool
	Styles   Styles
	viewport viewport.Model
	content  string
	height   int
	width    int
}

func New(jqtheme theme.Theme) Bubble {
	styles := DefaultStyles()
	styles.containerStyle = styles.containerStyle.BorderForeground(jqtheme.Inactive)
	styles.infoStyle = styles.infoStyle.BorderForeground(jqtheme.Inactive)
	v := viewport.New(1, 1)
	b := Bubble{
		Styles:   styles,
		viewport: v,
		content:  "",
	}
	return b
}

func (b *Bubble) SetBorderColor(color lipgloss.TerminalColor) {
	b.Styles.containerStyle = b.Styles.containerStyle.BorderForeground(color)
	b.Styles.infoStyle = b.Styles.infoStyle.BorderForeground(color)
}

func (b *Bubble) SetSize(width, height int) {
	b.width = width
	b.height = height

	b.Styles.containerStyle.
		Width(width - b.Styles.containerStyle.GetHorizontalFrameSize()/2).
		Height(height - b.Styles.containerStyle.GetVerticalFrameSize())

	b.viewport.Width = width - b.Styles.containerStyle.GetHorizontalFrameSize()
	b.viewport.Height = height - b.Styles.containerStyle.GetVerticalFrameSize() - 3
}

func (b *Bubble) GetContent() string {
	return b.content
}

func (b *Bubble) SetContent(content string) {
	b.content = content
	wrappedContent := lipgloss.NewStyle().Width(b.viewport.Width - 1).Render(content)

	b.viewport.SetContent(wrappedContent)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (b *Bubble) ScrollToTop() {
	b.viewport.GotoTop()
}

func (b Bubble) View() string {
	scrollPercent := fmt.Sprintf("%3.f%%", b.viewport.ScrollPercent()*100)

	info := b.Styles.infoStyle.Render(fmt.Sprintf("%s | %s", lipgloss.NewStyle().Italic(true).Render("output"), scrollPercent))
	line := strings.Repeat(" ", max(0, b.viewport.Width-lipgloss.Width(info)))

	footer := lipgloss.JoinHorizontal(lipgloss.Center, line, info)
	content := lipgloss.JoinVertical(lipgloss.Left, b.viewport.View(), footer)

	return b.Styles.containerStyle.Render(content)
}

func (Bubble) Init() tea.Cmd {
	return nil
}

func (b Bubble) Update(msg tea.Msg) (Bubble, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	b.viewport, cmd = b.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return b, tea.Batch(cmds...)
}
