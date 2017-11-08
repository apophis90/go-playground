package main

import (
	"os"

	"github.com/apophis90/lets-go/chapter-01/echo/v2/echo"
)

func main() {
	echo.Echo(os.Args[1:])
}
