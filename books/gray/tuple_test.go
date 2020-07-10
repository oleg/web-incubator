package gray

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_properties(t *testing.T) {
	result := Tuple{1.1, 2.2, 3.3}

	assert.EqualValues(t, 1.1, result.x)
	assert.EqualValues(t, 2.2, result.y)
	assert.EqualValues(t, 3.3, result.z)
}

func Test_add(t *testing.T) {
	t1 := Tuple{3, -2, 5}
	t2 := Tuple{-2, 3, 1}

	result := t1.add(t2)

	assert.Equal(t, Tuple{1, 1, 6}, result)
}

func Test_subtract(t *testing.T) {
	t1 := Tuple{3, -2, 5}
	t2 := Tuple{-2, 3, 1}

	result := t1.subtract(t2)

	assert.Equal(t, Tuple{5, -5, 4}, result)
}

func Test_negate(t *testing.T) {
	t1 := Tuple{-2, 3, -1}

	result := t1.negate()

	assert.Equal(t, Tuple{2, -3, 1}, result)
}

func Test_multiply(t *testing.T) {
	t1 := Tuple{1, -2, 3}

	result := t1.multiply(3.5)

	assert.Equal(t, Tuple{3.5, -7, 10.5}, result)
}

func Test_multiply_by_fraction(t *testing.T) {
	t1 := Tuple{1, -2, 3}

	result := t1.multiply(0.5)

	assert.Equal(t, Tuple{0.5, -1, 1.5}, result)
}

func Test_divide(t *testing.T) {
	t1 := Tuple{1, -2, 3}

	result := t1.divide(2)

	assert.Equal(t, Tuple{0.5, -1, 1.5}, result)
}

