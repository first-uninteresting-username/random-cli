package main

import (
	"flag"
	"fmt"
)

func main() {
	action := flag.String("action", "", "Select action")
	fmt.Println(action)
}
