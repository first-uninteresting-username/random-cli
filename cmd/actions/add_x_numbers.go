package actions

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

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

func addNums(in io.Reader, out io.Writer) error {
	fmt.Fprint(out, "Enter number of numbers to be added: ")
	reader := bufio.NewReader(in)
	num, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	n, err := strconv.Atoi(strings.TrimSpace(num))
	if err != nil {
		return err
	}

	var sum int

	for i := 0; i < n; i++ {
		fmt.Fprintf(out, "Enter your %dth number: ", i+1)

		numI, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		nI, err := strconv.Atoi(strings.TrimSpace(numI))
		if err != nil {
			return err
		}

		sum = sum + nI
	}

	_, err = fmt.Fprintf(out, "The sum is %d\n", sum)
	return err
}
