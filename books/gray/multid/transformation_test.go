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
