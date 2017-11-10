package main

/*
  This program demonstrates the concept of type declarations in Go, by
  introducing custom types for the temperature units Celsius and Fahrenheit.
  Note that, although both types share the same underlying type float64, Celsius
  and Fahrenheit are not of the same Go type!
*/

import "fmt"

type Celsius float64 // Declare type for temp in °C

func (c Celsius) String() string {
  return fmt.Sprintf("%g°C", c)
}

type Fahrenheit float64 // Declare type for temp in °F

const (
  AbsoluteZeroC Celsius = -273.15
  FreezingC Celsius = 0
  BoilingC  Celsius =  100
)

// Translate from °C to °F
func CtoF(c Celsius) Fahrenheit {
  return Fahrenheit(c * (9 / 5) + 32)
}

// Translate from °F to °C
func FtoC(f Fahrenheit) Celsius {
  return Celsius((f - 32) * (5 / 9))
}

func main() {
  fmt.Printf("Absolute zero: %.2f°C  | %.2f°F\n", AbsoluteZeroC, CtoF(AbsoluteZeroC))
  fmt.Printf("Freezing temp: %.2f°C | %.2f°F\n", FreezingC, CtoF(FreezingC))
  fmt.Printf("Boiling temp: %.2f°C | %.2f°F\n", BoilingC, CtoF(BoilingC))
}
