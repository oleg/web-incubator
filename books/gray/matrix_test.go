package gray

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_create_matrix(t *testing.T) {
	m := NewMatrix4(
		`| 1    | 2    | 3    | 4    |  
		 | 5.5  | 6.5  | 7.5  | 8.5  |
		 | 9    | 10   | 11   | 12   |
		 | 13.5 | 14.5 | 15.5 | 16.5 |`)

	assert.Equal(t, 1., m[0][0])
	assert.Equal(t, 4., m[0][3])
	assert.Equal(t, 5.5, m[1][0])
	assert.Equal(t, 7.5, m[1][2])
	assert.Equal(t, 13.5, m[3][0])
	assert.Equal(t, 16.5, m[3][3])
}

func Test_matrices_equal(t *testing.T) {
	m1 := NewMatrix4(
		`| 1 | 2 | 3 | 4 |
 		 | 5 | 6 | 7 | 8 |
	     | 9 | 8 | 7 | 6 |
	     | 5 | 4 | 3 | 2 |`)
	m2 := NewMatrix4(
		`| 1 | 2 | 3 | 4 |
	     | 5 | 6 | 7 | 8 |
	     | 9 | 8 | 7 | 6 |
	     | 5 | 4 | 3 | 2 |`)

	assert.Equal(t, m1, m2)
}

func Test_matrices_not_equal(t *testing.T) {
	m1 := NewMatrix4(
		`| 1 | 2 | 3 | 4 |
	     | 5 | 6 | 7 | 8 |
	     | 9 | 8 | 7 | 6 |
	     | 5 | 4 | 3 | 2 |`)
	m2 := NewMatrix4(
		`​| 2 | 3 | 4 | 5 |
	     | 6 | 7 | 8 | 9 |
	     | 8 | 7 | 6 | 5 |
	     | 4 | 3 | 2 | 1 |`)

	assert.NotEqual(t, m1, m2)
}

func Test_multiply_matrices(t *testing.T) {
	m1 := NewMatrix4(
		`| 1 | 2 | 3 | 4 |
	     | 5 | 6 | 7 | 8 |
	     | 9 | 8 | 7 | 6 |
	     | 5 | 4 | 3 | 2 |`)
	m2 := NewMatrix4(
		`| -2 | 1 | 2 |  3 |
	     |  3 | 2 | 1 | -1 |
	     |  4 | 3 | 6 |  5 |
	     |  1 | 2 | 7 |  8 |`)

	result := m1.multiply(m2)

	expected := NewMatrix4(
		`| 20|  22 |  50 |  48 |
	     | 44|  54 | 114 | 108 |
	     | 40|  58 | 110 | 102 |
	     | 16|  26 |  46 |  42 |`)

	assert.Equal(t, expected, result)
}

func Test_multiply_matrix_and_vector(t *testing.T) {
	m := NewMatrix4(
		`| 1 | 2 | 3 | 4 |
	     | 2 | 4 | 4 | 2 |
	     | 8 | 6 | 4 | 1 |
	     | 0 | 0 | 0 | 1 |`)
	v := Vector{1, 2, 3}
	result := m.multiplyVector(v)

	expected := Vector{18, 24, 33}

	assert.Equal(t, expected, result)
}

func Test_multiply_arrays_empty(t *testing.T) {
	var a [][]float64 = nil
	var b [][]float64 = nil

	c := multiply(a, b)

	assert.Nil(t, c)
}

func Test_multiply_arrays_single(t *testing.T) {
	a := [][]float64{{2}}
	b := [][]float64{{3}}

	c := multiply(a, b)

	expected := [][]float64{{6}}

	assert.Equal(t, expected, c)
}

func Test_multiply_arrays_two_three_on_three_five(t *testing.T) {
	a := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}
	b := [][]float64{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 0},
		{3, 4, 5, 6, 7},
	}

	c := multiply(a, b)

	expected := [][]float64{
		{1*1 + 2*6 + 3*3, 1*2 + 2*7 + 3*4, 1*3 + 2*8 + 3*5, 1*4 + 2*9 + 3*6, 1*5 + 2*0 + 3*7},
		{4*1 + 5*6 + 6*3, 4*2 + 5*7 + 6*4, 4*3 + 5*8 + 6*5, 4*4 + 5*9 + 6*6, 4*5 + 5*0 + 6*7},
	}
	assert.Equal(t, expected, c)
}

func Test_multiply_matrix_by_identity_matrix(t *testing.T) {
	m := NewMatrix4(
		`| 0 | 1 |  2 |  4 |
		 | 1 | 2 |  4 |  8 |
		 | 2 | 4 |  8 | 16 |
		 | 4 | 8 | 16 | 32 |`)

	r := m.multiply(IdentityMatrix)

	assert.Equal(t, m, r)
}

func Test_multiply_identity_matrix_by_vector(t *testing.T) {
	v := Vector{1, 2, 3}

	r := IdentityMatrix.multiplyVector(v)

	assert.Equal(t, v, r)
}

func Test_transpose_matrix(t *testing.T) {
	m := NewMatrix4(
		`| 0 | 9 | 3 | 0 |
		 | 9 | 8 | 0 | 8 |
		 | 1 | 8 | 5 | 3 |
		 | 0 | 0 | 5 | 8 |`)

	r := m.transpose()
	expected := NewMatrix4(
		`| 0 | 9 | 1 | 0 |
		 | 9 | 8 | 8 | 0 |
		 | 3 | 0 | 5 | 5 |
		 | 0 | 8 | 3 | 8 |`)

	assert.Equal(t, expected, r)
}

func Test_transpose_identity_matrix(t *testing.T) {

	m := IdentityMatrix.transpose()

	assert.Equal(t, IdentityMatrix, m)
}
