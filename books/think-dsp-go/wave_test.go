package think_dsp_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_create_wave_even(t *testing.T) {
	w := MakeWave([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 10)

	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, w.Ys)
	assert.Equal(t, []float64{0., 0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9}, w.Ts)
}

func Test_create_wave_odd(t *testing.T) {
	w := MakeWave([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10)

	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, w.Ys)
	assert.Equal(t, []float64{0., 0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1.}, w.Ts)
}
