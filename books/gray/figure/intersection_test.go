package figure

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_intersection_encapsulates_distance_and_object(t *testing.T) {
	s := Sphere()

	i := Inter{3.5, s}

	assert.Equal(t, 3.5, i.Distance)
	assert.Equal(t, s, i.Object)
}

func Test_aggregating_intersections(t *testing.T) {
	s := Sphere()

	i1 := Inter{1, s}
	i2 := Inter{2, s}

	xs := Inters{i1, i2}

	assert.Equal(t, xs[0].Distance, 1.)
	assert.Equal(t, xs[1].Distance, 2.)
}

func Test_hit_when_all_intersections_have_positive_distance(t *testing.T) {
	s := Sphere()

	i1 := Inter{1, s}
	i2 := Inter{2, s}
	xs := Inters{i2, i1}

	_, i := xs.Hit()

	assert.Equal(t, i1, i)
}

func Test_hit_intersections(t *testing.T) {
	s := Sphere()
	tests := []struct {
		name                 string
		intersections        Inters
		expectedFound        bool
		expectedIntersection Inter
	}{
		{"all intersections have positive t",
			Inters{Inter{2, s}, Inter{1, s}}, true, Inter{1, s}},
		{"some intersections have negative t",
			Inters{Inter{1, s}, Inter{-1, s}}, true, Inter{1, s}},
		{"all intersections have negative t",
			Inters{Inter{-1, s}, Inter{-2, s}}, false, Inter{}},
		{"is always the lowest non negative intersection",
			Inters{Inter{5, s}, Inter{7, s}, Inter{-3, s}, Inter{2, s}}, true, Inter{2, s}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			ok, hit := test.intersections.Hit()

			assert.Equal(t, test.expectedFound, ok)
			assert.Equal(t, test.expectedIntersection, hit)
		})
	}
}
