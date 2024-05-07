package statusbar

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/noahgorstein/jqp/tui/theme"
)

type Bubble struct {
	styles                styles
	StatusMessageLifetime time.Duration
	statusMessage         string
	statusMessageTimer    *time.Timer
}

func (b Bubble) Init() tea.Cmd {
	return nil
}

func (b Bubble) View() string {
	return b.styles.containerStyle.Render(b.statusMessage)
}

func (b *Bubble) SetSize(width int) {
	b.styles.containerStyle.Width(width)
}

func (b *Bubble) hideStatusMessage() {
	b.statusMessage = ""
	if b.statusMessageTimer != nil {
		b.statusMessageTimer.Stop()
	}
}

func (b Bubble) Update(msg tea.Msg) (Bubble, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case statusMessageTimeoutMsg:
		b.hideStatusMessage()
	case tea.WindowSizeMsg:
		b.SetSize(msg.Width)
	}
	return b, tea.Batch(cmd)
}

func New(jqtheme theme.Theme) Bubble {
	styles := defaultStyles()
	styles.successMessageStyle = styles.successMessageStyle.Foreground(jqtheme.Success)
	styles.errorMessageStyle = styles.errorMessageStyle.Foreground(jqtheme.Error)
	b := Bubble{
		styles: styles,
	}
	return b
}

type statusMessageTimeoutMsg struct{}

func (b *Bubble) NewStatusMessage(s string, success bool) tea.Cmd {
	if success {
		b.statusMessage = b.styles.successMessageStyle.Render(s)
	} else {
		b.statusMessage = b.styles.errorMessageStyle.Render(s)
	}

	if b.statusMessageTimer != nil {
		b.statusMessageTimer.Stop()
	}

	b.statusMessageTimer = time.NewTimer(b.StatusMessageLifetime)

	// Wait for timeout
	return func() tea.Msg {
		<-b.statusMessageTimer.C
		return statusMessageTimeoutMsg{}
	}
}
