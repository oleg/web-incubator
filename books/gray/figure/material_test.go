package figure

import (
	"github.com/stretchr/testify/assert"
	"gray/oned"
	"math"
	"testing"
)

func Test_default_material(t *testing.T) {
	m := DefaultMaterial()

	assert.Equal(t, oned.Color{1, 1, 1}, m.Color)
	assert.Equal(t, 0.1, m.Ambient)
	assert.Equal(t, 0.9, m.Diffuse)
	assert.Equal(t, 0.9, m.Specular)
	assert.Equal(t, 200.0, m.Shininess)
}

func Test_lighting(t *testing.T) {
	tests := []struct {
		name     string
		eyev     oned.Vector
		normalv  oned.Vector
		light    PointLight
		expected oned.Color
	}{
		{"Lighting with the eye between the light and the surface",
			oned.Vector{0, 0, -1},
			oned.Vector{0, 0, -1},
			PointLight{oned.Point{0, 0, -10}, oned.Color{1, 1, 1}},
			oned.Color{1.9, 1.9, 1.9}},
		{"Lighting with the eye between light and surface, eye offset 45°",
			oned.Vector{0, math.Sqrt2 / 2, -math.Sqrt2 / 2},
			oned.Vector{0, 0, -1},
			PointLight{oned.Point{0, 0, -10}, oned.Color{1, 1, 1}},
			oned.Color{1, 1, 1}},
		{"Lighting with eye opposite surface, light offset 45°",
			oned.Vector{0, 0, -1},
			oned.Vector{0, 0, -1},
			PointLight{oned.Point{0, 10, -10}, oned.Color{1, 1, 1}},
			oned.Color{0.7364, 0.7364, 0.7364}},
		{"Lighting with eye in the path of the reflection vector",
			oned.Vector{0, -math.Sqrt2 / 2, -math.Sqrt2 / 2},
			oned.Vector{0, 0, -1},
			PointLight{oned.Point{0, 10, -10}, oned.Color{1, 1, 1}},
			oned.Color{1.6364, 1.6364, 1.6364}},
		{"Lighting with the light behind the surface",
			oned.Vector{0, 0, -1},
			oned.Vector{0, 0, -1},
			PointLight{oned.Point{0, 0, 10}, oned.Color{1, 1, 1}},
			oned.Color{0.1, 0.1, 0.1}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			color := Lighting(DefaultMaterial(), test.light, oned.Point{}, test.eyev, test.normalv, false)

			oned.AssertColorEqualInDelta(t, test.expected, color)
		})
	}
}

func Test_lighting_with_surface_in_shadow(t *testing.T) {
	m := DefaultMaterial()
	eyeV := oned.Vector{0, 0, -1}
	normalV := oned.Vector{0, 0, -1}
	light := PointLight{oned.Point{0, 0, -10}, oned.Color{1, 1, 1}}

	r := Lighting(m, light, oned.Point{}, eyeV, normalV, true)

	assert.Equal(t, oned.Color{0.1, 0.1, 0.1}, r)
}

func Test_shadow(t *testing.T) {
	tests := []struct {
		name     string
		point    oned.Point
		expected bool
	}{
		{"There is no shadow when nothing is collinear with point and light",
			oned.Point{0, 10, 0}, false},
		{"The shadow when an object is between the point and the light",
			oned.Point{10, -10, 10}, true},
		{"There is no shadow when an object is behind the light",
			oned.Point{-20, 20, -20}, false},
		{"There is no shadow when an object is behind the point",
			oned.Point{-2, 2, -2}, false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := defaultWorld()

			r := w.IsShadowed(test.point)

			assert.Equal(t, test.expected, r)
		})
	}
}
