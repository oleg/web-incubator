package oned

import "math"

type Vector Tuple //w = 0
//todo: use link
func (t Vector) addVector(o Vector) Vector {
	return Vector(Tuple(t).add(Tuple(o)))
}

func (t Vector) subtractVector(o Vector) Vector {
	return Vector(Tuple(t).subtract(Tuple(o)))
}

func (t Vector) negate() Vector {
	return Vector(Tuple(t).negate())
}

func (t Vector) dot(o Vector) float64 {
	return Tuple(t).dot(Tuple(o))
}

func (t Vector) cross(o Vector) Vector {
	return Vector(Tuple(t).cross(Tuple(o)))
}

func (t Vector) multiplyScalar(scalar float64) Vector {
	return Vector(Tuple(t).multiplyScalar(scalar))
}


func (t Vector) magnitude() float64 { //todo: move to Tuple?
	return math.Sqrt(t.X*t.X + t.Y*t.Y + t.Z*t.Z)
}

func (t Vector) normalize() Vector { //todo: move to Tuple?
	magnitude := t.magnitude()
	return Vector{t.X / magnitude, t.Y / magnitude, t.Z / magnitude}
}
