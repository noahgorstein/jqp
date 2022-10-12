package inputdata

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/alecthomas/chroma/v2"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/noahgorstein/jqp/tui/theme"
	"github.com/noahgorstein/jqp/tui/utils"
)

type Bubble struct {
	Styles          Styles
	viewport        viewport.Model
	height          int
	width           int
	inputJson       []byte
	highlightedJson *bytes.Buffer
	filename        string
}

func New(inputJson []byte, filename string, theme theme.Theme) Bubble {
	styles := DefaultStyles()
	styles.containerStyle = styles.containerStyle.BorderForeground(theme.Inactive)
	styles.infoStyle = styles.infoStyle.BorderForeground(theme.Inactive)
	v := viewport.New(0, 0)
	b := Bubble{
		Styles:          styles,
		viewport:        v,
		inputJson:       inputJson,
		highlightedJson: highlightInputJson(inputJson, theme.ChromaStyle),
		filename:        filename,
	}
	return b
}

func (b *Bubble) SetBorderColor(color lipgloss.TerminalColor) {
	b.Styles.containerStyle.BorderForeground(color)
	b.Styles.infoStyle.BorderForeground(color)
}

func highlightInputJson(inputJson []byte, chromaStyle *chroma.Style) *bytes.Buffer {
	var f interface{}
	// TODO: error handling
	json.Unmarshal(inputJson, &f)

	var prettyJSON bytes.Buffer
	// TODO: error handling
	json.Indent(&prettyJSON, []byte(inputJson), "", "    ")

	buf := new(bytes.Buffer)
	utils.HighlightJson(buf, prettyJSON.String(), chromaStyle)

	return buf
}

func (b Bubble) GetInputJson() []byte {
	return b.inputJson
}

func (b *Bubble) SetSize(width, height int) {
	b.width = width
	b.height = height

	b.Styles.containerStyle.
		Width(width - b.Styles.containerStyle.GetHorizontalFrameSize()/2).
		Height(height - b.Styles.containerStyle.GetVerticalFrameSize())

	b.viewport.Width = width - b.Styles.containerStyle.GetHorizontalFrameSize() - 3
	b.viewport.Height = height - b.Styles.containerStyle.GetVerticalFrameSize() - 3

	renderedJson := lipgloss.NewStyle().Width(b.viewport.Width - 3).Render(b.highlightedJson.String())

	b.viewport.SetContent(renderedJson)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (b Bubble) View() string {
	scrollPercent := fmt.Sprintf("%3.f%%", b.viewport.ScrollPercent()*100)

	info := b.Styles.infoStyle.Render(fmt.Sprintf("%s | %s", lipgloss.NewStyle().Italic(true).Render(b.filename), scrollPercent))
	line := strings.Repeat(" ", max(0, b.viewport.Width-lipgloss.Width(info)))

	footer := lipgloss.JoinHorizontal(lipgloss.Center, line, info)
	content := lipgloss.JoinVertical(lipgloss.Left, b.viewport.View(), footer)

	return b.Styles.containerStyle.Render(content)
}

func (b Bubble) Init() tea.Cmd {
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
