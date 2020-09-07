package multid

import (
	"github.com/stretchr/testify/assert"
	"gray/oned"
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
