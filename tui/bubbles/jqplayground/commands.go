package jqplayground

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/atotto/clipboard"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/itchyny/gojq"
)

type successMsg struct {
	message string
}

type errorMsg struct {
	error error
}

type queryResultMsg struct {
	rawResults         string
	highlightedResults string
}

type writeToFileMsg struct{}

type copyQueryToClipboardMsg struct{}

func (b *Bubble) executeQuery(ctx context.Context) tea.Cmd {
	return func() tea.Msg {
		query, err := gojq.Parse(b.queryinput.GetInputValue())
		if err != nil {
			return errorMsg{
				error: err,
			}
		}
		var msgTemplate interface{}
		json.Unmarshal(b.inputdata.GetInputJson(), &msgTemplate)
		var results strings.Builder
		iter := query.RunWithContext(ctx, msgTemplate)
		for {
			v, ok := iter.Next()
			if !ok {
				break
			}
			if err, ok := v.(error); ok {
				return errorMsg{
					error: err,
				}
			}
			r, _ := gojq.Marshal(v)
			results.WriteString(fmt.Sprintf("%s\n", string(r)))
		}

		highlightedOutput := highlightJson([]byte(results.String()))
		return queryResultMsg{
			rawResults:         results.String(),
			highlightedResults: highlightedOutput.String(),
		}
	}
}

func (b Bubble) writeOutputToFile() tea.Cmd {
	return func() tea.Msg {
		err := os.WriteFile(b.fileselector.GetInput(), []byte(b.results), 0644)
		if err != nil {
			return errorMsg{
				error: err,
			}
		}
		return writeToFileMsg{}
	}
}

func (b Bubble) copyQueryToClipboard() tea.Cmd {
	return func() tea.Msg {
		err := clipboard.WriteAll(b.queryinput.GetInputValue())
		if err != nil {
			return errorMsg{
				error: err,
			}
		}
		return copyQueryToClipboardMsg{}
	}
}
