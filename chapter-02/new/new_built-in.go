package main

import "fmt"

func main() {
	fmt.Println(">> Create variable with new() built-in function ...")
	p := new(int) // Creates a variable of type int, initializes it with 0 and returns its address.
	fmt.Printf("Value of variable: %d\n", *p)
	fmt.Println(">> Change variable value ...")
	*p = 42
	fmt.Printf("New value is: %d\n", *p)
	fmt.Println(">> Create another variable with built-in new() ...")
	q := new(int)
	fmt.Printf("p == q yields: %t\n", p == q)
}
