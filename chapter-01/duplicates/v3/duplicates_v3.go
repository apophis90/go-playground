package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)

	for _, filename := range os.Args[1:] {
		// Read entire file.
		data, err := ioutil.ReadFile(filename)

		if err != nil {
			// Oops, something went wrong.
			fmt.Fprintf(os.Stderr, "duplicates_v3.go: %v\n", err)
			continue
		}

		// Convert bytes to string and split by "\n" character.
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, count := range counts {
		if count > 1 {
			fmt.Printf("Found duplicate: %d\t%q\n", count, line)
		}
	}
}
