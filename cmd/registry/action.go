package registry

import "io"

type Action struct {
	Name        string
	Description string
	Run         func(out io.Writer) error
}
