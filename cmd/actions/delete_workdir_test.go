package actions

import (
	"bytes"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

func TestDeleteWorkdir(t *testing.T) {
	buf := bytes.Buffer{}
	dir := t.TempDir()

	if err := os.MkdirAll(filepath.Join(dir, "subdir"), 0o755); err != nil {
		t.Fatalf("could not create subdir in %q", dir)
	}

	if err := os.WriteFile(filepath.Join(dir, "file.txt"), []byte("hello\n"), 0o644); err != nil {
		t.Fatalf("could not create file in %q", dir)
	}

	t.Chdir(dir)

	err := deleteWorkdir(nil, &buf)
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	want := deleteWorkdirOutput + "\n"
	got := buf.String()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	_, err = os.Stat(dir)
	if !errors.Is(err, fs.ErrNotExist) {
		t.Errorf("want 'fs.ErrNotExist' got %v", err)
	}
}
