package main

import (
	"bufio"
	"fmt"
	"github.com/oleg/incubator/adventofcode2020/misc"
	"io"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	d1 := newDevice()
	processInstructions(misc.MustOpen("day14/input.txt"), d1.setMask, d1.setMem)
	println(d1.sum())

	d2 := newDevice()
	processInstructions(misc.MustOpen("day14/input.txt"), d2.setMask, d2.setMem2)
	println(d2.sum())
}

type device struct {
	memory map[int64]int64
	mask   mask
}

func newDevice() *device {
	d := &device{}
	d.memory = make(map[int64]int64)
	d.mask = mask{}
	return d
}

func (d *device) setMask(m mask) {
	d.mask = m
}

func (d *device) setMem(addr, num int64) {
	num |= d.mask.ones
	num &= d.mask.zeros
	d.memory[addr] = num
}

func (d *device) setMem2(addr, num int64) {
	addr |= d.mask.ones
	l := len(d.mask.xs)
	top := int(math.Pow(2, float64(l)))
	for i := 0; i < top; i++ {
		s := []rune(fmt.Sprintf("%0"+fmt.Sprintf("%d", l)+"b", i))
		newAddrStr := []rune(fmt.Sprintf("%036b", addr))
		for p, v := range d.mask.xs {
			newAddrStr[v] = s[p]
		}
		newAddr, err := strconv.ParseInt(string(newAddrStr), 2, 64)
		if err != nil {
			panic(err)
		}
		d.memory[newAddr] = num
	}
}

func (d *device) sum() int64 {
	var sum int64
	for _, v := range d.memory {
		sum += v
	}
	return sum
}

type mask struct {
	zeros, ones int64
	xs          []int
}

func newMask(s string) mask {
	xs := make([]int, 0)
	var z, o strings.Builder
	for i, v := range s {
		switch v {
		case 'X':
			xs = append(xs, i)
			z.WriteRune('1')
			o.WriteRune('0')
		default:
			z.WriteRune(v)
			o.WriteRune(v)
		}
	}
	zeros, err := strconv.ParseInt(z.String(), 2, 64)
	if err != nil {
		panic(err)
	}
	ones, err := strconv.ParseInt(o.String(), 2, 64)
	if err != nil {
		panic(err)
	}
	return mask{zeros: zeros, ones: ones, xs: xs}
}

func (m mask) String() string {
	return fmt.Sprintf("%036b,%036b", m.zeros, m.ones)
}

func processInstructions(reader io.Reader, mask func(mask), operation func(int64, int64)) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		input := scanner.Text()
		if maskStr, found := getMaskStr(input); found {
			mask(newMask(maskStr))
		} else {
			operation(getMem(input))
		}
	}
}

func getMaskStr(input string) (string, bool) {
	prefix := "mask = "
	return strings.TrimPrefix(input, prefix), strings.HasPrefix(input, prefix)
}

func getMem(input string) (int64, int64) {
	re, err := regexp.Compile(`^mem\[(\d+)+] = (\d+)+`)
	if err != nil {
		panic(err)
	}
	memArgs := re.FindStringSubmatch(input)
	arg1, err := strconv.ParseInt(memArgs[1], 10, 36)
	if err != nil {
		panic(err)
	}
	arg2, err := strconv.ParseInt(memArgs[2], 10, 36)
	if err != nil {
		panic(err)
	}
	return arg1, arg2
}
