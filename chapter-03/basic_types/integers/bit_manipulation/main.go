package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func main() {
	n := newPatternOp()
	fmt.Printf("Here's a random bit pattern: %08b\n", n)

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("")
		fmt.Println("[0] Set bit to 1.")
		fmt.Println("[1] Clear bit.")
		fmt.Println("[2] Toggle bit.")
		fmt.Println("[3] Check bit.")
		fmt.Println("[4] New bit pattern.")
		fmt.Println("[5] Bye.")
		fmt.Println("")
		fmt.Print("> ")

		input := 0
		_, err := fmt.Scanf("%d", &input)

		if err != nil {
			if err == io.EOF {
				fmt.Println("EOF")
				os.Exit(0)
			} else {
				fmt.Fprintf(os.Stderr, "Error parsing input: %v\n", err.Error())
				clear()
				continue
			}
		}
		if input > 5 || input < 0 {
			fmt.Fprintf(os.Stderr, "Invalid input, value between %d and %d expected.\n", 0, 5)
			continue
		}
		n = invokeOp(int(input), n)
	}
}

func invokeOp(index int, n byte) byte {
	var result byte
	if index == 0 {
		result = setBitOp(n)
	} else if index == 1 {
		result = clearBitOp(n)
	} else if index == 2 {
		result = toggleBitOp(n)
	} else if index == 3 {
		result = checkBitOp(n)
	} else if index == 4 {
		result = newPatternOp()
	} else {
		exitOp()
	}
	fmt.Printf("Bit pattern is: %08b\n", result)
	return result
}

func setBitOp(n byte) byte {
	pos := readBitPos()
	if pos == 100 {
		return n
	}
	return byte(n | (1 << uint64(pos)))
}

func clearBitOp(n byte) byte {
	pos := readBitPos()
	if pos == 100 {
		return n
	}
	return byte(n &^ (1 << uint64(pos)))
}

func toggleBitOp(n byte) byte {
	pos := readBitPos()
	if pos == 100 {
		return n
	}
	return byte(n ^ (1 << uint64(pos)))
}

func checkBitOp(n byte) byte {
	shift := readBitPos()
	if shift != 100 {
		result := ((n >> shift) & 1)
		fmt.Printf("Bit at position %d is: %d\n", shift, result)
	}
	return n
}

func newPatternOp() byte {
	rand.Seed(time.Now().Unix())
	return byte(rand.Intn(256))
}

func exitOp() {
	fmt.Println("Bye.")
	os.Exit(0)
}

func readBitPos() uint64 {
	var bitPos uint64
	for {
		fmt.Print("Which bit should be chosen? [0-7] ")
		_, err := fmt.Scanf("%v", &bitPos)
		if err != nil {
			if err == io.EOF {
				return uint64(100)
			}
			fmt.Fprintf(os.Stderr, "Error parsing input: %v\n", err.Error())
			continue
		}
		if bitPos < 0 || bitPos > 7 {
			fmt.Fprintf(os.Stderr, "Index out of range: %d (should be 0-7)\n", bitPos)
			continue
		}
		break
	}
	return bitPos
}

func clear() {
	var in rune
	for {
		fmt.Scanf("%c", &in)
		if in == '\n' {
			break
		}
	}
}
