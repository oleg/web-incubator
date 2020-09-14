package figure

import (
	"gray/oned"
	"math"
)

type StripePattern struct {
	A oned.Color
	B oned.Color
}

func (p StripePattern) StripeAt(point oned.Point) oned.Color {
	if math.Mod(math.Floor(point.X), 2) == 0 {
		return p.A
	}
	return p.B
}
