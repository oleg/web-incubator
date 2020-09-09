package figure

import (
	"gray/multid"
	"gray/oned"
	"math"
)

type Sphere struct {
	transform multid.Matrix4
}

func MakeSphere() Sphere {
	return Sphere{multid.IdentityMatrix}
}

func (sphere Sphere) Transform() multid.Matrix4 {
	return sphere.transform
}

//todo or Sphere?
func (sphere Sphere) Intersect(worldRay Ray) Inters {
	ray := worldRay.Transform(sphere.Transform().Inverse())
	sphereToRay := ray.Origin.SubtractPoint(oned.Point{})
	a := ray.Direction.Dot(ray.Direction)
	b := 2 * ray.Direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1

	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return Inters{}
	}

	dSqrt := math.Sqrt(discriminant)
	t1 := (-b - dSqrt) / (2 * a)
	t2 := (-b + dSqrt) / (2 * a)
	return Inters{
		Inter{t1, sphere},
		Inter{t2, sphere},
	}
}
