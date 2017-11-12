package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var scanner = bufio.NewScanner(os.Stdin)

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

		scanner.Scan()
		input, err := strconv.ParseInt(scanner.Text(), 10, 0)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing input %v\n", scanner.Text())
			os.Exit(1)
		}
		if input > 5 || input < 0 {
			fmt.Fprintf(os.Stderr, "Invalid input, value between %d and %d expected.\n", 0, 4)
			os.Exit(2)
		}
		n = invokeOp(int(input), n)
	}
}

func invokeOp(index int, n byte) byte {
	var result byte
	if index == 0 {
		result = setBitOp(n)
		fmt.Printf("The new bit pattern is: %08b\n", result)
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
	return result
}

func setBitOp(n byte) byte {
	var bitPos uint64
	var err error
	for {
		fmt.Print("Which bit should be set? [0-7] ")
		scanner.Scan()
		bitPos, err = strconv.ParseUint(scanner.Text(), 10, 0)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing input: %v\n", scanner.Text())
			continue
		}
		if bitPos < 0 || bitPos > 7 {
			fmt.Fprintf(os.Stderr, "Index out of range: %d (should be 0-7)\n", bitPos)
			continue
		}
		break
	}
	return byte(n | (1 << bitPos))
}

func clearBitOp(n byte) byte {
	return byte(0)
}

func toggleBitOp(n byte) byte {
	return byte(0)
}

func checkBitOp(n byte) byte {
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
