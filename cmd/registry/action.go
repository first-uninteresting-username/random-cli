package registry

import (
	"io"
	"slices"
	"strings"
)

var actions []Action

type Action struct {
	Name        string
	Description string
	Run         func(out io.Writer) error
}

func sortActions() []Action {
	sorted := slices.Clone(actions)
	slices.SortFunc(sorted, func(a, b Action) int {
		return strings.Compare(a.Name, b.Name)
	})
	return sorted
}
