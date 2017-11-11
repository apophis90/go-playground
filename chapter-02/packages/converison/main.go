package main

import (
	"fmt"

	"github.com/apophis90/lets-go/chapter-02/packages/converison/length"
	"github.com/apophis90/lets-go/chapter-02/packages/converison/temp"
)

func main() {
	fmt.Printf("%s = %s, %s = %s\n", temp.Fahrenheit(32), temp.FtoC(temp.Fahrenheit(32)),
		temp.Celsius(32), temp.CtoF(temp.Celsius(32)))

	fmt.Printf("%s = %s, %s = %s\n", length.Meter(10), length.MtoF(length.Meter(10)),
		length.Foot(10), length.FtoM(length.Foot(10)))
}
