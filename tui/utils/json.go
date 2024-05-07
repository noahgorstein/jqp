package utils

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/formatters"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
)

const FourSpaces = "    "

func highlightJSON(w io.Writer, source string, style *chroma.Style) error {
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

func IsValidJSON(input []byte) bool {
	var js json.RawMessage
	return json.Unmarshal(input, &js) == nil
}

func IsValidJSONLines(input []byte) bool {
	if len(input) == 0 {
		return false
	}

	reader := bytes.NewReader(input)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		if !IsValidJSON(scanner.Bytes()) {
			return false
		}
	}
	return true
}

func prettifyJSON(input []byte, chromaStyle *chroma.Style) (*bytes.Buffer, error) {
	var buf bytes.Buffer
	if IsValidJSON(input) {
		err := json.Indent(&buf, []byte(input), "", FourSpaces)
		if err != nil {
			return nil, err
		}
	}
	if buf.Len() == 0 {
		err := highlightJSON(&buf, string(input), chromaStyle)
		if err != nil {
			return nil, err
		}
		return &buf, nil
	}
	var highlightedBuf bytes.Buffer
	err := highlightJSON(&highlightedBuf, buf.String(), chromaStyle)
	if err != nil {
		return nil, err
	}
	return &highlightedBuf, nil
}

func Prettify(inputJSON []byte, chromaStyle *chroma.Style, isJSONLines bool) (*bytes.Buffer, error) {
	if !isJSONLines {
		return prettifyJSON(inputJSON, chromaStyle)
	}
	var buf bytes.Buffer
	reader := bytes.NewReader(inputJSON)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Bytes()
		hightlighedLine, err := prettifyJSON(line, chromaStyle)
		if err != nil {
			return nil, err
		}
		_, err = buf.WriteString(fmt.Sprintf("%v\n", hightlighedLine))
		if err != nil {
			return nil, err
		}
	}
	return &buf, nil
}
