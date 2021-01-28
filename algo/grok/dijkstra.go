package grok

import (
	"container/list"
	"math"
)

type dijkstra struct {
	graph     map[string]map[string]int
	processed map[string]bool
	costs     map[string]int
	parents   map[string]string
}

func NewDijkstra(graph map[string]map[string]int) *dijkstra {
	return &dijkstra{
		graph:     graph,
		processed: make(map[string]bool),
		costs:     make(map[string]int),
		parents:   make(map[string]string),
	}
}

func (d *dijkstra) find(start, end string) ([]string, int) {
	for k, v := range d.graph[start] {
		d.costs[k] = v
		d.parents[k] = start
	}
	parent := d.lowestCost()
	for parent != "" {
		parentCost := d.costs[parent]
		for node, cost := range d.graph[parent] {
			newCost := parentCost + cost
			if _, found := d.costs[node]; !found {
				d.costs[node] = math.MaxInt32
			}
			if d.costs[node] > newCost {
				d.costs[node] = newCost
				d.parents[node] = parent
			}
		}
		d.processed[parent] = true
		parent = d.lowestCost()
	}
	l := list.New()
	for s := end; s != ""; s = d.parents[s] {
		l.PushFront(s)
	}
	path := make([]string, 0, l.Len())
	for e := l.Front(); e != nil; e = e.Next() {
		path = append(path, e.Value.(string))
	}
	return path, d.costs[end]
}

func (d *dijkstra) lowestCost() string {
	min := ""
	minCost := math.MaxInt32
	for node, cost := range d.costs {
		if cost < minCost && !d.processed[node] {
			minCost = cost
			min = node
		}
	}
	return min
}
