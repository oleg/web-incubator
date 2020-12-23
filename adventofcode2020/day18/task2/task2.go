package task2

import (
	"fmt"
	"strconv"
	"strings"
	"text/scanner"
)

type calc struct {
	acc  int
	op   string
	mult []int
}

func (h *calc) pushOp(op string) {
	switch op {
	case "+":
		h.op = "+"
	case "*":
		h.mult = append(h.mult, h.acc)
		h.op = ""
	default:
		panic("unexpected op: " + h.op)
	}
}
func (h *calc) push(v int) {
	switch h.op {
	case "":
		h.acc = v
	case "+":
		h.acc += v
	//case "*":
	//	h.acc *= v
	default:
		panic("unexpected op: " + h.op)
	}
}

func (h *calc) get() int {
	for _, v := range h.mult {
		h.acc *= v
	}
	return h.acc
}

func Evaluate(data string) int {
	var c calc
	lex := &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	lex.scan.Init(strings.NewReader("(" + data + ")"))
	lex.next()
	read(lex, &c)
	return c.get()
}

type lexer struct {
	scan  scanner.Scanner
	token rune
}

func (lex *lexer) next() {
	lex.token = lex.scan.Scan()
}

func (lex *lexer) text() string {
	return lex.scan.TokenText()
}

func read(lex *lexer, c *calc) {
	switch lex.token {
	case scanner.Int:
		i, _ := strconv.Atoi(lex.text())
		c.push(i)
		lex.next()
		return
	case '+', '*':
		c.pushOp(lex.text())
		lex.next()
		return
	case '(':
		lex.next()
		var listC calc
		readList(lex, &listC)
		c.push(listC.get())
		lex.next()
		return
	}
	panic(fmt.Sprintf("unexpected token %q", lex.text()))
}

func readList(lex *lexer, res *calc) {
	for !endList(lex) {
		read(lex, res)
	}
}

func endList(lex *lexer) bool {
	switch lex.token {
	case scanner.EOF:
		panic("end of file")
	case ')':
		return true
	}
	return false
}
