package grok

func bfs(graph map[string][]string, start string, check func(string) bool) (string, bool) {
	q := &queue{start}
	visited := make(map[string]bool)
	for v, found := q.dequeue(); found; v, found = q.dequeue() {
		if _, y := visited[v]; y {
			continue
		}
		if check(v) {
			return v, true
		}
		visited[v] = true
		for _, v := range graph[v] {
			q.enqueue(v)
		}
	}
	return "", false
}

type queue []string

func (q *queue) enqueue(s string) {
	*q = append(*q, s)
}

func (q *queue) dequeue() (string, bool) {
	if len(*q) == 0 {
		return "", false
	}
	first := (*q)[0]
	*q = (*q)[1:]
	return first, true
}
