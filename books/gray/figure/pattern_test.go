package figure

import (
	"github.com/stretchr/testify/assert"
	"gray/oned"
	"testing"
)

func Test_creating_stripe_pattern(t *testing.T) {
	pattern := StripePattern{oned.White, oned.Black}

	assert.Equal(t, oned.White, pattern.A)
	assert.Equal(t, oned.Black, pattern.B)
}

func Test_stripe_pattern_is_constant_in_y(t *testing.T) {
	pattern := StripePattern{oned.White, oned.Black}

	assert.Equal(t, oned.White, pattern.StripeAt(oned.Point{0, 0, 0}))
	assert.Equal(t, oned.White, pattern.StripeAt(oned.Point{0, 1, 0}))
	assert.Equal(t, oned.White, pattern.StripeAt(oned.Point{0, 2, 0}))
}

func Test_stripe_pattern_is_constant_in_z(t *testing.T) {
	pattern := StripePattern{oned.White, oned.Black}

	assert.Equal(t, oned.White, pattern.StripeAt(oned.Point{0, 0, 0}))
	assert.Equal(t, oned.White, pattern.StripeAt(oned.Point{0, 0, 1}))
	assert.Equal(t, oned.White, pattern.StripeAt(oned.Point{0, 0, 2}))
}

func Test_stripe_pattern_alternates_in_x(t *testing.T) {
	pattern := StripePattern{oned.White, oned.Black}

	assert.Equal(t, oned.White, pattern.StripeAt(oned.Point{0, 0, 0}))
	assert.Equal(t, oned.White, pattern.StripeAt(oned.Point{0.9, 0, 0}))

	assert.Equal(t, oned.Black, pattern.StripeAt(oned.Point{1, 0, 0}))
	assert.Equal(t, oned.Black, pattern.StripeAt(oned.Point{-0.1, 0, 0}))

	assert.Equal(t, oned.Black, pattern.StripeAt(oned.Point{-1, 0, 0}))
	assert.Equal(t, oned.White, pattern.StripeAt(oned.Point{-1.1, 0, 0}))
}
