package search

func LinearSearch(arr []int, elem int) (found bool, index int) {
	for i, e := range arr {
		if e == elem {
			return true, i
		}
	}
	return false, 0
}
