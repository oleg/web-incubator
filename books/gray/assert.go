package gray

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const delta = 0.000009

func AssertEqualInDelta(t *testing.T, expected, actual Vector) {
	assert.InDeltaMapValues(t, vectorToMap(expected), vectorToMap(actual), delta)
}
func vectorToMap(v Vector) map[string]float64 {
	return map[string]float64{"X": v.X, "Y": v.Y, "Z": v.Z}
}

func AssertColorEqualInDelta(t *testing.T, expected, actual Color) {
	assert.InDeltaMapValues(t, colorToMap(expected), colorToMap(actual), delta)
}
func colorToMap(v Color) map[string]float64 {
	return map[string]float64{"R": v.R(), "G": v.G(), "B": v.B()}
}

func AssertMatrixEqualInDelta(t *testing.T, expected, actual Matrix4) {
	assert.InDeltaMapValues(t, matrixToMap(expected), matrixToMap(actual), delta)
}
func matrixToMap(m Matrix4) map[string]float64 {
	r := map[string]float64{}
	for i, col := range m {
		for j, e := range col {
			k := fmt.Sprintf("%d:%d", i, j)
			r[k] = e
		}
	}
	return r
}
