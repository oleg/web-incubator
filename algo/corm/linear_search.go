package corm

//2.1-3
func LinearSearch(arr []int, elem int) (found bool, index int) {
	for i, e := range arr {
		if e == elem {
			return true, i
		}
	}
	return false, 0
}
