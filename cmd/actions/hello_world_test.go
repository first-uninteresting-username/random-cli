package actions

import (
	"bytes"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	buf := bytes.Buffer{}
	err := helloWorld(nil, &buf)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := "hello, world\n"
	got := buf.String()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
