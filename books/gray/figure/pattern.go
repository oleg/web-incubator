package figure

import (
	"gray/multid"
	"gray/oned"
	"math"
)

type StripePattern struct {
	A, B      oned.Color
	Transform multid.Matrix4
}

func MakeStripePattern(A, B oned.Color) StripePattern {
	return StripePattern{A, B, multid.IdentityMatrix}
}
func MakeStripePatternT(A, B oned.Color, transform multid.Matrix4) StripePattern {
	return StripePattern{A, B, transform}
}

func (p StripePattern) StripeAt(point oned.Point) oned.Color {
	if math.Mod(math.Floor(point.X), 2) == 0 {
		return p.A
	}
	return p.B
}

func (p StripePattern) StripeAtObject(object Shape, worldPoint oned.Point) oned.Color {
	objectPoint := object.Transform().Inverse().MultiplyPoint(worldPoint)
	patternPoint := p.Transform.Inverse().MultiplyPoint(objectPoint)
	return p.StripeAt(patternPoint)
}
