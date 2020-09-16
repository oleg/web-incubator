package figure

import (
	"gray/multid"
	"gray/oned"
	"math"
)

type Pattern interface {
	PatternAt(point oned.Point) oned.Color
	Transform() multid.Matrix4
}

func PatternAtShape(pattern Pattern, shape Shape, worldPoint oned.Point) oned.Color {
	objectPoint := shape.Transform().Inverse().MultiplyPoint(worldPoint)
	patternPoint := pattern.Transform().Inverse().MultiplyPoint(objectPoint)
	return pattern.PatternAt(patternPoint)
}

type StripePattern struct {
	A, B      oned.Color
	transform multid.Matrix4
}

func MakeStripePattern(A, B oned.Color) StripePattern {
	return StripePattern{A, B, multid.IdentityMatrix}
}
func MakeStripePatternT(A, B oned.Color, transform multid.Matrix4) StripePattern {
	return StripePattern{A, B, transform}
}

func (p StripePattern) PatternAt(point oned.Point) oned.Color {
	if math.Mod(math.Floor(point.X), 2) == 0 {
		return p.A
	}
	return p.B
}
func (p StripePattern) Transform() multid.Matrix4 {
	return p.transform
}

type GradientPattern struct {
	A, B      oned.Color
	transform multid.Matrix4
}

func MakeGradientPattern(a, b oned.Color) GradientPattern {
	return GradientPattern{a, b, multid.IdentityMatrix}
}
func MakeGradientPatternT(a, b oned.Color, transform multid.Matrix4) GradientPattern {
	return GradientPattern{a, b, transform}
}

func (p GradientPattern) PatternAt(point oned.Point) oned.Color {
	distance := p.B.Subtract(p.A)
	fraction := point.X - math.Floor(point.X)
	return p.A.Add(distance.MultiplyByScalar(fraction))
}

func (p GradientPattern) Transform() multid.Matrix4 {
	return p.transform
}
