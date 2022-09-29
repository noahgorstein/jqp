package jqplayground

import (
	"bytes"
	"encoding/json"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/quick"
	"github.com/noahgorstein/jqp/tui/utils"
)

func isValidJson(input []byte) bool {
	var js json.RawMessage
	return json.Unmarshal(input, &js) == nil
}

func highlightJson(input []byte, chromaStyle chroma.Style) *bytes.Buffer {

	if isValidJson(input) {
		var f interface{}
		json.Unmarshal(input, &f)
		var prettyJSON bytes.Buffer
		json.Indent(&prettyJSON, []byte(input), "", "    ")
		buf := new(bytes.Buffer)
		quick.Highlight(buf, prettyJSON.String(), "json", utils.GetTerminalColorSupport(), chromaStyle.Name)
		return buf
	}
	buf := new(bytes.Buffer)
	quick.Highlight(buf, string(input), "json", utils.GetTerminalColorSupport(), chromaStyle.Name)
	return buf
}
