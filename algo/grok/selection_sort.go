package grok

//SelectionSort sorts array in place
func SelectionSort(array []int) {
	l := len(array)
	for i := 0; i < l; i++ {
		min := i
		for j := i + 1; j < l; j++ {
			if array[j] < array[min] {
				min = j
			}
		}
		array[i], array[min] = array[min], array[i]
	}
}

func SelectionSortNew(array []int) []int {
	result := make([]int, 0, len(array))
	for len(array) != 0 {
		pos := 0
		min := array[pos]
		for i, v := range array {
			if v < min {
				min = v
				pos = i
			}
		}
		result = append(result, min)
		array = append(array[:pos], array[pos+1:]...)
	}
	return result
}

func SelectionSortNew2(array []int) []int {
	result := make([]int, 0, len(array))
	for len(array) != 0 {
		pos := 0
		for i, v := range array {
			if v < array[pos] {
				pos = i
			}
		}
		result = append(result, array[pos])
		array = append(array[:pos], array[pos+1:]...)
	}
	return result
}
