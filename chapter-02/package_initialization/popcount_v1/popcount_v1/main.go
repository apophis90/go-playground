package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/apophis90/lets-go/chapter-02/package_initialization/popcount"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Expected a single argument.")
		os.Exit(1)
	}

	n, err := strconv.ParseUint(os.Args[1], 10, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error processing input: %v\n", os.Args[1])
		os.Exit(2)
	}
	fmt.Println(popcount.PopCountv3(n))
}
