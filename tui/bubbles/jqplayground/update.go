package jqplayground

import (
	"fmt"
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/noahgorstein/jqp/tui/bubbles/state"
	"github.com/noahgorstein/jqp/tui/styles"
)

func (b Bubble) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		b.width = msg.Width
		b.height = msg.Height

		if b.state == state.Save {
			b.queryinput.SetWidth(msg.Width)
			b.inputdata.SetSize(
				int(float64(msg.Width)*0.5),
				msg.Height-lipgloss.Height(b.help.View())-lipgloss.Height(b.queryinput.View())-lipgloss.Height(b.statusbar.View())-lipgloss.Height(b.fileselector.View()))
			b.output.SetSize(
				int(float64(msg.Width)*0.5),
				msg.Height-lipgloss.Height(b.help.View())-lipgloss.Height(b.queryinput.View())-lipgloss.Height(b.statusbar.View())-lipgloss.Height(b.fileselector.View()))
			b.statusbar.SetSize(msg.Width)
			b.help.SetWidth(msg.Width)
			b.fileselector.SetSize(msg.Width)
		} else {
			b.queryinput.SetWidth(msg.Width)
			b.inputdata.SetSize(
				int(float64(msg.Width)*0.5),
				msg.Height-lipgloss.Height(b.help.View())-lipgloss.Height(b.queryinput.View())-lipgloss.Height(b.statusbar.View()))
			b.output.SetSize(
				int(float64(msg.Width)*0.5),
				msg.Height-lipgloss.Height(b.help.View())-lipgloss.Height(b.queryinput.View())-lipgloss.Height(b.statusbar.View()))
			b.statusbar.SetSize(msg.Width)
			b.help.SetWidth(msg.Width)
			b.fileselector.SetSize(msg.Width)
		}
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
				}
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
				case state.Output:
					b.state = state.Input
					b.help.SetState(state.Input)
					b.queryinput.SetBorderColor(styles.GREY)
					b.inputdata.SetBorderColor(styles.BLUE)
					b.output.SetBorderColor(styles.GREY)
				}
			}
		case tea.KeyEsc.String():
			if b.state == state.Save {
				b.state = state.Query
				b.help.SetState(state.Query)
				b.queryinput.SetBorderColor(styles.BLUE)

				b.queryinput.SetWidth(b.width)
				b.inputdata.SetSize(
					int(float64(b.width)*0.5),
					b.height-lipgloss.Height(b.help.View())-lipgloss.Height(b.queryinput.View())-lipgloss.Height(b.statusbar.View()))
				b.output.SetSize(
					int(float64(b.width)*0.5),
					b.height-lipgloss.Height(b.help.View())-lipgloss.Height(b.queryinput.View())-lipgloss.Height(b.statusbar.View()))
				b.statusbar.SetSize(b.width)
				b.help.SetWidth(b.width)
				b.fileselector.SetSize(b.width)
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
			b.inputdata.SetSize(
				int(float64(b.width)*0.5),
				b.height-lipgloss.Height(b.help.View())-lipgloss.Height(b.queryinput.View())-lipgloss.Height(b.statusbar.View())-lipgloss.Height(b.fileselector.View()))
			b.output.SetSize(
				int(float64(b.width)*0.5),
				b.height-lipgloss.Height(b.help.View())-lipgloss.Height(b.queryinput.View())-lipgloss.Height(b.statusbar.View())-lipgloss.Height(b.fileselector.View()))
			b.state = state.Save
			b.help.SetState(state.Save)
			b.setAllBordersInactive()
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
