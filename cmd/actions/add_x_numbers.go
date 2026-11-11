package actions

import (
	"bufio"
	"fmt"
	"io"

	"github.com/first-uninteresting-username/random-cli/cmd/helpers"
	"github.com/first-uninteresting-username/random-cli/cmd/registry"
)

// This action is a reimplementation of a school assignment.
// The original was written in C++98 (yes, really) and was far simpler.
func init() {
	err := registry.Register(
		registry.Action{
			Name:          "add",
			Description:   "Add number of ints specified by the user",
			IsDestructive: false,
			Run:           addNums,
		},
	)
	if err != nil {
		panic(err)
	}
}

const addNumsPrompt string = "Enter number of numbers to be added: "

func addNums(in io.Reader, out io.Writer) error {
	reader := bufio.NewReader(in)

	n, err := helpers.InputInt(reader, out, addNumsPrompt)
	if err != nil {
		return err
	}

	if n < 0 {
		n = -n
	}

	var sum int

	for i := range n {
		numCount := i + 1
		var suffix string
		switch {
		case numCount%10 == 1 && numCount%100 != 11:
			suffix = "st"
		case numCount%10 == 2 && numCount%100 != 12:
			suffix = "nd"
		case numCount%10 == 3 && numCount%100 != 13:
			suffix = "rd"

		default:
			suffix = "th"
		}
		s := fmt.Sprintf("Enter your %d%s number: ", numCount, suffix)
		nI, err := helpers.InputInt(reader, out, s)
		if err != nil {
			return err
		}

		sum = sum + nI
	}

	_, err = fmt.Fprintf(out, "The sum is %d\n", sum)
	return err
}
