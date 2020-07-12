package gray

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_convert_tuple_to_vector(t *testing.T) {
	vector := Vector{1.1, 2.1, 3.1}
	tuple := Tuple{1.1, 2.1, 3.1}
	assert.Equal(t, vector, Vector(tuple))
}

func Test_addVector_gives_vector(t *testing.T) {
	v1 := Vector{1, 2, 3}
	v2 := Vector{2, 3, 4}

	vector := v1.addVector(v2)

	assert.Equal(t, Vector{3, 5, 7}, vector)
}

func Test_subtractVector_gives_vector(t *testing.T) {
	v1 := Vector{5, 2, 2}
	v2 := Vector{5, 6, 1}

	vector := v1.subtractVector(v2)

	assert.Equal(t, Vector{0, -4, 1}, vector)
}

func Test_subtruct_zero_vector_negates_vector(t *testing.T) {
	zv := Vector{0, 0, 0}
	v1 := Vector{1, -2, 3}

	vector := zv.subtractVector(v1)

	assert.Equal(t, Vector{-1, 2, -3}, vector)
}

func Test_negate_negates_all_points(t *testing.T) {
	v := Vector{1, -2, 3}

	vector := v.negate()

	assert.Equal(t, Vector{-1, 2, -3}, vector)
}

func Test_magnitude_of_1_0_0(t *testing.T) {
	v := Vector{1, 0, 0}

	result := v.magnitude()

	assert.Equal(t, 1.0, result)
}

func Test_magnitude_of_0_1_0(t *testing.T) {
	v := Vector{0, 1, 0}

	result := v.magnitude()

	assert.Equal(t, 1.0, result)
}

func Test_magnitude_of_0_0_1(t *testing.T) {
	v := Vector{0, 0, 1}

	result := v.magnitude()

	assert.Equal(t, 1.0, result)
}

func Test_magnitude_of_1_2_3(t *testing.T) {
	v := Vector{1, 2, 3}

	result := v.magnitude()

	assert.Equal(t, math.Sqrt(14), result)
}

func Test_magnitude_of_m1_m2_m3(t *testing.T) {
	v := Vector{-1, -2, -3}

	result := v.magnitude()

	assert.Equal(t, math.Sqrt(14), result)
}

func Test_normalizing_vector_4_0_0(t *testing.T) {
	v := Vector{4, 0, 0}

	result := v.normalize()

	assert.Equal(t, Vector{1, 0, 0}, result)
}

func Test_normalizing_vector_1_2_3(t *testing.T) {
	v := Vector{1, 2, 3}

	result := v.normalize()

	AssertEqualInDelta(t, Vector{0.26726, 0.53452, 0.80178}, result)
}

func Test_magnitude_of_normalized_vector(t *testing.T) {
	v := Vector{1, 2, 3}

	result := v.normalize().magnitude()

	assert.Equal(t, 1.0, result)
}

func Test_dot_product_vector(t *testing.T) {
	v1 := Vector{1, 2, 3}
	v2 := Vector{2, 3, 4}

	result := v1.dot(v2)

	assert.Equal(t, 20.0, result)
}

func Test_cross_product_vector(t *testing.T) {
	v1 := Vector{1, 2, 3}
	v2 := Vector{2, 3, 4}

	assert.Equal(t, Vector{-1, 2, -1}, v1.cross(v2))
	assert.Equal(t, Vector{1, -2, 1}, v2.cross(v1))
}
