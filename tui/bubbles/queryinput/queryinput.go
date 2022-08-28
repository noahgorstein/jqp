package queryinput

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/noahgorstein/jqp/tui/styles"
)

type Bubble struct {
	Styles    Styles
	textinput textinput.Model
}

func New() Bubble {

	s := DefaultStyles()
	ti := textinput.New()
	ti.Focus()
	ti.BackgroundStyle.Height(1)
	ti.PromptStyle.Height(1)
	ti.TextStyle.Height(1)
	ti.Prompt = lipgloss.NewStyle().Foreground(styles.PINK).Render("jq > ")

	return Bubble{
		Styles:    s,
		textinput: ti,
	}
}

func (b *Bubble) SetBorderColor(color lipgloss.TerminalColor) {
	b.Styles.containerStyle.BorderForeground(color)
}

func (b Bubble) GetInputValue() string {
	return b.textinput.Value()
}

func (b Bubble) Init() tea.Cmd {
	return nil
}

func (b *Bubble) SetWidth(width int) {
	b.textinput.Width = width - b.Styles.containerStyle.GetHorizontalFrameSize() - lipgloss.Width(b.textinput.Prompt) - b.textinput.BackgroundStyle.GetHorizontalFrameSize() - 1
	b.Styles.containerStyle.Width(width - b.Styles.containerStyle.GetHorizontalFrameSize())
}

func (b Bubble) View() string {
	return b.Styles.containerStyle.Render(b.textinput.View())
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
