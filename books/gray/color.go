package gray

type Color struct {
	R, G, B float64 //todo: use standard color
}

func (t Color) add(o Color) Color {
	return Color{t.R + o.R, t.G + o.G, t.B + o.B}
}

func (t Color) subtract(o Color) Color {
	return Color{t.R - o.R, t.G - o.G, t.B - o.B}
}

func (t Color) multiplyByScalar(scalar float64) Color {
	return Color{t.R * scalar, t.G * scalar, t.B * scalar}
}

func (t Color) multiply(o Color) Color {
	return Color{t.R * o.R, t.G * o.G, t.B * o.B}
}
