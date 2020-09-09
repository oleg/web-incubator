package oned

type Ray struct {
	Origin    Point
	Direction Vector
}

func (ray Ray) Position(distance float64) Point {
	return ray.Origin.addVector(ray.Direction.multiplyScalar(distance))
}
