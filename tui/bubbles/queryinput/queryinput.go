package queryinput

import (
	"container/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/noahgorstein/jqp/tui/theme"
	"os"
	"bufio"
)

type Bubble struct {
	Styles    Styles
	textinput textinput.Model

	history         *list.List
	historyMaxLen   int
	historySelected *list.Element
}

func New(theme theme.Theme) Bubble {

	s := DefaultStyles()
	s.containerStyle.BorderForeground(theme.Primary)
	ti := textinput.New()
	ti.Focus()
	ti.PromptStyle.Height(1)
	ti.TextStyle.Height(1)
	ti.Prompt = lipgloss.NewStyle().Bold(true).Foreground(theme.Secondary).Render("jq > ")

	// set to empty pointer
	historySelected := (*list.Element)(nil)
	historyList := list.New()

	if history, err := loadHistory(".jqp_history"); err == nil {
		for _, entry := range history {
			historyList.PushBack(entry)
		}

		historySelected = historyList.Front()
	}

	return Bubble{
		Styles:    s,
		textinput: ti,
		history:       historyList,
		historySelected: historySelected,
		historyMaxLen: 512,
	}
}

func loadHistory(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
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

	updateHistoryFile(".jqp_history", b.history)
}

func updateHistoryFile(filename string, history *list.List) {
	file, err := os.Create(filename)
	if err != nil {
		return
	}
	defer file.Close()

	// remove duplicates from history
	seen := make(map[string]bool)

	writer := bufio.NewWriter(file)
	for e := history.Front(); e != nil; e = e.Next() {
		value, ok := e.Value.(string)
		if ok && !seen[value] {
			writer.WriteString(value + "\n")
			seen[value] = true
		}
	}
	writer.Flush()
}


func (b Bubble) Init() tea.Cmd {
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
	if msg, ok := msg.(tea.KeyMsg); ok {
		switch msg.Type {
		case tea.KeyUp:
			if b.history.Len() == 0 {
				break
			}
			n := b.historySelected.Next()
			if n != nil {
				b.textinput.SetValue(n.Value.(string))
				b.textinput.CursorEnd()
				b.historySelected = n
			}
		case tea.KeyDown:
			if b.history.Len() == 0 {
				break
			}
			p := b.historySelected.Prev()
			if p != nil {
				b.textinput.SetValue(p.Value.(string))
				b.textinput.CursorEnd()
				b.historySelected = p
			}
		case tea.KeyEnter:
			b.RotateHistory()
		}
	}

	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	b.textinput, cmd = b.textinput.Update(msg)
	cmds = append(cmds, cmd)

	return b, tea.Batch(cmds...)

}
