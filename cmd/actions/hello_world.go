package actions

import (
	"fmt"
	"io"

	"github.com/first-uninteresting-username/random-cli/cmd/registry"
)

func init() {
	err := registry.Register(
		registry.Action{
			Name:        "hello-world",
			Description: "Display the familiar greeting",
			Run:         helloWorld,
		},
	)
	if err != nil {
		panic(err)
	}
}

func helloWorld(out io.Writer) error {
	if _, err := fmt.Fprintln(out, "hello, world"); err != nil {
		return err
	}
	return nil
}
