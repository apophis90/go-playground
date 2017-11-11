package main

/*
  This program reads a single numerical value from its command-line arguments
  or waits to be passed a value from stdin of no command-line argument is given.
  Then, it interprets the value as a temperature (Celsius/Fahrenheit) or length
  (meters/feet) and converts the value between different units.
*/

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/apophis90/lets-go/chapter-02/packages/converison/length"
	"github.com/apophis90/lets-go/chapter-02/packages/converison/temp"
)

func main() {
	var input string

	// If there are not enough CLI arguments, try reading from stdin.
	if len(os.Args) < 2 {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input = scanner.Text()
	} else {
		input = os.Args[1]
	}

	// Try to convert input received as a CLI arg or via stdin to a float value.
	val, err := strconv.ParseFloat(input, 64)
	if err != nil {
		// Print error and exit if float conversion fails.
		fmt.Fprintf(os.Stderr, "Error processing argument: %v (numerical value expected)\n", input)
		os.Exit(2)
	}

	fmt.Printf("%s = %s, %s = %s\n", temp.Fahrenheit(val), temp.FtoC(temp.Fahrenheit(val)),
		temp.Celsius(val), temp.CtoF(temp.Celsius(val)))

	fmt.Printf("%s = %s, %s = %s\n", length.Meter(val), length.MtoF(length.Meter(val)),
		length.Foot(val), length.FtoM(length.Foot(val)))
}
