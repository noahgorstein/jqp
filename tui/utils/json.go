package utils

import (
	"io"

	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/formatters"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
)

const FourSpaces = "    "

func highlightJson(w io.Writer, source string, style *chroma.Style) error {
	l := lexers.Get("json")
	if l == nil {
		l = lexers.Fallback
	}
	l = chroma.Coalesce(l)

	f := formatters.Get(getTerminalColorSupport())
	if f == nil {
		f = formatters.Fallback
	}

	if style == nil {
		style = styles.Fallback
	}

	it, err := l.Tokenise(nil, source)
	if err != nil {
		return err
	}
	return f.Format(w, style, it)
}

func IsValidJson(input []byte) bool {
	var js json.RawMessage
	return json.Unmarshal(input, &js) == nil
}

func IsValidJsonLines(input []byte) bool {
	reader := bytes.NewReader(input)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		if !IsValidJson(scanner.Bytes()) {
			return false
		}
	}
	return true
}

func prettifyJson(input []byte, chromaStyle *chroma.Style) *bytes.Buffer {

	var buf bytes.Buffer
	if IsValidJson(input) {
		json.Indent(&buf, []byte(input), "", FourSpaces)
	}
	if buf.Len() == 0 {
		highlightJson(&buf, string(input), chromaStyle)
		return &buf
	}
	var highlightedBuf bytes.Buffer
	highlightJson(&highlightedBuf, buf.String(), chromaStyle)
	return &highlightedBuf
}

func Prettify(inputJson []byte, chromaStyle *chroma.Style, isJsonLines bool) *bytes.Buffer {
	if isJsonLines {
		var buf bytes.Buffer
		reader := bytes.NewReader(inputJson)
		scanner := bufio.NewScanner(reader)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			line := scanner.Bytes()
			hightlighedLine := prettifyJson(line, chromaStyle)
			buf.WriteString(fmt.Sprintf("%v\n", hightlighedLine))
		}
		return &buf
	}
	return prettifyJson(inputJson, chromaStyle)
}
