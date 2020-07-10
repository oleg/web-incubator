package gray

import "math"

type Vector Tuple //w = 0

func (t Vector) addVector(o Vector) Vector {
	return Vector(Tuple(t).add(Tuple(o)))
}

func (t Vector) subtractVector(o Vector) Vector {
	return Vector(Tuple(t).subtract(Tuple(o)))
}

func (t Vector) negate() Vector {
	return Vector(Tuple(t).negate())
}

func (t Vector) magnitude() float64 { //todo: move to Tuple?
	return math.Sqrt(t.x*t.x + t.y*t.y + t.z*t.z)
}

func (t Vector) normalize() Vector { //todo: move to Tuple?
	magnitude := t.magnitude()
	return Vector{t.x / magnitude, t.y / magnitude, t.z / magnitude}
}

func (t Vector) dot(o Vector) float64 { //todo: move to Tuple?
	return t.x*o.x + t.y*o.y + t.z*o.z
}

func (t Vector) cross(o Vector) Vector { //todo: move to tuple?
	return Vector{
		t.y*o.z - t.z*o.y,
		t.z*o.x - t.x*o.z,
		t.x*o.y - t.y*o.x}
}
