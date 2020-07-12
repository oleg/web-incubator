package gray

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const delta = 0.000009

func AssertEqualInDelta(t *testing.T, expected Vector, actual Vector) {
	assert.InDeltaMapValues(t, vectorToMap(expected), vectorToMap(actual), delta)
}
func vectorToMap(v Vector) map[string]float64 {
	return map[string]float64{"X": v.X, "Y": v.Y, "Z": v.Z}
}

func AssertColorEqualInDelta(t *testing.T, expected Color, actual Color) {
	assert.InDeltaMapValues(t, colorToMap(expected), colorToMap(actual), delta)
}

func colorToMap(v Color) map[string]float64 {
	return map[string]float64{"R": v.R(), "G": v.G(), "B": v.B()}
}
