package misc

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func MustOpen(filename string) io.Reader {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return file
}

//todo test?
func MustReadInts(reader io.Reader) []int {
	ints := make([]int, 0, 10)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		ints = append(ints, MustAtoi(scanner.Text()))
	}
	return ints
}

func MustAtoi(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return num
}

func ReadStructs(reader io.Reader, f func(string)) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		f(scanner.Text())
	}
}
