package io

import (
	"bufio"
	"os"
	"strconv"
)

//todo test?
func ReadInts(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	ints := make([]int, 0, 10)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		ints = append(ints, number)
	}
	return ints, nil
}
