package main

import (
	"flag"
	"fmt"
	"os"

	_ "github.com/first-uninteresting-username/random-cli/cmd/actions"
	"github.com/first-uninteresting-username/random-cli/cmd/registry"
)

func main() {
	actionName := flag.String("action", "", "Select action")
	list := flag.Bool("list", false, "List all action")
	flag.Parse()

	if *list == true {
		err := registry.ListActions()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	action, err := registry.GetAction(*actionName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	action.Run(os.Stdout)

	os.Exit(0)
}
