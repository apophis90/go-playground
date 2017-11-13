package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
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

		var input int64
		hasMore := scanner.Scan()

		if reachedEOF(scanner.Err(), hasMore) {
			fmt.Printf("\nEOF\n")
			os.Exit(0)
		}

		in := scanner.Text()
		input, err := strconv.ParseInt(in, 10, 0)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing input: %v\n", in)
			continue
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
	fmt.Printf("\nBit pattern is: %08b\n", result)
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
	var err error
	// We use a new Scanner here since the outer one can't be recovered once it
	// has received EOF.
	var scanner = bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Which bit should be chosen? [0-7] ")
		hasMore := scanner.Scan()

		if reachedEOF(scanner.Err(), hasMore) {
			fmt.Println("Abort")
			return uint64(100)
		}

		bitPos, err = strconv.ParseUint(scanner.Text(), 10, 0)

		if err != nil {
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

/*
	From the docs, we know that a scanner's Scan() method returns "false" when there
	is no more token available. If Err() returns "nil" at the same time, we can be
	sure to have reached EOF (via ctrl+d from command line).
*/
func reachedEOF(err error, hasMore bool) bool {
	return err == nil && !hasMore
}
