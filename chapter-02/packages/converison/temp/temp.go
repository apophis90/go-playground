package temp

import "fmt"

// Celsius poses a custom type for temp in °C
type Celsius float64

// Fahrenheit poses a custom type for temp in °F
type Fahrenheit float64

func (c Celsius) String() string {
	return fmt.Sprintf("%.2f°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%.2f°F", f)
}

// CtoF converts Celsius temperature to Fahrenheit.
func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit(c*(9/5) + 32)
}

// FtoC converts Fahrenheit temperature to Celsius.
func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * (5 / 9))
}
