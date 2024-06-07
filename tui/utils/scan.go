package utils

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
)

// ScanLinesWithDynamicBufferSize scans the input byte slice line by line, using a dynamically
// increasing buffer size. It starts with an initial buffer size of 64KB and doubles the buffer
// size each time a line exceeds the current buffer size, up to the specified maximum buffer size.
//
// If a line exceeds the maximum buffer size, it returns an error.
//
// The processLine function is called for each line and should return an error if processing fails.
//
// The function returns an error if the input exceeds the maximum buffer size or if any other
// error occurs during line processing. It returns nil if all lines are processed successfully.
func ScanLinesWithDynamicBufferSize(input []byte, maxBufferSize int, processLine func([]byte) error) error {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	initialBufferSize := 64 * 1024 // 64KB initial buffer size

	for bufferSize := initialBufferSize; bufferSize <= maxBufferSize; bufferSize *= 2 {
		if err := scanWithBufferSize(scanner, bufferSize, maxBufferSize, processLine); err != nil {
			if errors.Is(err, bufio.ErrTooLong) {
				// Buffer size is too small, retry with a larger buffer
				continue
			}
			return err
		}
		// All lines are processed successfully
		return nil
	}

	// Input exceeds maximum buffer size
	return fmt.Errorf("input exceeds maximum buffer size of %d bytes", maxBufferSize)
}

func scanWithBufferSize(scanner *bufio.Scanner, bufferSize, maxBufferSize int, processLine func([]byte) error) error {
	scanner.Buffer(make([]byte, bufferSize), maxBufferSize)

	for scanner.Scan() {
		if err := processLine(scanner.Bytes()); err != nil {
			return err
		}
	}

	return scanner.Err()
}
