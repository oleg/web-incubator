package figure

import (
	"github.com/stretchr/testify/assert"
	"gray/multid"
	"gray/oned"
	"testing"
)

func Test(t *testing.T) {

}

func Test_ray_intersects_sphere_at_two_points(t *testing.T) {
	tests := []struct {
		name     string
		ray      oned.Ray
		expected []float64
	}{
		{"A ray intersects a sphere at two points",
			oned.Ray{oned.Point{0, 0, -5}, oned.Vector{0, 0, 1}},
			[]float64{4, 6}},
		{"A ray intersects a sphere at a tangent",
			oned.Ray{oned.Point{0, 1, -5}, oned.Vector{0, 0, 1}},
			[]float64{5, 5}},
		{"A ray misses a sphere",
			oned.Ray{oned.Point{0, 2, -5}, oned.Vector{0, 0, 1}},
			[]float64{}},
		{"A ray originates inside a sphere",
			oned.Ray{oned.Point{0, 0, 0}, oned.Vector{0, 0, 1}},
			[]float64{-1, 1}},
		{"A sphere is behind a ray",
			oned.Ray{oned.Point{0, 0, 5}, oned.Vector{0, 0, 1}},
			[]float64{-6, -4}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := Sphere()

			xs := Intersect(s, test.ray)

			assert.Equal(t, test.expected, xs)
		})
	}
}

func Intersect(s interface{}, ray oned.Ray) []interface{} {
	/*
		​ 	​# the vector from the sphere's center, to the ray origin​
		​ 	​# remember: the sphere is centered at the world origin​
		​ 	sphere_to_ray ← ray.origin - point(0, 0, 0)
		​
		​ 	a ← dot(ray.direction, ray.direction)
		​ 	b ← 2 * dot(ray.direction, sphere_to_ray)
		​ 	c ← dot(sphere_to_ray, sphere_to_ray) - 1
		​
		​ 	discriminant ← b² - 4 * a * c
	*/
	sphereToRay := ray.Origin.SubtractPoint(oned.Point{})
	a := multid.Matrix4{}

	return nil
}

func Sphere() interface{} {
	return nil
}

/*
​ 	​Scenario​: A ray intersects a sphere at two points
​ 	  ​Given​ r ← ray(point(0, 0, -5), vector(0, 0, 1))
​ 	    ​And​ s ← sphere()
​ 	  ​When​ xs ← intersect(s, r)
​ 	  ​Then​ xs.count = 2
​ 	    ​And​ xs[0] = 4.0
​ 	    ​And​ xs[1] = 6.0
*/
