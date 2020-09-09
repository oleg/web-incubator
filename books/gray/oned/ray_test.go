package oned

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_creating_and_querying_a_ray(t *testing.T) {
	origin := Point{1, 2, 3}
	direction := Vector{4, 5, 6}

	ray := Ray{origin, direction}

	assert.Equal(t, origin, ray.Origin)
	assert.Equal(t, direction, ray.Direction)
}

func Test_Computing_point_from_distance(t *testing.T) {

	tests := []struct {
		name     string
		distance float64
		expected Point
	}{
		{"0", 0, Point{2, 3, 4}},
		{"1", 1, Point{3, 3, 4}},
		{"-1", -1, Point{1, 3, 4}},
		{"2.5", 2.5, Point{4.5, 3, 4}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := Ray{Point{2, 3, 4}, Vector{1, 0, 0}}
			p := r.Position(test.distance)
			assert.Equal(t, test.expected, p)
		})
	}
}
