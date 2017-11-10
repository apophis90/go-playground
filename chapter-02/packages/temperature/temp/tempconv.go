package tempconv

import "fmt"

// Celsius poses a custom type for temp in °C
type Celsius float64

// Fahrenheit poses a custom type for temp in °F
type Fahrenheit float64

// Kelvin poses a custom type for temp in K
type Kelvin float64

const (
	// AbsoluteZeroC defines absolute zero in °C
	AbsoluteZeroC Celsius = -273.15

	// FreezingC defines the freezing point in °C
	FreezingC Celsius = 0

	// BoilingC defines the boiling point in °C
	BoilingC Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%.2f°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%.2f°F", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%.2fK", k)
}
