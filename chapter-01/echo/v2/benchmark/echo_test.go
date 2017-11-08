package test

import (
	"testing"

	"github.com/apophis90/lets-go/chapter-01/echo/v2/echo"
)

func BenchmarkEcho(b *testing.B) {
	args := []string{"Benchmarking", "Hello", "World"}
	echo.Echo(args)
}
