package actions

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"testing"
)

func TestEcho(t *testing.T) {
	const prompt = "Enter a string: "

	cases := []struct {
		name     string
		input    string
		wantEcho string
		wantErr  error
	}{
		{name: "one line", input: "hello\n", wantEcho: "hello\n", wantErr: nil},
		{name: "no newline", input: "hello", wantEcho: "", wantErr: io.EOF},
		{name: "empty", input: "", wantEcho: "", wantErr: io.EOF},
		{name: "two lines", input: "a\nb\n", wantEcho: "a\n", wantErr: nil},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			buf := bytes.Buffer{}
			err := echo(strings.NewReader(c.input), &buf)

			if !errors.Is(err, c.wantErr) {
				t.Errorf("err = %v, want %v", err, c.wantErr)
			}

			want := prompt + c.wantEcho
			if got := buf.String(); got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		})
	}
}
