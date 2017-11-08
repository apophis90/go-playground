package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, fahrenheitToCelsius(freezingF))
	fmt.Printf("%g°F = %g°C\n", boilingF, fahrenheitToCelsius(boilingF))
}

func fahrenheitToCelsius(tempFahrenheit float64) float64 {
	return (tempFahrenheit - 32) * 5 / 9
}
