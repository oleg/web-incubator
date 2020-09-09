package figure

import (
	"github.com/stretchr/testify/assert"
	"gray/multid"
	"gray/oned"
	"testing"
)

func Test_ray_intersects_sphere_at_two_points(t *testing.T) {
	tests := []struct {
		name     string
		ray      Ray
		expected []float64
	}{
		{"A ray intersects a sphere at two points",
			Ray{oned.Point{0, 0, -5}, oned.Vector{0, 0, 1}},
			[]float64{4, 6}},
		{"A ray intersects a sphere at a tangent",
			Ray{oned.Point{0, 1, -5}, oned.Vector{0, 0, 1}},
			[]float64{5, 5}},
		{"A ray misses a sphere",
			Ray{oned.Point{0, 2, -5}, oned.Vector{0, 0, 1}},
			[]float64{}},
		{"A ray originates inside a sphere",
			Ray{oned.Point{0, 0, 0}, oned.Vector{0, 0, 1}},
			[]float64{-1, 1}},
		{"A sphere is behind a ray",
			Ray{oned.Point{0, 0, 5}, oned.Vector{0, 0, 1}},
			[]float64{-6, -4}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := MakeSphere()

			xs := s.Intersect(test.ray)

			assert.Len(t, xs, len(test.expected))
			for i, expected := range test.expected {
				assert.Equal(t, expected, xs[i].Distance)
			}
		})
	}
}

func Test_intersect_sets_object_on_intersection(t *testing.T) {
	r := Ray{oned.Point{0, 0, -5}, oned.Vector{0, 0, 1}}
	s := MakeSphere()

	res := s.Intersect(r)

	assert.Equal(t, 2, len(res))
	assert.Equal(t, s, res[0].Object)
	assert.Equal(t, s, res[1].Object)
}

func Test_sphere_default_transformation(t *testing.T) {
	s := MakeSphere()

	r := s.Transform()

	assert.Equal(t, multid.IdentityMatrix, r)
}
func Test_changing_sphere_transformation(t *testing.T) {
	tr := multid.Translation(2, 3, 4)
	s := Sphere{tr}

	r := s.Transform()

	assert.Equal(t, tr, r)
}

func Test_intersecting_scaled_sphere_with_ray(t *testing.T) {
	r := Ray{oned.Point{0, 0, -5}, oned.Vector{0, 0, 1}}
	s := Sphere{multid.Scaling(2, 2, 2)}

	xs := s.Intersect(r)

	assert.Equal(t, 2, len(xs))
	assert.Equal(t, 3., xs[0].Distance)
	assert.Equal(t, 7., xs[1].Distance)
}
func Test_intersecting_translated_sphere_with_ray(t *testing.T) {
	r := Ray{oned.Point{0, 0, -5}, oned.Vector{0, 0, 1}}
	s := Sphere{multid.Translation(5, 0, 0)}

	xs := s.Intersect(r)

	assert.Equal(t, 0, len(xs))
}
