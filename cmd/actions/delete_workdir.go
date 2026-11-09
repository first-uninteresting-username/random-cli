package actions

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/first-uninteresting-username/random-cli/cmd/registry"
)

func init() {
	err := registry.Register(
		registry.Action{
			Name:          "delete-workdir",
			Description:   "Delete the directory where the program was run from",
			IsDestructive: true,
			Run:           deleteWorkdir,
		},
	)
	if err != nil {
		panic(err)
	}
}

func deleteWorkdir(out io.Writer) error {
	workdir, err := os.Getwd()
	if err != nil {
		return err
	}
	if err := os.RemoveAll(workdir); err != nil {
		log.Printf("warning: failed to remove %s: %v", workdir, err)
	}

	if _, err := fmt.Fprintln(out, "Current directory removed!"); err != nil {
		return err
	}
	return nil
}
