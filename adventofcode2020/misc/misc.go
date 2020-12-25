package misc

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func TrimNewLine(str string) string {
	return strings.TrimPrefix(str, "\n")
}

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

func MustReadFileToString(filename string) string {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
