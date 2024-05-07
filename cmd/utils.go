package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/noahgorstein/jqp/tui/utils"
)

// isValidInput checks the validity of input data as JSON or JSON lines.
// It takes a byte slice 'data' and returns two boolean values indicating
// whether the data is valid JSON and valid JSON lines, along with an error
// if the data is not valid in either format.
func isValidInput(data []byte) (bool, bool, error) {
	isValidJSON := utils.IsValidJSON(data)
	isValidJSONLines := utils.IsValidJSONLines(data)
	if !isValidJSON && !isValidJSONLines {
		return false, false, errors.New("Data is not valid JSON or LDJSON")
	}
	return isValidJSON, isValidJSONLines, nil
}

func streamToBytes(stream io.Reader) ([]byte, error) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(stream)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func isInputFromPipe() bool {
	fi, _ := os.Stdin.Stat()
	return (fi.Mode() & os.ModeCharDevice) == 0
}

func getFile() (*os.File, error) {
	if flags.filepath == "" {
		return nil, errors.New("please provide an input file")
	}
	if !fileExists(flags.filepath) {
		return nil, errors.New("the file provided does not exist")
	}
	file, e := os.Open(flags.filepath)
	if e != nil {
		return nil, fmt.Errorf("Unable to open file: %w", e)
	}
	return file, nil
}

func fileExists(filepath string) bool {
	info, e := os.Stat(filepath)
	if os.IsNotExist(e) {
		return false
	}
	return !info.IsDir()
}
