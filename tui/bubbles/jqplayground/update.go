package jqplayground

import (
	"context"
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/noahgorstein/jqp/tui/bubbles/state"
)

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

func (b Bubble) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	prevState := b.state

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		b.width = msg.Width
		b.height = msg.Height
		b.resizeBubbles()
	case tea.KeyMsg:
		switch msg.String() {
		case tea.KeyCtrlC.String():
			if b.state != state.Running {
				return b, tea.Quit
			}
			if b.cancel != nil {
				b.cancel()
				b.cancel = nil
			}
			b.state = state.Query
		case tea.KeyTab.String():
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
		case tea.KeyShiftTab.String():
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
		case tea.KeyEsc.String():
			if b.state == state.Save {
				b.state = state.Query
			}
		case tea.KeyEnter.String():
			if b.state == state.Save {
				cmd = b.writeOutputToFile()
				cmds = append(cmds, cmd)
			} else if b.state == state.Query {
				b.queryinput.RotateHistory()
				b.state = state.Running
				var ctx context.Context
				ctx, b.cancel = context.WithCancel(context.Background())
				cmd = b.executeQuery(ctx)
				cmds = append(cmds, cmd)
			}
		case tea.KeyCtrlS.String():
			b.state = state.Save
		case tea.KeyCtrlY.String():
			if b.state != state.Save {
				cmd = b.copyQueryToClipboard()
				cmds = append(cmds, cmd)
			}
		}
	case queryResultMsg:
		b.state = state.Query
		b.output.ScrollToTop()
		b.output.SetContent(msg.highlightedResults)
		b.results = msg.rawResults
		cmd = b.statusbar.NewStatusMessage("Successfully executed query.", true)
		cmds = append(cmds, cmd)
	case writeToFileMsg:
		b.state = state.Query
		cmd = b.statusbar.NewStatusMessage(fmt.Sprintf("Successfully wrote results to file: %s", b.fileselector.GetInput()), true)
		cmds = append(cmds, cmd)
		b.fileselector.SetInput(b.workingDirectory)
	case copyQueryToClipboardMsg:
		cmd = b.statusbar.NewStatusMessage("Successfully copied query to system clipboard.", true)
		cmds = append(cmds, cmd)
	case errorMsg:
		if b.state == state.Running {
			b.state = state.Query
		}
		cmd = b.statusbar.NewStatusMessage(msg.error.Error(), false)
		cmds = append(cmds, cmd)

	}

	if b.state != prevState {
		switch b.state {
		case state.Query:
			b.queryinput.SetBorderColor(b.theme.Primary)
			b.inputdata.SetBorderColor(b.theme.Inactive)
			b.output.SetBorderColor(b.theme.Inactive)
			cmds = append(cmds, textinput.Blink)
		case state.Input:
			b.queryinput.SetBorderColor(b.theme.Inactive)
			b.inputdata.SetBorderColor(b.theme.Primary)
			b.output.SetBorderColor(b.theme.Inactive)
		case state.Output:
			b.queryinput.SetBorderColor(b.theme.Inactive)
			b.inputdata.SetBorderColor(b.theme.Inactive)
			b.output.SetBorderColor(b.theme.Primary)
		case state.Save:
			b.queryinput.SetBorderColor(b.theme.Inactive)
			b.inputdata.SetBorderColor(b.theme.Inactive)
			b.output.SetBorderColor(b.theme.Inactive)
		}
		b.help.SetState(b.state)
		// Help menu may overflow when we switch sections
		// so we need resize when active section changed.
		// We also need to resize when file selector (dis)appears.
		b.resizeBubbles()
	}

	switch b.state {
	case state.Query:
		b.queryinput, cmd = b.queryinput.Update(msg)
		cmds = append(cmds, cmd)
	case state.Input:
		b.inputdata, cmd = b.inputdata.Update(msg)
		cmds = append(cmds, cmd)
	case state.Output:
		b.output, cmd = b.output.Update(msg)
		cmds = append(cmds, cmd)
	case state.Save:
		b.fileselector, cmd = b.fileselector.Update(msg)
		cmds = append(cmds, cmd)
	}

	b.statusbar, cmd = b.statusbar.Update(msg)
	cmds = append(cmds, cmd)

	b.help, cmd = b.help.Update(msg)
	cmds = append(cmds, cmd)

	return b, tea.Batch(cmds...)
}
