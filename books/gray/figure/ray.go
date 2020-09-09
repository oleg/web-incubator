package figure

import "gray/oned"

type Ray struct {
	Origin    oned.Point
	Direction oned.Vector
}

func (ray Ray) Position(distance float64) oned.Point {
	return ray.Origin.AddVector(ray.Direction.MultiplyScalar(distance))
}
