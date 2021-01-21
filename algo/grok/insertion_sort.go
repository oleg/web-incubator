package grok

func InsertionSort(array []int) {
	for i := 1; i < len(array); i++ {
		curr := i
		for j := i - 1; j >= 0; j-- {
			if array[j] <= array[curr] {
				break
			}
			array[j], array[curr] = array[curr], array[j]
			curr = j
		}
	}
}

func InsertionSort2(array []int) {
	for i := 1; i < len(array); i++ {
		curr := i
		key := array[curr]
		beforeCurr := curr - 1
		for ; beforeCurr >= 0 && array[beforeCurr] > key; beforeCurr-- {
			array[curr] = array[beforeCurr]
			curr = beforeCurr
		}
		array[curr] = key
	}
}
