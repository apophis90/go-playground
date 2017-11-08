package test

import (
	"testing"

	"github.com/apophis90/lets-go/chapter-01/echo/v1/echo"
)

// BenchmarkEcho defines a benchmark test for the echo function.
func BenchmarkEcho(b *testing.B) {
	args := []string{"Testing", "Hello", "World"}
	echo.Echo(args)
}
