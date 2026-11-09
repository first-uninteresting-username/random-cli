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

func flipCoin(out io.Writer) error {
	n := rand.Int()
	result := "tails"
	if n%2 == 0 {
		result = "heads"
	}

	if _, err := fmt.Fprintln(out, result); err != nil {
		return err
	}
	return nil
}
