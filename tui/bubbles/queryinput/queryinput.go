package queryinput

import (
	"container/list"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/noahgorstein/jqp/tui/theme"
)

type Bubble struct {
	Styles    Styles
	textinput textinput.Model

	history         *list.List
	historyMaxLen   int
	historySelected *list.Element
}

func New(jqtheme theme.Theme) Bubble {
	s := DefaultStyles()
	s.containerStyle.BorderForeground(jqtheme.Primary)
	ti := textinput.New()
	ti.Focus()
	ti.PromptStyle.Height(1)
	ti.TextStyle.Height(1)
	ti.Prompt = lipgloss.NewStyle().Bold(true).Foreground(jqtheme.Secondary).Render("jq > ")

	return Bubble{
		Styles:    s,
		textinput: ti,

		history:       list.New(),
		historyMaxLen: 512,
	}
}

func (b *Bubble) SetBorderColor(color lipgloss.TerminalColor) {
	b.Styles.containerStyle.BorderForeground(color)
}

func (b Bubble) GetInputValue() string {
	return b.textinput.Value()
}

func (b *Bubble) RotateHistory() {
	b.history.PushFront(b.textinput.Value())
	b.historySelected = b.history.Front()
	if b.history.Len() > b.historyMaxLen {
		b.history.Remove(b.history.Back())
	}
}

func (Bubble) Init() tea.Cmd {
	return textinput.Blink
}

func (b *Bubble) SetWidth(width int) {
	b.Styles.containerStyle.Width(width - b.Styles.containerStyle.GetHorizontalFrameSize())
	b.textinput.Width = width - b.Styles.containerStyle.GetHorizontalFrameSize() - 1
}

func (b Bubble) View() string {
	return b.Styles.containerStyle.Render(b.textinput.View())
}

func (b Bubble) Update(msg tea.Msg) (Bubble, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return b.updateKeyMsg(msg)
	default:
		var cmd tea.Cmd
		b.textinput, cmd = b.textinput.Update(msg)
		return b, cmd
	}
}

func (b *Bubble) SetQuery(query string) {
	b.textinput.SetValue(query)
}

func (b Bubble) updateKeyMsg(msg tea.KeyMsg) (Bubble, tea.Cmd) {
	switch msg.Type {
	case tea.KeyUp:
		return b.handleKeyUp()
	case tea.KeyDown:
		return b.handleKeyDown()
	case tea.KeyEnter:
		b.RotateHistory()
		return b, nil
	default:
		var cmd tea.Cmd
		b.textinput, cmd = b.textinput.Update(msg)
		return b, cmd
	}
}

func (b Bubble) handleKeyUp() (Bubble, tea.Cmd) {
	if b.history.Len() == 0 {
		return b, nil
	}
	n := b.historySelected.Next()
	if n != nil {
		b.textinput.SetValue(n.Value.(string))
		b.textinput.CursorEnd()
		b.historySelected = n
	}
	return b, nil
}

func (b Bubble) handleKeyDown() (Bubble, tea.Cmd) {
	if b.history.Len() == 0 {
		return b, nil
	}
	p := b.historySelected.Prev()
	if p != nil {
		b.textinput.SetValue(p.Value.(string))
		b.textinput.CursorEnd()
		b.historySelected = p
	}
	return b, nil
}
