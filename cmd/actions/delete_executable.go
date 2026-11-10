package actions

import (
	"fmt"
	"io"
	"os"

	"github.com/first-uninteresting-username/random-cli/cmd/registry"
)

func init() {
	err := registry.Register(
		registry.Action{
			Name:          "delete-executable",
			Description:   "Delete the executable of the program",
			IsDestructive: true,
			Run:           deleteExecutable,
		},
	)
	if err != nil {
		panic(err)
	}
}

func deleteExecutable(_ io.Reader, out io.Writer) error {
	exe, err := os.Executable()
	if err != nil {
		return err
	}
	if err := os.Remove(exe); err != nil {
		return err
	}

	if _, err := fmt.Fprintln(out, "Executable of this CLI removed!"); err != nil {
		return err
	}
	return nil
}
