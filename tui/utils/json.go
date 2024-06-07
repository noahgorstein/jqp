package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/formatters"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
)

const FourSpaces = "    "

// IsValidInput checks the validity of input data as JSON or JSON lines.
// It takes a byte slice 'data' and returns two boolean values indicating
// whether the data is valid JSON and valid JSON lines, along with an error
// if the data is not valid in either format.
func IsValidInput(data []byte) (isValidJSON bool, isValidJSONLines bool, err error) {
	if len(data) == 0 {
		err = errors.New("Data is not valid JSON or NDJSON")
		return false, false, err
	}

	isValidJSON = IsValidJSON(data) == nil
	isValidJSONLines = IsValidJSONLines(data) == nil

	if !isValidJSON && !isValidJSONLines {
		err = errors.New("Data is not valid JSON or NDJSON")
		return false, false, err
	}

	return isValidJSON, isValidJSONLines, nil
}

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

func IsValidJSON(input []byte) error {
	var js json.RawMessage
	return json.Unmarshal(input, &js)
}

func IsValidJSONLines(input []byte) error {
	maxBufferSize := 100 * 1024 * 1024 // 100MB
	err := ScanLinesWithDynamicBufferSize(input, maxBufferSize, IsValidJSON)
	if err != nil {
		return err
	}
	return nil
}

func indentJSON(input *[]byte, output *bytes.Buffer) error {
	err := IsValidJSON(*input)
	if err != nil {
		return nil
	}
	err = json.Indent(output, []byte(*input), "", FourSpaces)
	if err != nil {
		return err
	}
	return nil
}

func prettifyJSON(input []byte, chromaStyle *chroma.Style) (*bytes.Buffer, error) {
	var indentedBuf bytes.Buffer
	err := indentJSON(&input, &indentedBuf)
	if err != nil {
		return nil, err
	}
	if indentedBuf.Len() == 0 {
		err := highlightJSON(&indentedBuf, string(input), chromaStyle)
		if err != nil {
			return nil, err
		}
		return &indentedBuf, nil
	}
	var highlightedBuf bytes.Buffer
	err = highlightJSON(&highlightedBuf, indentedBuf.String(), chromaStyle)
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
	processLine := func(line []byte) error {
		hightlighedLine, err := prettifyJSON(line, chromaStyle)
		if err != nil {
			return err
		}
		_, err = buf.WriteString(fmt.Sprintf("%v\n", hightlighedLine))
		if err != nil {
			return err
		}
		return nil
	}

	const maxBufferSize = 100 * 1024 * 1024 // 100MB max buffer size
	err := ScanLinesWithDynamicBufferSize(inputJSON, maxBufferSize, processLine)
	if err != nil {
		return nil, err
	}
	return &buf, nil
}
