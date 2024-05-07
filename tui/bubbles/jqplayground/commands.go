package jqplayground

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/atotto/clipboard"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/itchyny/gojq"

	"github.com/noahgorstein/jqp/tui/utils"
)

type errorMsg struct {
	error error
}

type queryResultMsg struct {
	rawResults         string
	highlightedResults string
}

type writeToFileMsg struct{}

type copyQueryToClipboardMsg struct{}

type copyResultsToClipboardMsg struct{}

// executeQuery executes a query using the provided query input and input data,
// returning a command that produces a message containing the results of the query.
// It parses the query input, processes the input data according to whether it's in JSON
// lines format or not, and then iterates over the results of the query, formatting them
// and returning them as a message. If an error occurs during parsing, processing, or
// iterating over the results, an error message is returned instead.
func (b *Bubble) executeQuery(ctx context.Context) tea.Cmd {
	return func() tea.Msg {
		var results strings.Builder
		query, err := gojq.Parse(b.queryinput.GetInputValue())
		if err != nil {
			return errorMsg{error: err}
		}

		processInput := func(data []byte) error {
			var obj any
			if err := json.Unmarshal(data, &obj); err != nil {
				return err
			}

			iter := query.RunWithContext(ctx, obj)
			for {
				v, ok := iter.Next()
				if !ok {
					break
				}
				if err, ok := v.(error); ok {
					return err
				}
				r, err := gojq.Marshal(v)
				if err != nil {
					continue
				}
				results.WriteString(fmt.Sprintf("%s\n", string(r)))
			}
			return nil
		}

		if b.isJSONLines {
			scanner := bufio.NewScanner(bytes.NewReader(b.inputdata.GetInputJSON()))
			for scanner.Scan() {
				line := scanner.Bytes()
				if err := processInput(line); err != nil {
					return errorMsg{error: err}
				}
			}
		} else {
			if err := processInput(b.inputdata.GetInputJSON()); err != nil {
				return errorMsg{error: err}
			}
		}

		highlightedOutput := utils.Prettify([]byte(results.String()), b.theme.ChromaStyle, true)
		return queryResultMsg{
			rawResults:         results.String(),
			highlightedResults: highlightedOutput.String(),
		}
	}
}

func (b Bubble) saveOutput() tea.Cmd {
	if b.fileselector.GetInput() == "" {
		return b.copyOutputToClipboard()
	}
	return b.writeOutputToFile()
}

func (b Bubble) copyOutputToClipboard() tea.Cmd {
	return func() tea.Msg {
		err := clipboard.WriteAll(b.results)
		if err != nil {
			return errorMsg{
				error: err,
			}
		}
		return copyResultsToClipboardMsg{}
	}
}

func (b Bubble) writeOutputToFile() tea.Cmd {
	return func() tea.Msg {
		err := os.WriteFile(b.fileselector.GetInput(), []byte(b.results), 0o644)
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
