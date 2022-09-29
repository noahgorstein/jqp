package jqplayground

import (
	"os"
	"time"

	"github.com/noahgorstein/jqp/tui/bubbles/fileselector"
	"github.com/noahgorstein/jqp/tui/bubbles/help"
	"github.com/noahgorstein/jqp/tui/bubbles/inputdata"
	"github.com/noahgorstein/jqp/tui/bubbles/output"
	"github.com/noahgorstein/jqp/tui/bubbles/queryinput"
	"github.com/noahgorstein/jqp/tui/bubbles/state"
	"github.com/noahgorstein/jqp/tui/bubbles/statusbar"
	"github.com/noahgorstein/jqp/tui/theme"
)

func (b Bubble) GetState() state.State {
	return b.state
}

type Bubble struct {
	width            int
	height           int
	workingDirectory string
	state            state.State
	queryinput       queryinput.Bubble
	inputdata        inputdata.Bubble
	output           output.Bubble
	help             help.Bubble
	statusbar        statusbar.Bubble
	fileselector     fileselector.Bubble
	results          string
	cancel           func()
	theme            theme.Theme
}

func New(inputJson []byte, filename string, theme theme.Theme) Bubble {

	workingDirectory, _ := os.Getwd()

	sb := statusbar.New(theme)
	sb.StatusMessageLifetime = time.Second * 10
	fs := fileselector.New(theme)

	fs.SetInput(workingDirectory)

	b := Bubble{
		workingDirectory: workingDirectory,
		state:            state.Query,
		queryinput:       queryinput.New(theme),
		inputdata:        inputdata.New(inputJson, filename, theme),
		output:           output.New(theme),
		help:             help.New(theme),
		statusbar:        sb,
		fileselector:     fs,
		theme:            theme,
	}
	return b
}
