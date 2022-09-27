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
}

func New(inputJson []byte, filename string) Bubble {

	workingDirectory, _ := os.Getwd()

	sb := statusbar.New()
	sb.StatusMessageLifetime = time.Second * 10
	fs := fileselector.New()

	fs.SetInput(workingDirectory)

	b := Bubble{
		workingDirectory: workingDirectory,
		state:            state.Query,
		queryinput:       queryinput.New(),
		inputdata:        inputdata.New(inputJson, filename),
		output:           output.New(),
		help:             help.New(),
		statusbar:        sb,
		fileselector:     fs,
	}
	return b
}
