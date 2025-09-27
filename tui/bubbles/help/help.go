package help

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbletea"

	"github.com/noahgorstein/jqp/tui/bubbles/state"
	"github.com/noahgorstein/jqp/tui/theme"
)

type Bubble struct {
	state          state.State
	help           help.Model
	keys           keyMap
	Styles         Styles
	showInputPanel bool
}

func New(jqtheme theme.Theme) Bubble {
	styles := DefaultStyles()
	model := help.New()
	model.Styles.ShortKey = styles.helpKeyStyle.Foreground(jqtheme.Primary)
	model.Styles.ShortDesc = styles.helpTextStyle.Foreground(jqtheme.Secondary)
	model.Styles.ShortSeparator = styles.helpSeparatorStyle.Foreground(jqtheme.Inactive)

	return Bubble{
		state:          state.Query,
		Styles:         styles,
		help:           model,
		keys:           keys,
		showInputPanel: true, // Default to showing input panel
	}
}

//nolint:revive // switch statement complexity is acceptable here
func (b Bubble) collectHelpBindings() []key.Binding {
	k := b.keys
	bindings := []key.Binding{}

	// Create dynamic toggle binding based on panel visibility
	toggleText := "show input panel"
	if b.showInputPanel {
		toggleText = "hide input panel"
	}
	toggleBinding := key.NewBinding(
		key.WithKeys("ctrl+t"),
		key.WithHelp("ctrl+t", toggleText),
	)

	switch b.state {
	case state.Query:
		bindings = append(bindings, k.submit, k.section, k.copyQuery, toggleBinding, k.save)
	case state.Running:
		bindings = append(bindings, k.abort)
	case state.Input, state.Output:
		bindings = append(bindings, k.section, k.navigate, k.page, k.copyQuery, toggleBinding, k.save)
	case state.Save:
		bindings = append(bindings, k.back)
	default:
		// No additional bindings for unknown states
	}

	return bindings
}

func (b *Bubble) SetWidth(width int) {
	b.Styles.helpbarStyle = b.Styles.helpbarStyle.Width(width - 1)
}

func (Bubble) Init() tea.Cmd {
	return nil
}

func (b Bubble) View() string {
	return b.Styles.helpbarStyle.Render(b.help.ShortHelpView(b.collectHelpBindings()))
}

func (b *Bubble) SetState(mode state.State) {
	b.state = mode
}

func (b *Bubble) SetInputPanelVisibility(visible bool) {
	b.showInputPanel = visible
}

func (b Bubble) Update(msg tea.Msg) (Bubble, tea.Cmd) {
	var cmd tea.Cmd

	if msg, ok := msg.(tea.WindowSizeMsg); ok {
		b.SetWidth(msg.Width)
	}

	return b, tea.Batch(cmd)
}
