package main

import (
	"flag"
	"fmt"
)

func main() {
	action := flag.String("action", "", "Select action")
	flag.Parse()
	fmt.Println(*action)
}
