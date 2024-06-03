package jqplayground

import (
	"context"
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/noahgorstein/jqp/tui/bubbles/inputdata"
	"github.com/noahgorstein/jqp/tui/bubbles/state"
)

func (b Bubble) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	prevState := b.state
	b.handleMessage(msg, &cmds)
	b.updateState(prevState, &cmds)
	b.updateComponents(msg, &cmds)

	return b, tea.Batch(cmds...)
}

func totalHeight(bubbles ...interface{ View() string }) int {
	var height int
	for _, bubble := range bubbles {
		height += lipgloss.Height(bubble.View())
	}
	return height
}

func (b *Bubble) resizeBubbles() {
	b.queryinput.SetWidth(b.width)
	b.statusbar.SetSize(b.width)
	b.help.SetWidth(b.width)
	height := b.height
	if b.state == state.Save {
		b.fileselector.SetSize(b.width)
		height -= totalHeight(b.help, b.queryinput, b.statusbar, b.fileselector)
	} else {
		height -= totalHeight(b.help, b.queryinput, b.statusbar)
	}
	b.inputdata.SetSize(b.width/2, height)
	b.output.SetSize(b.width/2, height)
}

//nolint:revive // don't see a more elegant way to reduce complexity here since types can't be keys in a map
func (b *Bubble) handleMessage(msg tea.Msg, cmds *[]tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		b.handleWindowSizeMsg(msg)
	case tea.KeyMsg:
		b.handleKeyMsg(msg, cmds)
	case queryResultMsg:
		b.handleQueryResultMsg(msg, cmds)
	case writeToFileMsg:
		b.handleWriteToFileMsg(msg, cmds)
	case copyResultsToClipboardMsg:
		b.handleCopyResultsToClipboardMsg(cmds)
	case copyQueryToClipboardMsg:
		b.handleCopyQueryToClipboardMsg(cmds)
	case errorMsg:
		b.handleErrorMsg(msg, cmds)
	case InvalidInputMsg:
		b.handleInvalidInput(cmds)
	case inputdata.InputReadyMsg:
		b.state = state.Query
	}
}

func (b *Bubble) handleQueryResultMsg(msg queryResultMsg, cmds *[]tea.Cmd) {
	b.state = state.Query
	b.output.ScrollToTop()
	b.output.SetContent(msg.highlightedResults)
	b.results = msg.rawResults
	*cmds = append(*cmds, b.statusbar.NewStatusMessage("Successfully executed query.", true))
}

func (b *Bubble) handleWriteToFileMsg(_ writeToFileMsg, cmds *[]tea.Cmd) {
	b.state = state.Query
	*cmds = append(*cmds, b.statusbar.NewStatusMessage(fmt.Sprintf("Successfully wrote results to file: %s", b.fileselector.GetInput()), true))
	b.fileselector.SetInput(b.workingDirectory)
}

func (b *Bubble) handleCopyResultsToClipboardMsg(cmds *[]tea.Cmd) {
	b.state = state.Query
	*cmds = append(*cmds, b.statusbar.NewStatusMessage("Successfully copied results to system clipboard.", true))
}

func (b *Bubble) handleCopyQueryToClipboardMsg(cmds *[]tea.Cmd) {
	*cmds = append(*cmds, b.statusbar.NewStatusMessage("Successfully copied query to system clipboard.", true))
}

func (b *Bubble) handleErrorMsg(msg errorMsg, cmds *[]tea.Cmd) {
	if b.state == state.Running {
		b.state = state.Query
	}
	*cmds = append(*cmds, b.statusbar.NewStatusMessage(msg.error.Error(), false))
}

func (b *Bubble) handleWindowSizeMsg(msg tea.WindowSizeMsg) {
	b.width = msg.Width
	b.height = msg.Height
	b.resizeBubbles()
}

func (b *Bubble) handleKeyMsg(msg tea.KeyMsg, cmds *[]tea.Cmd) {
	keyHandlers := map[tea.KeyType]func(){
		tea.KeyCtrlC:    func() { b.handleCtrlC(cmds) },
		tea.KeyTab:      b.handleTab,
		tea.KeyShiftTab: b.handleShiftTab,
		tea.KeyEsc:      b.handleEsc,
		tea.KeyEnter:    func() { b.handleEnter(cmds) },
		tea.KeyCtrlS:    b.handleCtrlS,
		tea.KeyCtrlY:    func() { b.handleCtrlY(cmds) },
	}
	if handler, ok := keyHandlers[msg.Type]; ok {
		handler()
	}
}

func (b *Bubble) handleInvalidInput(cmds *[]tea.Cmd) {
	b.ExitMessage = "Data is not valid JSON or NDJSON"
	*cmds = append(*cmds, tea.Quit)
}

func (b *Bubble) handleCtrlC(cmds *[]tea.Cmd) {
	if b.state != state.Running {
		*cmds = append(*cmds, tea.Quit)
	}
	if b.cancel != nil {
		b.cancel()
		b.cancel = nil
	}
	b.state = state.Query
}

func (b *Bubble) handleTab() {
	if b.state != state.Save {
		switch b.state {
		case state.Query:
			b.state = state.Input
		case state.Input:
			b.state = state.Output
		case state.Output:
			b.state = state.Query
		}
	}
}

func (b *Bubble) handleShiftTab() {
	if b.state != state.Save {
		switch b.state {
		case state.Query:
			b.state = state.Output
		case state.Input:
			b.state = state.Query
		case state.Output:
			b.state = state.Input
		}
	}
}

func (b *Bubble) handleEsc() {
	if b.state == state.Save {
		b.state = state.Query
	}
}

func (b *Bubble) executeQuery(cmds *[]tea.Cmd) {
	b.queryinput.RotateHistory()
	b.state = state.Running
	var ctx context.Context
	ctx, b.cancel = context.WithCancel(context.Background())
	*cmds = append(*cmds, b.executeQueryCommand(ctx))
}

func (b *Bubble) handleEnter(cmds *[]tea.Cmd) {
	if b.state == state.Save {
		*cmds = append(*cmds, b.saveOutput())
	}
	if b.state == state.Query {
		b.executeQuery(cmds)
	}
}

func (b *Bubble) handleCtrlS() {
	b.state = state.Save
}

func (b *Bubble) handleCtrlY(cmds *[]tea.Cmd) {
	if b.state != state.Save {
		*cmds = append(*cmds, b.copyQueryToClipboard())
	}
}

func (b *Bubble) updateState(prevState state.State, cmds *[]tea.Cmd) {
	if b.state != prevState {
		b.updateActiveComponent(cmds)
		b.help.SetState(b.state)

		// Help menu may overflow when we switch sections
		// so we need resize when active section changed.
		// We also need to resize when file selector (dis)appears.
		b.resizeBubbles()
	}
}

func (b *Bubble) updateActiveComponent(cmds *[]tea.Cmd) {
	switch b.state {
	case state.Query:
		b.setComponentBorderColors(b.theme.Primary, b.theme.Inactive, b.theme.Inactive)
		*cmds = append(*cmds, textinput.Blink)
	case state.Input:
		b.setComponentBorderColors(b.theme.Inactive, b.theme.Primary, b.theme.Inactive)
	case state.Output:
		b.setComponentBorderColors(b.theme.Inactive, b.theme.Inactive, b.theme.Primary)
	case state.Save:
		b.setComponentBorderColors(b.theme.Inactive, b.theme.Inactive, b.theme.Inactive)
	}
}

func (b *Bubble) setComponentBorderColors(query, input, output lipgloss.Color) {
	b.queryinput.SetBorderColor(query)
	b.inputdata.SetBorderColor(input)
	b.output.SetBorderColor(output)
}

func (b *Bubble) updateComponents(msg tea.Msg, cmds *[]tea.Cmd) {
	var cmd tea.Cmd
	dispatch := map[state.State]func(msg tea.Msg, cmds *[]tea.Cmd){
		state.Query: func(msg tea.Msg, cmds *[]tea.Cmd) {
			b.queryinput, cmd = b.queryinput.Update(msg)
			*cmds = append(*cmds, cmd)

		},
		state.Input: func(msg tea.Msg, cmds *[]tea.Cmd) {
			b.inputdata, cmd = b.inputdata.Update(msg)
			*cmds = append(*cmds, cmd)
		},
		state.Output: func(msg tea.Msg, cmds *[]tea.Cmd) {
			b.output, cmd = b.output.Update(msg)
			*cmds = append(*cmds, cmd)
		},
		state.Save: func(msg tea.Msg, cmds *[]tea.Cmd) {
			b.fileselector, cmd = b.fileselector.Update(msg)
			*cmds = append(*cmds, cmd)
		},
		state.Loading: func(msg tea.Msg, cmds *[]tea.Cmd) {
			b.inputdata, cmd = b.inputdata.Update(msg)
			*cmds = append(*cmds, cmd)
		},
	}

	if updateFunc, ok := dispatch[b.state]; ok {
		updateFunc(msg, cmds)
	}

	b.statusbar, cmd = b.statusbar.Update(msg)
	*cmds = append(*cmds, cmd)

	b.help, cmd = b.help.Update(msg)
	*cmds = append(*cmds, cmd)

}
