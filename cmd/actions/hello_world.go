package actions

import (
	"fmt"
	"io"

	"github.com/first-uninteresting-username/random-cli/cmd/registry"
)

func init() {
	err := registry.Register(
		registry.Action{
			Name:          "hello-world",
			Description:   "Display the familiar greeting",
			IsDestructive: false,
			Run:           helloWorld,
		},
	)
	if err != nil {
		panic(err)
	}
}

func helloWorld(_ io.Reader, out io.Writer) error {
	_, err := fmt.Fprintln(out, "hello, world")
	return err
}
