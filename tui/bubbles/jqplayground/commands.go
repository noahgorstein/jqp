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

// processQueryResults iterates through the results of a gojq query on the provided JSON object
// and appends the formatted results to the provided string builder.
func processQueryResults(ctx context.Context, results *strings.Builder, query *gojq.Query, obj any) error {
	iter := query.RunWithContext(ctx, obj)
	for {
		v, ok := iter.Next()
		if !ok {
			break
		}

		if err, ok := v.(error); ok {
			return err
		}

		if r, err := gojq.Marshal(v); err == nil {
			results.WriteString(fmt.Sprintf("%s\n", string(r)))
		}
	}
	return nil
}

func processJSONWithQuery(ctx context.Context, results *strings.Builder, query *gojq.Query, data []byte) error {
	var obj any
	if err := json.Unmarshal(data, &obj); err != nil {
		return err
	}
	err := processQueryResults(ctx, results, query, obj)
	if err != nil {
		return err
	}

	return nil
}

func processJSONLinesWithQuery(ctx context.Context, results *strings.Builder, query *gojq.Query, data []byte) error {
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		line := scanner.Bytes()
		if err := processJSONWithQuery(ctx, results, query, line); err != nil {
			return err
		}
	}
	return nil
}

func (b *Bubble) executeQueryOnInput(ctx context.Context) (string, error) {
	var results strings.Builder
	query, err := gojq.Parse(b.queryinput.GetInputValue())
	if err != nil {
		return "", err
	}

	processor := processJSONWithQuery

	if b.isJSONLines {
		processor = processJSONLinesWithQuery
	}
	if err := processor(ctx, &results, query, b.inputdata.GetInputJSON()); err != nil {
		return "", err
	}
	return results.String(), nil
}

func (b *Bubble) executeQueryCommand(ctx context.Context) tea.Cmd {
	return func() tea.Msg {
		results, err := b.executeQueryOnInput(ctx)
		if err != nil {
			return errorMsg{error: err}
		}
		highlightedOutput, err := utils.Prettify([]byte(results), b.theme.ChromaStyle, true)
		if err != nil {
			return errorMsg{error: err}
		}
		return queryResultMsg{
			rawResults:         results,
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
		err := os.WriteFile(b.fileselector.GetInput(), []byte(b.results), 0o600)
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
