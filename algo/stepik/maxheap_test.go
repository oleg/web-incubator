package stepik

import (
	"algo/assert"
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestSolveMaxheap(t *testing.T) {
	in := `6
Insert 200
Insert 10
ExtractMax
Insert 5
Insert 500
ExtractMax`
	out := new(bytes.Buffer)

	solveMaxheap(strings.NewReader(in), out)

	if out.String() != "200\n500\n" {
		t.Errorf("Wrong output '%s'", out.String())
	}
}

func TestSolveMaxheap2(t *testing.T) {
	in := `8
Insert 200
Insert 10
Insert 5
Insert 500
ExtractMax
ExtractMax
ExtractMax
ExtractMax`
	out := new(bytes.Buffer)

	solveMaxheap(strings.NewReader(in), out)

	if out.String() != "500\n200\n10\n5\n" {
		t.Errorf("Wrong output '%s'", out.String())
	}
}
func TestSolveMaxheap3(t *testing.T) {
	in := `11
Insert 2
Insert 3
Insert 18
Insert 15
Insert 18
Insert 12
Insert 12
Insert 2
ExtractMax
ExtractMax
ExtractMax`
	out := new(bytes.Buffer)

	solveMaxheap(strings.NewReader(in), out)

	if out.String() != "18\n18\n15\n" {
		t.Errorf("Wrong output '%s'", out.String())
	}
}

func TestSolveMaxheap4(t *testing.T) {
	in := `9
Insert 53
Insert 7
Insert 22
Insert 6
Insert 5
Insert 21
Insert 20
ExtractMax
ExtractMax`
	out := new(bytes.Buffer)

	solveMaxheap(strings.NewReader(in), out)

	if out.String() != "53\n22\n" {
		t.Errorf("Wrong output '%s'", out.String())
	}
}

func TestSolveMaxheap5(t *testing.T) {
	in := `5
Insert 10
Insert 10
Insert 8
ExtractMax
ExtractMax`
	out := new(bytes.Buffer)

	solveMaxheap(strings.NewReader(in), out)

	if out.String() != "10\n10\n" {
		t.Errorf("Wrong output '%s'", out.String())
	}
}
func TestSolveMaxheap6(t *testing.T) {
	in := `37
Insert 3
Insert 0
ExtractMax
ExtractMax
Insert 32323
Insert 334
Insert 11111
ExtractMax
ExtractMax
Insert 323123123
Insert 100000000
Insert 323123123
Insert 100000000
Insert 323123123
Insert 100000000
Insert 323123123
Insert 100000000
Insert 323123123
Insert 100000000
Insert 323123123
Insert 100000000
Insert 323123123
Insert 100000000
Insert 323123123
Insert 100000000
ExtractMax
ExtractMax
ExtractMax
ExtractMax
ExtractMax
ExtractMax
ExtractMax
ExtractMax
ExtractMax
ExtractMax
ExtractMax
ExtractMax`
	out := new(bytes.Buffer)

	solveMaxheap(strings.NewReader(in), out)

	if out.String() != "3\n0\n32323\n11111\n323123123\n323123123\n323123123\n323123123\n323123123\n323123123\n323123123\n323123123\n100000000\n100000000\n100000000\n100000000\n" {
		t.Errorf("Wrong output '%s'", out.String())
	}
}

func TestSiftUpSmall(t *testing.T) {
	mh := maxheap{3, 2, 10}

	mh.siftUp(2)

	expected := maxheap{10, 2, 3}
	if !reflect.DeepEqual(mh, expected) {
		t.Errorf("Wrong result %v, expected %v", mh, expected)
	}
}

func TestSiftUpLong(t *testing.T) {
	mh := maxheap{10, 9, 8, 7, 6, 5, 4, 100}
	mh.siftUp(len(mh) - 1)

	expected := maxheap{100, 10, 8, 9, 6, 5, 4, 7}
	if !reflect.DeepEqual(mh, expected) {
		t.Errorf("Wrong result %v, expected %v", mh, expected)
	}
}

func TestExtractVal(t *testing.T) {
	tests := []struct {
		name     string
		expected int
		heap     maxheap
	}{
		{"case 1", 100, maxheap{100, 10, 8, 9, 6, 5, 4, 7}},
		{"case 2", 5, maxheap{5, 4, 3, 2, 1}},
		{"case 3", 3, maxheap{3, 2, 1}},
		{"case 4", 1, maxheap{1}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			extracted := test.heap.extract()
			if extracted != test.expected {
				t.Errorf("Wrong extracted element %v, expected %v", extracted, test.expected)
			}
			if !test.heap.isCorrect() {
				t.Errorf("Heap is in wrong state \n%v\n", test.heap.sprint())
			}
		})
	}
}

func TestSiftDown(t *testing.T) {
	tests := []struct {
		name string
		heap maxheap
	}{
		{"correct big", maxheap{100, 10, 8, 9, 6, 5, 4, 7}},
		{"correct small", maxheap{1}},
		{"incorrect one level", maxheap{2, 3, 1}},
		{"incorrect several levels", maxheap{1, 100, 10, 8, 9, 6, 5, 4, 7}},
		{"incorrect stop at middle", maxheap{25, 100, 80, 70, 40, 20, 10, 5}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.heap.siftDown(0)
			if !test.heap.isCorrect() {
				t.Errorf("Heap is in wrong state \n%v\n", test.heap.sprint())
			}
		})
	}
}

func TestIsCorrect(t *testing.T) {
	tests := []struct {
		name     string
		expected bool
		heap     maxheap
	}{
		{"case 1", true, maxheap{3, 2, 1}},
		{"case 2", true, maxheap{1}},
		{"case 3", false, maxheap{5, 3, 10}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			extracted := test.heap.isCorrect()
			if extracted != test.expected {
				t.Errorf("Wrong extracted element %v, expected %v, \n%v",
					extracted, test.expected, test.heap.sprint())
			}
		})
	}
}

func TestParent(t *testing.T) {
	assert.Equal(t, -1, parent(0))
	assert.Equal(t, 0, parent(1))
	assert.Equal(t, 0, parent(2))
	assert.Equal(t, 1, parent(3))
	assert.Equal(t, 1, parent(4))
	assert.Equal(t, 2, parent(5))
	assert.Equal(t, 2, parent(6))
}

func TestChildren(t *testing.T) {
	tests := []struct {
		parent, child1, child2 int
	}{
		{0, 1, 2},
		{1, 3, 4},
		{2, 5, 6},
		{3, 7, 8},
		{4, 9, 10},
		{5, 11, 12},
		{6, 13, 14},
		{7, 15, 16},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("case %d", test.parent), func(t *testing.T) {
			a, b := children(test.parent)
			assert.Equal(t, test.child1, a)
			assert.Equal(t, test.child2, b)
		})
	}
}
