package gray

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_color(t *testing.T) {
	c := Color{-0.5, 0.4, 1.7}

	assert.Equal(t, -0.5, c.R)
	assert.Equal(t, 0.4, c.G)
	assert.Equal(t, 1.7, c.B)
}

func Test_adding_colors(t *testing.T) {
	c1 := Color{R: 0.9, G: 0.6, B: 0.75}
	c2 := Color{R: 0.7, G: 0.1, B: 0.25}

	result := c1.add(c2)

	AssertColorEqualInDelta(t, Color{1.6, 0.7, 1.0}, result)
}

func Test_subtracting_colors(t *testing.T) {
	c1 := Color{R: 0.9, G: 0.6, B: 0.75}
	c2 := Color{R: 0.7, G: 0.1, B: 0.25}

	result := c1.subtract(c2)

	AssertColorEqualInDelta(t, Color{0.2, 0.5, 0.5}, result)
}

func Test_multiplying_by_scalar(t *testing.T) {
	c1 := Color{R: 0.2, G: 0.3, B: 0.4}

	result := c1.multiplyByScalar(2)

	AssertColorEqualInDelta(t, Color{0.4, 0.6, 0.8}, result)
}

func Test_multiply_colors(t *testing.T) {
	c1 := Color{R: 1, G: 0.2, B: 0.4}
	c2 := Color{R: 0.9, G: 1, B: 0.1}

	result := c1.multiply(c2)

	AssertColorEqualInDelta(t, Color{0.9, 0.2, 0.04}, result)
}
