package sort

func insertionSort(arr []int) {
	for k := 1; k < len(arr); k++ {
		key := arr[k]
		p := k - 1
		for p >= 0 && arr[p] > key {
			arr[p+1] = arr[p]
			p--
		}
		arr[p+1] = key
	}
}
