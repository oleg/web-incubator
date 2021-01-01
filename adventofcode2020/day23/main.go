package main

import (
	"fmt"
	"strings"
)

func main() {
	cb := newCycleBuffer([]int{2, 4, 7, 8, 1, 9, 3, 5, 6})
	moveN(100, cb)
	println(cb.String())
}

func moveN(n int, cb *cycleBuffer) {
	for i := 0; i < n; i++ {
		move(cb)
	}
}

func move(cb *cycleBuffer) {
	pickup := cb.cut3()
	dest := cb.dest(pickup)
	find := cb.find(dest)
	cb.insert(find, pickup)
}

type node struct {
	value int
	next  *node
}

type cycleBuffer struct {
	head *node
	len  int
}

func newCycleBuffer(items []int) *cycleBuffer {
	n := &node{value: items[0]}
	head := n
	for _, v := range items[1:] {
		tmp := node{value: v}
		n.next, n = &tmp, &tmp
	}
	n.next = head
	return &cycleBuffer{head: head, len: len(items)}
}

func (cb *cycleBuffer) cut3() *cycleBuffer {
	h := cb.head
	a := h.next
	b := a.next
	c := b.next
	h.next = c.next
	c.next = a
	cb.len = cb.len - 3
	return &cycleBuffer{head: a, len: 3}
}

func (cb *cycleBuffer) dest(pickup *cycleBuffer) int {
	next := func(curr int) int {
		curr--
		if curr == 0 {
			return cb.len + pickup.len
		}
		return curr
	}

	destination := next(cb.head.value)
	for pickup.contains(destination) {
		destination = next(destination)
	}
	return destination
}

func (cb *cycleBuffer) contains(value int) bool {
	return cb.find(value) != nil
}

func (cb *cycleBuffer) find(value int) *node {
	n := cb.head
	for i := 0; i < cb.len; i++ {
		if n.value == value {
			return n
		}
		n = n.next
	}
	return nil
}

func (cb *cycleBuffer) insert(after *node, other *cycleBuffer) {
	n := other.head
	for i := 0; i < other.len-1; i++ {
		n = n.next
	}
	after.next, n.next = other.head, after.next
	cb.len += other.len
	cb.head = cb.head.next
}

func (cb *cycleBuffer) String() string {
	n := cb.head
	var b strings.Builder
	for i := 0; i < cb.len; i++ {
		b.WriteString(fmt.Sprintf("%d", n.value))
		n = n.next
	}
	return b.String()
}
