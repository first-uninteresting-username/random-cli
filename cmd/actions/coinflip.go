package actions

import (
	"fmt"
	"io"
	"math/rand/v2"

	"github.com/first-uninteresting-username/random-cli/cmd/registry"
)

func init() {
	err := registry.Register(
		registry.Action{
			Name:          "coinflip",
			Description:   "Flip a coin and print heads or tails",
			IsDestructive: false,
			Run:           flipCoin,
		},
	)
	if err != nil {
		panic(err)
	}
}

func flipCoin(_ io.Reader, out io.Writer) error {
	n := rand.Int()
	result := "tails"
	if n%2 == 0 {
		result = "heads"
	}

	_, err := fmt.Fprintln(out, result)
	return err
}
