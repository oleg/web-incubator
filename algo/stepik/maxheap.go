package stepik

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
)

func solveMaxheap(r io.Reader, w io.Writer) {
	executeMaxheapOperations(r, w)
}

func executeMaxheapOperations(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	size, _ := strconv.Atoi(scanner.Text())
	mh := maxheap(make([]int, 0, size))
	for scanner.Scan() {
		words := strings.Split(scanner.Text(), " ")
		switch words[0] {
		case "Insert":
			operand, _ := strconv.Atoi(words[1])
			mh.insert(operand)
		case "ExtractMax":
			_, _ = fmt.Fprintln(w, mh.extract())
		default:
			panic("wrong operation " + words[0])
		}
	}
}

type maxheap []int

func (m *maxheap) extract() int {
	l := len(*m) - 1
	m.swap(0, l)
	min := (*m)[l]
	*m = (*m)[:l]
	if len(*m) > 0 {
		m.siftDown(0)
	}
	return min
}

func (m *maxheap) insert(operand int) {
	*m = append(*m, operand)
	m.siftUp(len(*m) - 1)
}

func (m *maxheap) siftUp(pos int) {
	if pos == 0 {
		return
	}
	pPos := parent(pos)
	it := (*m)[pos]
	p := (*m)[pPos]
	if p < it {
		m.swap(pos, pPos)
	}
	m.siftUp(pPos)
}

func (m *maxheap) siftDown(pos int) {
	l := len(*m)
	it := (*m)[pos]
	aPos, bPos := children(pos)
	if aPos < l && bPos < l {
		a := (*m)[aPos]
		b := (*m)[bPos]
		if it < a || it < b {
			if a > b {
				m.swap(pos, aPos)
				m.siftDown(aPos)
			} else {
				m.swap(pos, bPos)
				m.siftDown(bPos)
			}
		}
		return
	}
	if aPos < l {
		a := (*m)[aPos]
		if it < a {
			m.swap(pos, aPos)
			m.siftDown(aPos)
		}
		return
	}
	if bPos < l {
		b := (*m)[bPos]
		if it < b {
			m.swap(pos, bPos)
			m.siftDown(bPos)
		}
		return
	}
}

func (m *maxheap) swap(pos1, pos2 int) {
	(*m)[pos1], (*m)[pos2] = (*m)[pos2], (*m)[pos1]
}

func (m *maxheap) isCorrect() bool {
	for i, v := range *m {
		if i != 0 && v > (*m)[parent(i)] {
			return false
		}
	}
	return true
}

func parent(i int) int {
	return int(math.Ceil(float64(i)/2) - 1)
}

func children(i int) (int, int) {
	return 2*i + 1, 2*i + 2
}

func (m *maxheap) sprint() string {
	var out strings.Builder
	n := len(*m)
	power := 0
	levelLength := math.Pow(2, float64(power))
	tab := n * 2
	for _, v := range *m {
		format := fmt.Sprintf("%%%dd", tab)
		levelLength--
		if levelLength == 0 {
			tab = tab/2 + tab/8
			power++
			levelLength = math.Pow(2, float64(power))
			format += "\n"
		}
		out.WriteString(fmt.Sprintf(format, v))
	}
	return out.String()
}
