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
	isJSONLines      bool
}

func New(inputJSON []byte, filename string, query string, jqtheme theme.Theme, isJSONLines bool) (Bubble, error) {
	workingDirectory, err := os.Getwd()
	if err != nil {
		return Bubble{}, err
	}

	sb := statusbar.New(jqtheme)
	sb.StatusMessageLifetime = time.Second * 10
	fs := fileselector.New(jqtheme)

	fs.SetInput(workingDirectory)

	inputData, err := inputdata.New(inputJSON, filename, jqtheme, isJSONLines)
	if err != nil {
		return Bubble{}, err
	}
	queryInput := queryinput.New(jqtheme)
	if query != "" {
		queryInput.SetQuery(query)
	}

	b := Bubble{
		workingDirectory: workingDirectory,
		state:            state.Query,
		queryinput:       queryInput,
		inputdata:        inputData,
		output:           output.New(jqtheme),
		help:             help.New(jqtheme),
		statusbar:        sb,
		fileselector:     fs,
		theme:            jqtheme,
		isJSONLines:      isJSONLines,
	}
	return b, nil
}
