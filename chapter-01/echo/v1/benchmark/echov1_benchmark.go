package test

import "testing"
import "main"

func BenchmarkEcho(b *testing.B) {
  Echo(["Hello", "World"])
}
