package main

import "fmt"

func main() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("x: %08b\n", x) // "00100010"
	fmt.Printf("y: %08b\n", y) // "00000110"

	fmt.Printf("x&y (bitwise AND): %08b\n", x&y)       // "00000010" (bitwise AND, intersection)
	fmt.Printf("x|y (bitwise OR): %08b\n", x|y)        // "00100110" (bitwise OR, union)
	fmt.Printf("x^y (bitwise XOR): %08b\n", x^y)       // "00100100" (bitwise XOR, symmetric difference)
	fmt.Printf("x&^y (bitwise AND NOT): %08b\n", x&^y) // "00100000" (bitwise AND NOT, difference)

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 {
			fmt.Printf("Bit %d is set in x.\n", i) // Prints indices of bits set in x ("1", "5")
		}
	}

	fmt.Printf("Multiply x by 2: %08b\n", x<<1) // "01000100" (1x multiply by 2 means left-shift pattern by one position)
	fmt.Printf("Divide x by 2: %08b\n", x>>1)   // "00010001" (1x divide by 2 means right-shift pattern by one position)
}
