package jqplayground

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/noahgorstein/jqp/tui/bubbles/state"
	"github.com/noahgorstein/jqp/tui/styles"
)

func (b *Bubble) resizeBubbles(width, height int) {
	b.queryinput.SetWidth(width)
	b.statusbar.SetSize(width)
	b.help.SetWidth(width)
	b.fileselector.SetSize(width)
	if b.state == state.Save {
		b.inputdata.SetSize(
			int(float64(width)*0.5),
			height-lipgloss.Height(b.help.View())-lipgloss.Height(b.queryinput.View())-lipgloss.Height(b.statusbar.View())-lipgloss.Height(b.fileselector.View()))
		b.output.SetSize(
			int(float64(width)*0.5),
			height-lipgloss.Height(b.help.View())-lipgloss.Height(b.queryinput.View())-lipgloss.Height(b.statusbar.View())-lipgloss.Height(b.fileselector.View()))
	} else {
		b.inputdata.SetSize(
			int(float64(width)*0.5),
			height-lipgloss.Height(b.help.View())-lipgloss.Height(b.queryinput.View())-lipgloss.Height(b.statusbar.View()))
		b.output.SetSize(
			int(float64(width)*0.5),
			height-lipgloss.Height(b.help.View())-lipgloss.Height(b.queryinput.View())-lipgloss.Height(b.statusbar.View()))
	}
}

func (b Bubble) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		b.width = msg.Width
		b.height = msg.Height
		b.resizeBubbles(msg.Width, msg.Height)
	case tea.KeyMsg:
		switch msg.String() {
		case tea.KeyCtrlC.String():
			return b, tea.Quit
		case tea.KeyTab.String():
			if b.state != state.Save {
				switch b.state {
				case state.Query:
					b.state = state.Input
					b.help.SetState(state.Input)
					b.queryinput.SetBorderColor(styles.GREY)
					b.inputdata.SetBorderColor(styles.BLUE)
					b.output.SetBorderColor(styles.GREY)
				case state.Input:
					b.state = state.Output
					b.help.SetState(state.Output)
					b.queryinput.SetBorderColor(styles.GREY)
					b.inputdata.SetBorderColor(styles.GREY)
					b.output.SetBorderColor(styles.BLUE)
				case state.Output:
					b.state = state.Query
					b.help.SetState(state.Query)
					b.queryinput.SetBorderColor(styles.BLUE)
					b.inputdata.SetBorderColor(styles.GREY)
					b.output.SetBorderColor(styles.GREY)

					cmds = append(cmds, textinput.Blink)
				}
				// help menu may overflow when we switch sections so we need resize when active section changed
				b.resizeBubbles(b.width, b.height)
			}
		case tea.KeyShiftTab.String():
			if b.state != state.Save {
				switch b.state {
				case state.Query:
					b.state = state.Output
					b.help.SetState(state.Output)
					b.queryinput.SetBorderColor(styles.GREY)
					b.inputdata.SetBorderColor(styles.GREY)
					b.output.SetBorderColor(styles.BLUE)
				case state.Input:
					b.state = state.Query
					b.help.SetState(state.Query)
					b.queryinput.SetBorderColor(styles.BLUE)
					b.inputdata.SetBorderColor(styles.GREY)
					b.output.SetBorderColor(styles.GREY)

					cmds = append(cmds, textinput.Blink)
				case state.Output:
					b.state = state.Input
					b.help.SetState(state.Input)
					b.queryinput.SetBorderColor(styles.GREY)
					b.inputdata.SetBorderColor(styles.BLUE)
					b.output.SetBorderColor(styles.GREY)
				}
				// help menu may overflow when we switch sections so we need resize when active section changed
				b.resizeBubbles(b.width, b.height)
			}
		case tea.KeyEsc.String():
			if b.state == state.Save {
				b.state = state.Query
				b.help.SetState(state.Query)
				b.queryinput.SetBorderColor(styles.BLUE)
				b.resizeBubbles(b.width, b.height)

				cmds = append(cmds, textinput.Blink)
			}
		case tea.KeyEnter.String():
			if b.state == state.Save {
				cmd = b.writeOutputToFile()
				cmds = append(cmds, cmd)
			} else {
				cmd = b.executeQuery()
				cmds = append(cmds, cmd)
			}
		case tea.KeyCtrlS.String():
			b.state = state.Save
			b.help.SetState(state.Save)
			b.resizeBubbles(b.width, b.height)
			b.setAllBordersInactive()
		case tea.KeyCtrlY.String():
			if b.state != state.Save {
				cmd = b.copyQueryToClipboard()
				cmds = append(cmds, cmd)
			}
		}
	case queryResultMsg:
		b.output.ScrollToTop()
		b.output.SetContent(msg.highlightedResults)
		b.results = msg.rawResults
		cmd = b.statusbar.NewStatusMessage("Successfully executed query.", true)
		cmds = append(cmds, cmd)
	case writeToFileMsg:
		b.state = state.Query
		b.help.SetState(state.Query)
		b.queryinput.SetBorderColor(styles.BLUE)

		cmd = b.statusbar.NewStatusMessage(fmt.Sprintf("Successfully wrote results to file: %s", b.fileselector.GetInput()), true)
		cmds = append(cmds, cmd)

		b.fileselector.SetInput(b.workingDirectory)
		b.inputdata.SetSize(
			int(float64(b.width)*0.5),
			b.height-lipgloss.Height(b.help.View())-lipgloss.Height(b.queryinput.View())-lipgloss.Height(b.statusbar.View()))
		b.output.SetSize(
			int(float64(b.width)*0.5),
			b.height-lipgloss.Height(b.help.View())-lipgloss.Height(b.queryinput.View())-lipgloss.Height(b.statusbar.View()))

		cmds = append(cmds, textinput.Blink)

	case copyQueryToClipboardMsg:
		cmd = b.statusbar.NewStatusMessage("Successfully copied query to system clipboard.", true)
		cmds = append(cmds, cmd)
	case errorMsg:
		cmd = b.statusbar.NewStatusMessage(msg.error.Error(), false)
		cmds = append(cmds, cmd)

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
