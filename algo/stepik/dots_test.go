package stepik

import (
	"algo/assert"
	"bytes"
	"strings"
	"testing"
)

func TestReadIntervals(t *testing.T) {
	input := strings.NewReader(`3
1 2
3 4
5 6`)

	intervals := readIntervals(input)

	assert.EqualSlice(t, intervals, []interval{
		{1, 2},
		{3, 4},
		{5, 6},
	})
}

func TestWritePoints(t *testing.T) {
	output := new(bytes.Buffer)

	writePoints(output, []int{1, 3, 2})

	expected :=
		`3
1 3 2`
	assert.Equal(t, output.String(), expected)
}

func TestPlaceDots1(t *testing.T) {
	intervals := []interval{
		{1, 3},
		{2, 5},
		{3, 6},
	}

	dots := placeDots(intervals)

	assert.EqualSlice(t, dots, []int{3})
}

func TestPlaceDots2(t *testing.T) {
	intervals := []interval{
		{2, 3},
		{1, 4},
		{5, 6},
	}

	dots := placeDots(intervals)

	assert.EqualSlice(t, dots, []int{3, 6})
}

func TestPlaceDotsEmpty(t *testing.T) {
	var intervals []interval

	dots := placeDots(intervals)

	assert.Equal(t, len(dots), 0)
}

func TestInterval(t *testing.T) {
	i := interval{from: 1, to: 10}

	assert.Equal(t, i.contains(1), true)
	assert.Equal(t, i.contains(2), true)
	assert.Equal(t, i.contains(9), true)
	assert.Equal(t, i.contains(10), true)

	assert.Equal(t, i.contains(-10), false)
	assert.Equal(t, i.contains(0), false)
	assert.Equal(t, i.contains(11), false)
	assert.Equal(t, i.contains(100), false)
}

func TestProcess(t *testing.T) {
	input := strings.NewReader(`3
1 2
3 4
5 6
`)
	output := new(bytes.Buffer)

	process(input, output)

	expected :=
		`3
2 4 6`

	assert.Equal(t, output.String(), expected)
}
