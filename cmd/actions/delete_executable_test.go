package actions

import (
	"bytes"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

func TestDeleteExecutable(t *testing.T) {
	buf := bytes.Buffer{}
	dir := t.TempDir()

	filePath := filepath.Join(dir, "file.txt")
	if err := os.WriteFile(filePath, []byte("hello\n"), 0o644); err != nil {
		t.Fatalf("could not create file in %q", dir)
	}

	err := deleteAndReport(&buf, filePath)
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	want := "Executable of this program removed!\n"
	got := buf.String()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	_, err = os.Stat(filePath)
	if !errors.Is(err, fs.ErrNotExist) {
		t.Errorf("want 'fs.ErrNotExist' got %v", err)
	}
}
