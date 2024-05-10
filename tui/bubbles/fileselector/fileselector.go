package fileselector

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/noahgorstein/jqp/tui/theme"
)

type Bubble struct {
	Styles    Styles
	textinput textinput.Model
}

func New(jqtheme theme.Theme) Bubble {
	s := DefaultStyles()
	ti := textinput.New()
	ti.Focus()
	ti.PromptStyle = s.promptStyle.Foreground(jqtheme.Secondary)
	s.inputLabelStyle.Foreground(jqtheme.Primary)

	return Bubble{
		Styles:    s,
		textinput: ti,
	}
}

func (b Bubble) GetInput() string {
	return b.textinput.Value()
}

func (b *Bubble) SetInput(input string) {
	b.textinput.SetValue(input)
}

func (Bubble) Init() tea.Cmd {
	return nil
}

func (b Bubble) View() string {
	return b.Styles.containerStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			b.Styles.inputLabelStyle.Render("Save output to file (leave empty to copy to clipboard): "),
			b.textinput.View()))
}

func (b Bubble) Update(msg tea.Msg) (Bubble, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	b.textinput, cmd = b.textinput.Update(msg)
	cmds = append(cmds, cmd)

	return b, tea.Batch(cmds...)
}

func (b Bubble) SetSize(width int) {
	b.Styles.containerStyle.
		Width(width - b.Styles.containerStyle.GetHorizontalFrameSize())
}
