package main

import (
	"fmt"

	"github.com/apophis90/lets-go/chapter-02/packages/temperature/temp"
)

func main() {
	fmt.Printf("Absolute zero: %s | %s | %s\n",
		tempconv.AbsoluteZeroC, tempconv.CtoF(tempconv.AbsoluteZeroC), tempconv.CtoK(tempconv.AbsoluteZeroC))
	fmt.Printf("Freezing point: %s | %s | %s\n",
		tempconv.FreezingC, tempconv.CtoF(tempconv.FreezingC), tempconv.CtoK(tempconv.FreezingC))
	fmt.Printf("Boiling point: %s | %s | %s\n",
		tempconv.BoilingC, tempconv.CtoF(tempconv.BoilingC), tempconv.CtoK(tempconv.BoilingC))
}
