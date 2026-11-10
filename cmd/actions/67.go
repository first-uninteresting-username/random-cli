package actions

import (
	"fmt"
	"io"
	"strings"

	"github.com/first-uninteresting-username/random-cli/cmd/registry"
)

func init() {
	err := registry.Register(
		registry.Action{
			Name:          "six-seven",
			Description:   "Write 67 67^2 times",
			IsDestructive: false,
			Run:           sixSeven,
		},
	)
	if err != nil {
		panic(err)
	}
}

func sixSeven(_ io.Reader, out io.Writer) error {

	// The fact those 2 variables have almost identical names is a part of the joke
	var six_seven_lines strings.Builder
	var six_seven_line string

	for range 67 {
		six_seven_line = six_seven_line + "67"
	}
	six_seven_line = six_seven_line + "\n"

	for range 67 {
		_, err := six_seven_lines.WriteString(six_seven_line)
		if err != nil {
			return err
		}
	}

	if _, err := fmt.Fprintln(out, six_seven_lines.String()); err != nil {
		return err
	}
	return nil
}
