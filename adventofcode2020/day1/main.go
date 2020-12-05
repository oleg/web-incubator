package main

import "github.com/oleg/incubator/adventofcode2020/io"

func main() {
	ints, err := io.ReadInts("day1/input.txt")
	if err != nil {
		panic(err)
	}
	a, b := find1(ints)
	println(a * b)
}

func find1(input []int) (int, int) {
	for _, a := range input {
		for _, b := range input {
			if a != b && a+b == 2020 {
				return a, b
			}
		}
	}
	return 0, 0
}

func find2(input []int) (int, int) {
	for _, a := range input {
		for _, b := range input {
			if a != b && a+b == 2020 {
				return a, b
			}
		}
	}
	return 0, 0
}
