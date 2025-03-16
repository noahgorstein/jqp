package inputdata

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/noahgorstein/jqp/tui/theme"
	"github.com/noahgorstein/jqp/tui/utils"
)

type Bubble struct {
	styles               Styles
	viewport             viewport.Model
	height               int
	width                int
	inputJSON            []byte
	highlightedJSON      *bytes.Buffer
	filename             string
	theme                theme.Theme
	setInitialContentSub chan setPrettifiedContentMsg
}

func New(inputJSON []byte, filename string, jqtheme theme.Theme) (Bubble, error) {
	styles := DefaultStyles()
	styles.containerStyle = styles.containerStyle.BorderForeground(jqtheme.Inactive)
	styles.infoStyle = styles.infoStyle.BorderForeground(jqtheme.Inactive)

	v := viewport.New(0, 0)
	v.SetContent("Loading...")

	b := Bubble{
		styles:               styles,
		viewport:             v,
		inputJSON:            inputJSON,
		filename:             filename,
		theme:                jqtheme,
		setInitialContentSub: make(chan setPrettifiedContentMsg),
	}
	return b, nil
}

func (b *Bubble) SetBorderColor(color lipgloss.TerminalColor) {
	b.styles.containerStyle = b.styles.containerStyle.BorderForeground(color)
	b.styles.infoStyle = b.styles.infoStyle.BorderForeground(color)
}

func (b Bubble) GetInputJSON() []byte {
	return b.inputJSON
}

func (b Bubble) GetHighlightedInputJSON() []byte {
	return b.highlightedJSON.Bytes()
}

func (b *Bubble) SetSize(width, height int) {
	b.width = width
	b.height = height

	b.styles.containerStyle = b.styles.containerStyle.
		Width(width - b.styles.containerStyle.GetHorizontalFrameSize()/2).
		Height(height - b.styles.containerStyle.GetVerticalFrameSize())

	b.viewport.Width = width - b.styles.containerStyle.GetHorizontalFrameSize()
	b.viewport.Height = height - b.styles.containerStyle.GetVerticalFrameSize() - 3
}

func (b Bubble) View() string {
	scrollPercent := fmt.Sprintf("%3.f%%", b.viewport.ScrollPercent()*100)

	info := b.styles.infoStyle.Render(fmt.Sprintf("%s | %s", lipgloss.NewStyle().Italic(true).Render(b.filename), scrollPercent))
	line := strings.Repeat(" ", max(0, b.viewport.Width-lipgloss.Width(info)))

	footer := lipgloss.JoinHorizontal(lipgloss.Center, line, info)
	content := lipgloss.JoinVertical(lipgloss.Left, b.viewport.View(), footer)

	return b.styles.containerStyle.Render(content)
}

func (b *Bubble) SetContent(content string) {
	formattedContent := lipgloss.NewStyle().Width(b.viewport.Width - 3).Render(content)
	b.viewport.SetContent(formattedContent)
}

// ReadyMsg signals that the inputdata Bubble has loaded the user's data
// into the viewport
type ReadyMsg struct{}

// setPrettifiedContentMsg contains the input data prettified
type setPrettifiedContentMsg struct {
	Content *bytes.Buffer
}

// prettifyContentCmd sends the initial prettified content to the provided channel.
//
// Prettifying the input data can be an expensive operation particularly for large inputs, so it is performed here and
// sent through the channel to ensure the prettified data is available without blocking other operations.
func (b Bubble) prettifyContentCmd(sub chan setPrettifiedContentMsg, isJSONLines bool) tea.Cmd {
	return func() tea.Msg {
		prettifiedData, _ := utils.Prettify(b.inputJSON, b.theme.ChromaStyle, isJSONLines)
		sub <- setPrettifiedContentMsg{Content: prettifiedData}
		return nil
	}
}

// A command that waits for a setPrettifiedContentMsg on a channel.
func waitForPrettifiedContent(sub chan setPrettifiedContentMsg) tea.Cmd {
	return func() tea.Msg {
		return setPrettifiedContentMsg(<-sub)
	}
}

func (b Bubble) Init(isJSONLines bool) tea.Cmd {
	return tea.Batch(
		b.prettifyContentCmd(b.setInitialContentSub, isJSONLines),
		waitForPrettifiedContent(b.setInitialContentSub))
}

func (b Bubble) Update(msg tea.Msg) (Bubble, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	if msg, ok := msg.(setPrettifiedContentMsg); ok {
		b.highlightedJSON = msg.Content
		b.SetContent(msg.Content.String())
		return b, func() tea.Msg {
			return ReadyMsg{}
		}
	}

	b.viewport, cmd = b.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return b, tea.Batch(cmds...)
}
