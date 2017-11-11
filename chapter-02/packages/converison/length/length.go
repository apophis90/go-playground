package length

import "fmt"

// Meter defines a custom type for length in m.
type Meter float64

// Foot defines a custom type for length in ft.
type Foot float64

func (m Meter) String() string {
	return fmt.Sprintf("%.2f m", m)
}

func (f Foot) String() string {
	return fmt.Sprintf("%.2f ft", f)
}

// MtoF converts a length in meters to (engl.) foot.
func MtoF(m Meter) Foot {
	return Foot(m / 0.3084)
}

// FtoM converts a length in (engl.) foot to meters.
func FtoM(f Foot) Meter {
	return Meter(f * 0.3048)
}
