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
​ 	     | 5 | 6 | 7 | 8 |
​ 	     | 9 | 8 | 7 | 6 |
​ 	     | 5 | 4 | 3 | 2 |`)
	m2 := NewMatrix4(
		`| 1 | 2 | 3 | 4 |
​ 	     | 5 | 6 | 7 | 8 |
​ 	     | 9 | 8 | 7 | 6 |
​ 	     | 5 | 4 | 3 | 2 |`)

	assert.Equal(t, m1, m2)
}

func Test_matrices_not_equal(t *testing.T) {
	m1 := NewMatrix4(
		`| 1 | 2 | 3 | 4 |
​ 	     | 5 | 6 | 7 | 8 |
​ 	     | 9 | 8 | 7 | 6 |
​ 	     | 5 | 4 | 3 | 2 |`)
	m2 := NewMatrix4(
		`​| 2 | 3 | 4 | 5 |
​ 	     | 6 | 7 | 8 | 9 |
​ 	     | 8 | 7 | 6 | 5 |
​ 	     | 4 | 3 | 2 | 1 |`)

	assert.NotEqual(t, m1, m2)
}

/*
​ 	​Scenario​: Matrix equality with different matrices
​ 	  ​Given​ the following matrix A:
​ 	      | 1 | 2 | 3 | 4 |
​ 	      | 5 | 6 | 7 | 8 |
​ 	      | 9 | 8 | 7 | 6 |
​ 	      | 5 | 4 | 3 | 2 |
​ 	    ​And​ the following matrix B:
​ 	      | 2 | 3 | 4 | 5 |
​ 	      | 6 | 7 | 8 | 9 |
​ 	      | 8 | 7 | 6 | 5 |
​ 	      | 4 | 3 | 2 | 1 |
​ 	  ​Then​ A != B
*/
