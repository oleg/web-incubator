package stepik

import (
	"fmt"
	"io"
	"sort"
	"strings"
)

func solveHoffmanEncode(r io.Reader, w io.Writer) {
	writeHoffmannEncodeOutput(w, encodeHoffman(readHoffmannEncodeInput(r)))
}

func readHoffmannEncodeInput(r io.Reader) string {
	var input string
	_, _ = fmt.Fscan(r, &input)
	return input
}

func writeHoffmannEncodeOutput(w io.Writer, enc hoffmanEncoding) {
	_, _ = fmt.Fprintf(w, "%d %d\n", len(enc.dict), len(enc.code))

	keys := make([]string, 0, len(enc.dict))
	for k := range enc.dict {
		keys = append(keys, string(k))
	}
	sort.Strings(keys)
	for _, k := range keys {
		_, _ = fmt.Fprintf(w, "%s: %s\n", k, enc.dict[rune(k[0])])
	}
	_, _ = fmt.Fprint(w, enc.code)
}

type hoffmanEncoding struct {
	dict map[rune]string
	code string
}

type node struct {
	letter rune
	freq   int
	l, r   *node
}

func (n *node) fillPrefixDict(prefix string, dict map[rune]string) {
	if n == nil {
		return
	}
	dict[n.letter] = prefix
	n.l.fillPrefixDict(prefix+"0", dict)
	n.r.fillPrefixDict(prefix+"1", dict)
}

type priorityQueue []node

func (pq *priorityQueue) insert(n node) {
	*pq = append(*pq, n)
}

func (pq *priorityQueue) findMin() node {
	if len(*pq) == 0 {
		return node{}
	}
	mPos, min := 0, (*pq)[0]
	for i, v := range *pq {
		if v.freq < min.freq {
			min = v
			mPos = i
		}
	}
	*pq = append((*pq)[:mPos], (*pq)[mPos+1:]...)
	return min
}

func encodeHoffman(input string) hoffmanEncoding {
	freq := calcFreq(input)
	if len(freq) == 1 {
		return hoffmanEncoding{
			dict: map[rune]string{
				rune(input[0]): "0",
			},
			code: strings.Repeat("0", freq[rune(input[0])]),
		}
	}
	dict := makeDictionary(freq)
	code := encode(input, dict)
	return hoffmanEncoding{dict: dict, code: code}
}

func makeDictionary(freq map[rune]int) map[rune]string {
	pq := buildQueue(freq)
	dict := make(map[rune]string)
	pq[0].fillPrefixDict("", dict)
	delete(dict, 0)
	return dict
}

func encode(input string, dict map[rune]string) string {
	var b strings.Builder
	for _, v := range input {
		b.WriteString(dict[v])
	}
	return b.String()
}

func buildQueue(freq map[rune]int) priorityQueue {
	pq := priorityQueue{}
	for k, v := range freq {
		pq.insert(node{letter: k, freq: v})
	}
	for i := 0; i < len(freq)-1; i++ {
		n1 := pq.findMin()
		n2 := pq.findMin()
		if n1.freq > n2.freq {
			n1, n2 = n2, n1
		}
		pq.insert(node{letter: 0, freq: n1.freq + n2.freq, l: &n1, r: &n2})
	}
	return pq
}

func calcFreq(input string) map[rune]int {
	freq := make(map[rune]int)
	for _, v := range input {
		freq[v]++
	}
	return freq
}
