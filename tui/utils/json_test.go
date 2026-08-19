package utils

import (
	"bytes"
	"regexp"
	"testing"
)

var ansiEscape = regexp.MustCompile(`\x1b\[[0-9;]*m`)

func TestIndentUnit(t *testing.T) {
	tests := []struct {
		name   string
		spaces int
		want   string
	}{
		{name: "two spaces", spaces: 2, want: "  "},
		{name: "four spaces", spaces: 4, want: FourSpaces},
		{name: "one space", spaces: 1, want: " "},
		{name: "zero falls back to default", spaces: 0, want: FourSpaces},
		{name: "negative falls back to default", spaces: -3, want: FourSpaces},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := indentUnit(tt.spaces); got != tt.want {
				t.Errorf("indentUnit(%d) = %q, want %q", tt.spaces, got, tt.want)
			}
		})
	}
}

func assertIndentJSON(t *testing.T, indent, want string) {
	t.Helper()

	data := []byte(`{"a":{"b":1}}`)
	var out bytes.Buffer
	if err := indentJSON(&data, &out, indent); err != nil {
		t.Fatalf("indentJSON returned error: %v", err)
	}
	if out.String() != want {
		t.Errorf("indentJSON with %q =\n%q\nwant\n%q", indent, out.String(), want)
	}
}

func TestIndentJSON(t *testing.T) {
	t.Run("two spaces", func(t *testing.T) {
		assertIndentJSON(t, "  ", "{\n  \"a\": {\n    \"b\": 1\n  }\n}")
	})
	t.Run("four spaces", func(t *testing.T) {
		assertIndentJSON(t, FourSpaces, "{\n    \"a\": {\n        \"b\": 1\n    }\n}")
	})
}

func TestIndentJSONInvalidInputIsNoop(t *testing.T) {
	var out bytes.Buffer
	data := []byte("not json")
	if err := indentJSON(&data, &out, "  "); err != nil {
		t.Fatalf("indentJSON returned error for invalid input: %v", err)
	}
	if out.Len() != 0 {
		t.Errorf("expected no output for invalid input, got %q", out.String())
	}
}

func TestPrettifyRespectsIndent(t *testing.T) {
	input := []byte(`{"a":{"b":1}}`)

	buf, err := Prettify(input, nil, false, 2)
	if err != nil {
		t.Fatalf("Prettify returned error: %v", err)
	}

	// Highlighting wraps tokens in ANSI escapes, so strip them before comparing
	// against the expected 2-space indented output.
	got := ansiEscape.ReplaceAllString(buf.String(), "")
	want := "{\n  \"a\": {\n    \"b\": 1\n  }\n}"
	if got != want {
		t.Errorf("Prettify with indent 2 =\n%q\nwant\n%q", got, want)
	}
}
