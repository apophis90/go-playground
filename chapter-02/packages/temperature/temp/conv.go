package tempconv

// CtoF converts Celsius temperature to Fahrenheit.
func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit(c*(9/5) + 32)
}

// FtoC converts Fahrenheit temperature to Celsius.
func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * (5 / 9))
}

// CtoK converts Celsius temperature to Kelvin.
func CtoK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}
