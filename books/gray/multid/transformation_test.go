package multid

import (
	"github.com/stretchr/testify/assert"
	"gray/oned"
	"math"
	"testing"
)

func Test_multiply_point_by_translation_matrix(t *testing.T) {
	tr := Translation(5, -3, 2)
	p := oned.Point{-3, 4, 5}

	r := tr.multiplyPoint(p)

	assert.Equal(t, oned.Point{2, 1, 7}, r)
}

func Test_multiply_point_by_inverse_of_translation_matrix(t *testing.T) {
	tr := Translation(5, -3, 2)
	inv := tr.inverse()
	p := oned.Point{-3, 4, 5}

	r := inv.multiplyPoint(p)

	assert.Equal(t, oned.Point{-8, 7, 3}, r)
}

func Test_scaling_matrix_applied_to_point(t *testing.T) {
	tr := Scaling(2, 3, 4)
	p := oned.Point{-4, 6, 8}

	r := tr.multiplyPoint(p)

	assert.Equal(t, oned.Point{-8, 18, 32}, r)
}

func Test_scaling_matrix_applied_to_vector(t *testing.T) {
	tr := Scaling(2, 3, 4)
	v := oned.Vector{-4, 6, 8}

	r := tr.multiplyVector(v)

	assert.Equal(t, oned.Vector{-8, 18, 32}, r)
}

func Test_multiplying_inverse_of_scaling_matrix(t *testing.T) {
	tr := Scaling(2, 3, 4)
	inv := tr.inverse()
	v := oned.Vector{-4, 6, 8}

	r := inv.multiplyVector(v)

	assert.Equal(t, oned.Vector{-2, 2, 2}, r)
}

func Test_reflection_is_scaling_by_negative_value(t *testing.T) {
	tr := Scaling(-1, 1, 1)
	p := oned.Point{2, 3, 4}

	r := tr.multiplyPoint(p)

	assert.Equal(t, oned.Point{-2, 3, 4}, r)
}

func Test_rotating_point_around_x_axis(t *testing.T) {
	tests := []struct {
		name     string
		rotation float64
		expected oned.Point
	}{
		{"half quarter", math.Pi / 4, oned.Point{0, math.Sqrt2 / 2, math.Sqrt2 / 2}},
		{"full quarter", math.Pi / 2, oned.Point{0, 0, 1}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tr := RotationX(test.rotation)

			r := tr.multiplyPoint(oned.Point{0, 1, 0})

			AssertPointEqualInDelta(t, test.expected, r)
		})
	}
}

func Test_rotating_point_around_y_axis(t *testing.T) {
	tests := []struct {
		name     string
		rotation float64
		expected oned.Point
	}{
		{"half quarter", math.Pi / 4, oned.Point{math.Sqrt2 / 2, 0, math.Sqrt2 / 2}},
		{"full quarter", math.Pi / 2, oned.Point{1, 0, 0}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tr := RotationY(test.rotation)

			r := tr.multiplyPoint(oned.Point{0, 0, 1})

			AssertPointEqualInDelta(t, test.expected, r)
		})
	}
}

func Test_rotating_point_around_z_axis(t *testing.T) {
	tests := []struct {
		name     string
		rotation float64
		expected oned.Point
	}{
		{"half quarter", math.Pi / 4, oned.Point{-math.Sqrt2 / 2, math.Sqrt2 / 2, 0}},
		{"full quarter", math.Pi / 2, oned.Point{-1, 0, 0}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tr := RotationZ(test.rotation)

			r := tr.multiplyPoint(oned.Point{0, 1, 0})

			AssertPointEqualInDelta(t, test.expected, r)
		})
	}
}

func Test_inverse_of_x_rotation_rotates_in_opposite_direction(t *testing.T) {
	p := oned.Point{0, 1, 0}
	halfQuarter := RotationX(math.Pi / 4)
	inv := halfQuarter.inverse()

	r := inv.multiplyPoint(p)

	expected := oned.Point{0, math.Sqrt2 / 2, -math.Sqrt2 / 2}
	AssertPointEqualInDelta(t, expected, r)
}

func Test_shearing_transformation(t *testing.T) {
	tests := []struct {
		name           string
		transformation Matrix4
		expected       oned.Point
	}{
		{"x in proportion to y", Shearing(1, 0, 0, 0, 0, 0), oned.Point{5, 3, 4}},
		{"x in proportion to z", Shearing(0, 1, 0, 0, 0, 0), oned.Point{6, 3, 4}},
		{"y in proportion to x", Shearing(0, 0, 1, 0, 0, 0), oned.Point{2, 5, 4}},
		{"y in proportion to z", Shearing(0, 0, 0, 1, 0, 0), oned.Point{2, 7, 4}},
		{"z in proportion to x", Shearing(0, 0, 0, 0, 1, 0), oned.Point{2, 3, 6}},
		{"z in proportion to y", Shearing(0, 0, 0, 0, 0, 1), oned.Point{2, 3, 7}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tr := test.transformation
			p := oned.Point{2, 3, 4}

			r := tr.multiplyPoint(p)

			assert.Equal(t, test.expected, r)
		})
	}
}
