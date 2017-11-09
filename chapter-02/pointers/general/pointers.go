package main

import "fmt"

func main() {
	fmt.Println(">> Testing variable updates via pointers ...")
	i := 1
	fmt.Printf("Old value of i is: %d\n", i)
	p := &i // p holds memory address of var i
	*p = 2  // * dereferences memory address stored in p, so var i gets assigned a new value.
	fmt.Printf("New value of i is: %d\n", i)

	fmt.Println(">> Testing functions returning pointers to local variables ...")
	fmt.Printf("Value of local var i is: %d\n", *f()) // *f() dereferences returned pointer.

	fmt.Println(">> Testing pointers as function arguments ...")
	x := 42
	y := inc(&x) // inc modifies value referenced by &x and returns a COPY of the result.
	fmt.Printf("Integer value held by x: %d\n", x)
	fmt.Println("Modifying variable y ...")
	y = 0
	fmt.Printf("y should have a new value: %d, x should remain unchanged: %d\n", y, x)
}

func f() *int {
	i := 42
	return &i
}

func inc(x *int) int {
	*x++      // Increment int value x points to.
	return *x // Return value (value referenced by x gets copied!)
}
