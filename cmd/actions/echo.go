package actions

import (
	"bufio"
	"fmt"
	"io"

	"github.com/first-uninteresting-username/random-cli/cmd/registry"
)

func init() {
	err := registry.Register(
		registry.Action{
			Name:          "echo",
			Description:   "Echo your string",
			IsDestructive: false,
			Run:           echo,
		},
	)
	if err != nil {
		panic(err)
	}
}

func echo(in io.Reader, out io.Writer) error {
	fmt.Fprint(out, "Enter a string: ")
	reader := bufio.NewReader(in)
	line, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	if _, err := fmt.Fprint(out, line); err != nil {
		return err
	}
	return nil
}
