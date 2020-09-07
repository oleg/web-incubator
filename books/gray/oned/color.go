package oned

type Color Tuple

func NewColor(R, G, B float64) Color {
	return Color{R, G, B}
}

//todo: use link
func (t Color) R() float64 {
	return t.X
}

func (t Color) G() float64 {
	return t.Y
}

func (t Color) B() float64 {
	return t.Z
}

func (t Color) add(o Color) Color {
	return Color(Tuple(t).add(Tuple(o)))
}

func (t Color) subtract(o Color) Color {
	return Color(Tuple(t).subtract(Tuple(o)))
}

func (t Color) multiplyByScalar(scalar float64) Color {
	return Color(Tuple(t).multiplyScalar(scalar))
}

func (t Color) multiply(o Color) Color {
	return Color(Tuple(t).hadamard(Tuple(o)))
}
