package utils

import (
	"bufio"
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
	isValidJSON = IsValidJSON(data)
	isValidJSONLines = IsValidJSONLines(data)
	if !isValidJSON && !isValidJSONLines {
		return false, false, errors.New("Data is not valid JSON or NDJSON")
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

func indentJSON(input *[]byte, output *bytes.Buffer) error {
	if IsValidJSON(*input) {
		err := json.Indent(output, []byte(*input), "", FourSpaces)
		if err != nil {
			return err
		}
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
