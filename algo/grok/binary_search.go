package grok

func BinarySearch(target int, array []int) (int, bool) {
	low := 0
	high := len(array) - 1
	for low <= high {
		middle := (low + high) / 2
		curr := array[middle]
		if curr == target {
			return middle, true
		}
		if curr < target {
			low = middle + 1
		}
		if curr > target {
			high = middle - 1
		}
	}
	return 0, false
}
