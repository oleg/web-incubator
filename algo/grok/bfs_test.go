package grok

import (
	"algo/assert"
	"strings"
	"testing"
)

func TestBfsOnSimpleGraph(t *testing.T) {
	graph := testGraph()
	start := "a-1"

	item, found := bfs(graph, start, func(s string) bool {
		return strings.HasSuffix(s, "777")
	})

	assert.Equal(t, found, true)
	assert.Equal(t, "f-777", item)
}
func TestBfsOnSimpleGraphWrongStart(t *testing.T) {
	graph := testGraph()
	start := "XXX"

	item, found := bfs(graph, start, func(s string) bool {
		return strings.HasSuffix(s, "777")
	})

	assert.Equal(t, found, false)
	assert.Equal(t, "", item)
}
func TestBfsOnCircular(t *testing.T) {
	graph := map[string][]string{
		"a-1": {"b-1"},
		"b-1": {"c-1"},
		"c-1": {"a-1"},
	}
	start := "a-1"

	item, found := bfs(graph, start, func(s string) bool {
		return strings.HasSuffix(s, "777")
	})

	assert.Equal(t, found, false)
	assert.Equal(t, "", item)
}

func testGraph() map[string][]string {
	return map[string][]string{
		"a-1": {"b-1", "b-2", "b-3", "b-4"},
		"b-1": {"c-1", "c-2", "c-3"},
		"b-2": {"d-1", "d-2", "d-3", "d-4"},
		"b-3": {"e-1", "e-2"},
		"b-4": {"f-1", "f-777"},
	}

}

func TestQueue(t *testing.T) {
	q := queue{}

	q.enqueue("1")
	assertEqualQueue(t, q, queue{"1"})

	q.enqueue("2")
	assertEqualQueue(t, q, queue{"1", "2"})

	q.enqueue("3")
	assertEqualQueue(t, q, queue{"1", "2", "3"})

	q.enqueue("4")
	assertEqualQueue(t, q, queue{"1", "2", "3", "4"})

	item, found := q.dequeue()
	assert.Equal(t, found, true)
	assert.Equal(t, item, "1")
	assertEqualQueue(t, q, queue{"2", "3", "4"})

	item, found = q.dequeue()
	assert.Equal(t, found, true)
	assert.Equal(t, item, "2")
	assertEqualQueue(t, q, queue{"3", "4"})

	q.enqueue("5")
	assertEqualQueue(t, q, queue{"3", "4", "5"})

	item, found = q.dequeue()
	assert.Equal(t, found, true)
	assert.Equal(t, item, "3")
	assertEqualQueue(t, q, queue{"4", "5"})

	item, found = q.dequeue()
	assert.Equal(t, true, found)
	assert.Equal(t, "4", item)
	assertEqualQueue(t, queue{"5"}, q)

	item, found = q.dequeue()
	assert.Equal(t, true, found)
	assert.Equal(t, "5", item)
	assertEqualQueue(t, queue{}, q)

	item, found = q.dequeue()
	assert.Equal(t, false, found)
	assert.Equal(t, "", item)
	assertEqualQueue(t, queue{}, q)

	item, found = q.dequeue()
	assert.Equal(t, false, found)
	assert.Equal(t, "", item)
	assertEqualQueue(t, queue{}, q)

	q.enqueue("7")
	assertEqualQueue(t, queue{"7"}, q)
}

func assertEqualQueue(t *testing.T, actual, expected queue) {
	assert.EqualSlice(t, []string(actual), []string(expected))
}
