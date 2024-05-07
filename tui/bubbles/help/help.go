package help

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbletea"

	"github.com/noahgorstein/jqp/tui/bubbles/state"
	"github.com/noahgorstein/jqp/tui/theme"
)

type Bubble struct {
	state  state.State
	width  int
	help   help.Model
	keys   keyMap
	Styles Styles
}

func New(jqtheme theme.Theme) Bubble {
	styles := DefaultStyles()
	model := help.New()
	model.Styles.ShortKey = styles.helpKeyStyle.Foreground(jqtheme.Primary)
	model.Styles.ShortDesc = styles.helpTextStyle.Foreground(jqtheme.Secondary)
	model.Styles.ShortSeparator = styles.helpSeparatorStyle.Foreground(jqtheme.Inactive)

	return Bubble{
		state:  state.Query,
		Styles: styles,
		help:   model,
		keys:   keys,
	}
}

func (b Bubble) collectHelpBindings() []key.Binding {
	k := b.keys
	bindings := []key.Binding{}
	switch b.state {
	case state.Query:
		bindings = append(bindings, k.submit, k.section, k.copyQuery, k.save)
	case state.Running:
		bindings = append(bindings, k.abort)
	case state.Input, state.Output:
		bindings = append(bindings, k.section, k.navigate, k.page, k.copyQuery, k.save)
	case state.Save:
		bindings = append(bindings, k.back)
	}

	return bindings
}

func (b *Bubble) SetWidth(width int) {
	b.Styles.helpbarStyle.Width(width - 1)
}

func (b Bubble) Init() tea.Cmd {
	return nil
}

func (b Bubble) View() string {
	return b.Styles.helpbarStyle.Render(b.help.ShortHelpView(b.collectHelpBindings()))
}

func (b *Bubble) SetState(mode state.State) {
	b.state = mode
}

func (b Bubble) Update(msg tea.Msg) (Bubble, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		b.SetWidth(msg.Width)
	}
	return b, tea.Batch(cmd)
}
