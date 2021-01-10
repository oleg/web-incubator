package stepik

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestReadIntervals(t *testing.T) {
	input := strings.NewReader(`3
1 2
3 4
5 6`)

	intervals := readIntervals(input)

	expected := []interval{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	assert.Equal(t, expected, intervals)
}

func TestWritePoints(t *testing.T) {
	output := new(bytes.Buffer)

	writePoints(output, []int{1, 3, 2})

	expected :=
		`3
1 3 2`
	assert.Equal(t, expected, output.String())
}

func TestPlaceDots1(t *testing.T) {
	intervals := []interval{
		{1, 3},
		{2, 5},
		{3, 6},
	}

	dots := placeDots(intervals)

	expected := []int{3}
	assert.Equal(t, expected, dots)
}

func TestPlaceDots2(t *testing.T) {
	intervals := []interval{
		{2, 3},
		{1, 4},
		{5, 6},
	}

	dots := placeDots(intervals)

	expected := []int{3, 6}
	assert.Equal(t, expected, dots)
}

func TestPlaceDotsEmpty(t *testing.T) {
	var intervals []interval

	dots := placeDots(intervals)

	assert.Empty(t, dots)
}

func TestInterval(t *testing.T) {
	i := interval{from: 1, to: 10}

	assert.True(t, i.contains(1))
	assert.True(t, i.contains(2))
	assert.True(t, i.contains(9))
	assert.True(t, i.contains(10))

	assert.False(t, i.contains(-10))
	assert.False(t, i.contains(0))
	assert.False(t, i.contains(11))
	assert.False(t, i.contains(100))
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

	assert.Equal(t, expected, output.String())
}
