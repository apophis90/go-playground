package echo

import (
	"fmt"
)

func Echo(args []string) {
	output, sep := "", ""
	for _, arg := range args {
		output += sep + arg
		sep = " "
	}
	fmt.Println(output)
}
