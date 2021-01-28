package grok

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFind(t *testing.T) {
	graph := map[string]map[string]int{
		"book":   {"poster": 0, "lp": 5},
		"poster": {"guitar": 30, "drums": 35},
		"lp":     {"guitar": 15, "drums": 20},
		"guitar": {"piano": 20},
		"drums":  {"piano": 10},
	}

	d := NewDijkstra(graph)
	path, cost := d.find("book", "piano")
	fmt.Printf("%v\n", d.costs)
	fmt.Printf("%v\n", d.parents)
	assert.Equal(t, []string{"book", "lp", "drums", "piano"}, path)
	assert.Equal(t, 35, cost)
}
