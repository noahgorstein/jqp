package statusbar

import (
	tea "github.com/charmbracelet/bubbletea"
	"time"
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

func New() Bubble {
	styles := defaultStyles()
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
