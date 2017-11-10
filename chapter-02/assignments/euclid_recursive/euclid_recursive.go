package main

/*
  This is the classical Euclidean algorithm for computing the greatest common
  divisor of two non-negative natural numbers a and b.
  What it does is subtracting the smaller number from the greater number until
  both numbers become equal, which is the greatest common divisor that has been
  searched.
*/

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "euklid_old.go: Expect exactly 2 arguments!")
		os.Exit(1)
	}
	a, err := strconv.Atoi(os.Args[1])
	if err != nil {
		os.Exit(2)
		fmt.Fprintln(os.Stderr, "euklid_old.go: Error reading first argument.")
	}
	b, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, "euklid_old.go: Error reading second argument.")
		os.Exit(2)
	}
	if a < 0 || b < 0 {
		fmt.Fprintln(os.Stderr, "euklid_old.go: Algorithm defined only for positive numbers.")
		os.Exit(3)
	}
	fmt.Printf("Greatest common divisor of %d and %d is: %d\n", a, b, gcd(a, b))
}

func gcd(a int, b int) int {
	if a == 0 {
		return b
	}

	if b != 0 {
		if a > b {
			return gcd(a-b, b)
		}
		return gcd(a, b-a)
	}
	return a
}
