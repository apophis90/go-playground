package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Running programm " + os.Args[0] + ": ")
	echo(os.Args[1:])
}

func echo(args []string) {
	for index, val := range args {
		fmt.Println(fmt.Sprintf("Arg #%d: %s", index, val))
	}
}
