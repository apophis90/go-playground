package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline") // Returns a pointer to the flag value.
var sep = flag.String("sep", " ", "separator")

func main() {
	flag.Parse()                               // Update flags from default values
	fmt.Print(strings.Join(flag.Args(), *sep)) // Access programm args via flag.Args()
	if !*n {
		fmt.Println()
	}
}
