package main

import "github.com/oleg/incubator/adventofcode2020/io"

func main() {
	ints, err := io.ReadInts("day1/input.txt")
	if err != nil {
		panic(err)
	}
	a1, b1 := find1(ints)
	println(a1 * b1)
	a2, b2, c2 := find2(ints)
	println(a2 * b2 * c2)
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

func find2(input []int) (int, int, int) {
	for _, a := range input {
		for _, b := range input {
			for _, c := range input {
				if a != b && a != c && b != c && a+b+c == 2020 {
					return a, b, c
				}
			}
		}
	}
	return 0, 0, 0
}
