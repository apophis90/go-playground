package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "gcd.go: Expect exactly 2 arguments!")
		os.Exit(1)
	}
	a, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "gcd.go: Error reading first argument.")
		os.Exit(2)
	}
	b, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, "gcd.go: Error reading second argument.")
		os.Exit(2)
	}
	fmt.Printf("Greatest common divisor of %d and %d is: %d\n", a, b, gcd(a, b))
}

func gcd(x int, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
