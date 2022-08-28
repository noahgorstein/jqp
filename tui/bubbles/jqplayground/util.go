package jqplayground

import (
	"bytes"
	"encoding/json"

	"github.com/alecthomas/chroma/quick"
	"github.com/noahgorstein/jqp/tui/styles"
	"github.com/noahgorstein/jqp/tui/utils"
)

func isValidJson(input []byte) bool {
	var js json.RawMessage
	return json.Unmarshal(input, &js) == nil
}

func highlightJson(input []byte) *bytes.Buffer {

	if isValidJson(input) {
		var f interface{}
		json.Unmarshal(input, &f)
		var prettyJSON bytes.Buffer
		json.Indent(&prettyJSON, []byte(input), "", "    ")
		buf := new(bytes.Buffer)
		quick.Highlight(buf, prettyJSON.String(), "json", "terminal16", utils.GetChromaTheme())
		return buf
	}
	buf := new(bytes.Buffer)
	quick.Highlight(buf, string(input), "json", "terminal16", utils.GetChromaTheme())
	return buf
}

func (b *Bubble) setAllBordersInactive() {
	b.queryinput.SetBorderColor(styles.GREY)
	b.inputdata.SetBorderColor(styles.GREY)
	b.output.SetBorderColor(styles.GREY)
}
