package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	// Pass stdin reader to scanner.
	input := bufio.NewScanner(os.Stdin)

	// Move on to next token. Return true if there is one, false otherwise.
	for input.Scan() {
		counts[input.Text()]++
	}

	for line, count := range counts {
		if count > 1 {
			// Use %q conversion character (verb) to get a quoted string.
			fmt.Printf("Found duplicate: %d\t%q\n", count, line)
		}
	}
}
