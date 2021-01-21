package grok

func QuickSortLomuto(arr []int) {
	quickSortL(arr, 0, len(arr)-1)
}

func quickSortL(arr []int, lo, hi int) {
	if lo < hi {
		p := partitionL(arr, lo, hi)
		quickSortL(arr, lo, p-1)
		quickSortL(arr, p+1, hi)
	}
}

func partitionL(arr []int, lo int, hi int) int {
	i := lo
	for j := lo; j < hi; j++ {
		if arr[j] < arr[hi] {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[hi] = arr[hi], arr[i]
	return i
}

func QuickSortNew(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	if len(arr) == 2 {
		if arr[0] > arr[1] {
			arr[0], arr[1] = arr[1], arr[0]
		}
		return arr
	}
	pivot := arr[0]
	smaller := make([]int, 0)
	grater := make([]int, 0)
	for i := 1; i < len(arr); i++ {
		if arr[i] < pivot {
			smaller = append(smaller, arr[i])
		} else {
			grater = append(grater, arr[i])
		}
	}
	res := append(QuickSortNew(smaller), pivot)
	res = append(res, QuickSortNew(grater)...)
	return res
}
