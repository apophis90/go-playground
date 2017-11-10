package main

/*
  This program computes the n-th fibonacci number, reading n from standard
  input.
*/

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "fib.go: Expect a single argument.")
		os.Exit(1)
	}

	nth, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "fib.go: Error processing input.")
		os.Exit(2)
	}

	fmt.Printf("Computing nth fibonacci number with n = %d ...\n", nth)
	fmt.Printf("Result: %d\n", nthFibonacciNum(nth))
}

func nthFibonacciNum(n int) int {
	x, y := 0, 1 // Tuple assignment
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
