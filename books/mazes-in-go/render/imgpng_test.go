package render

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"mazes/maze"
	"os"
	"testing"
)

func Test_renders_simple_1x1_grid_as_png(t *testing.T) {
	g := maze.NewGrid(1, 1)
	output, err := os.Create("test/output_1x1.png")
	assert.Nil(t, err)

	err = ToPng(g, output)

	assert.Nil(t, err)
	actual := mustReadFile(t, "test/output_1x1.png")
	expected := mustReadFile(t, "test/expected_output_1x1.png")
	assert.Equal(t, expected, actual)

}
func Test_renders_simple_2x2_grid_as_png(t *testing.T) {
	g := maze.NewGrid(2, 2)
	output, err := os.Create("test/output_2x2.png")
	assert.Nil(t, err)

	err = ToPng(g, output)

	assert.Nil(t, err)
	actual := mustReadFile(t, "test/output_2x2.png")
	expected := mustReadFile(t, "test/expected_output_2x2.png")
	assert.Equal(t, expected, actual)

}

func Test_renders_complex_4x4_grid_as_png(t *testing.T) {
	g := complex4x4Grid()
	output, err := os.Create("test/output_4x4.png")
	assert.Nil(t, err)

	err = ToPng(g, output)

	assert.Nil(t, err)
	actual := mustReadFile(t, "test/output_4x4.png")
	expected := mustReadFile(t, "test/expected_output_4x4.png")
	assert.Equal(t, expected, actual)
}

func mustReadFile(t *testing.T, filename string) []byte {
	actual, err := ioutil.ReadFile(filename)
	assert.Nil(t, err)
	return actual
}
