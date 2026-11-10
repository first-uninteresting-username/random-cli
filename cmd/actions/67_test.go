package actions

import (
	"bytes"
	"strings"
	"testing"
)

func TestSixSeven(t *testing.T) {
	buf := bytes.Buffer{}
	err := sixSeven(nil, &buf)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := buf.String()
	got = strings.TrimSuffix(got, "\n")
	lines := strings.Split(got, "\n")

	lineCount := len(lines)
	if lineCount != 67 {
		t.Fatalf("got %d lines, want 67", lineCount)
	}

	wantLine := strings.Repeat("67", 67)
	for n := range 67 {
		line := lines[n]
		if line != wantLine {
			t.Errorf("line %d does not contain '67' repeated 67 times", n)
		}
	}

}
