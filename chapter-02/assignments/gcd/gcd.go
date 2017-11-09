package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "gcd.go: Expect exactly 2 arguments!")
	}
	a, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Oops, something went wrong.")
	}
	b, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Oops, something went wrong.")
	}
	fmt.Printf("Greatest common divisor of %d and %d is: %d\n", a, b, gcd(a, b))
}

func gcd(x int, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
